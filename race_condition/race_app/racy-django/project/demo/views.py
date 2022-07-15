import random
import time

from django.db import transaction
from django.http import HttpRequest, HttpResponse
# Only because this is a quick and simple demo
from django.views.decorators.csrf import csrf_exempt

from .models import Account, Transaction


def random_delay() -> int:
    # Simulate random factors such as processing and validating the data, and
    # this one stupid network call that I was making for realtime related stuff
    # (which should have been put in a queue and delegated to a worker process).
    time.sleep(random.randint(3, 10))


@csrf_exempt
def atomic_long_delay(request: HttpRequest) -> HttpResponse:
    """
    In this example, we:
        - Use an atomic code block
        - Introduce a long (possibly random) time gap between when we read
        v/s when we commit the transaction (end of the atomic code block).
        - Don't use row locking.

    In this case if we have more than 1 (sync) gunicorn worker running and
    the frequency of requests to this endpoint (or any endpoint modifying the
    account balance) is greater than the time it takes for the atomic code
    block to run, then we can expect race conditions to occur.

    Typically, these kinds of things won't be found in development (unless you
    have really experienced dev teams with good code review skills).

    Note: We are not using using F() expressions due to complexities in the
    business logic. The random_delay represents processing the data related
    to the account and the request. During this time, we really don't want
    anything else to be using the account data (due to the risk of reading
    stale data).
    """
    try:
        account_id = int(request.POST["account"])
        amount = int(request.POST["amount"])
        if amount < 0:
            raise ValueError
    except (KeyError, ValueError):
        return HttpResponse(status=400)

    with transaction.atomic():
        try:
            # Possibly reading stale data:
            account = Account.objects.get(id=account_id)
        except Account.DoesNotExist:
            return HttpResponse(status=404)
        # Enough time for another request to be made or for the data read to
        # become stale:
        random_delay()
        account.balance += amount
        Transaction.objects.create(account=account, amount=amount)  # Ok.
        account.save()  # Overwritting possible.

    return HttpResponse(200)


@csrf_exempt
def non_atomic_long_delay(request: HttpRequest) -> HttpResponse:
    """
    This is more or less the same as atomic_long_delay. The only difference is
    that now the delay is between when we read v/s when we write instead of
    when we read v/s when we commit the transaction.
    
    In both of these examples we write at the very end of the transaction, so
    really, both of these functions are the exact same. I want to show that
    atomic transactions the real issue here - it's the delay b/w read and save.
    """
    try:
        account_id = int(request.POST["account"])
        amount = int(request.POST["amount"])
        if amount < 0:
            raise ValueError
    except (KeyError, ValueError):
        return HttpResponse(status=400)

    try:
        account = Account.objects.get(id=account_id)
    except Account.DoesNotExist:
        return HttpResponse(status=404)
    random_delay()
    account.balance += amount
    Transaction.objects.create(account=account, amount=amount)
    account.save()

    return HttpResponse(200)


@csrf_exempt
def atomic_no_delay(request: HttpRequest) -> HttpResponse:
    """
    Due to the absence of a large delay, the window for the race condition to
    occur is considerably reduced and you would need the request frequency to
    be much higher (and with more gunicorn workers to handle them in parallel)
    for a race condition to be triggered.

    Race conditions are now unlikely to occur but still very much possible.
    
    If you're only using 1 sync gunicorn worker then race conditions are not
    possible just like in the atomic_long_delay example; but you'll have
    terrible throughput and most requests would probably time out during high
    load.
    """
    try:
        account_id = int(request.POST["account"])
        amount = int(request.POST["amount"])
        if amount < 0:
            raise ValueError
    except (KeyError, ValueError):
        return HttpResponse(status=400)

    with transaction.atomic():
        try:
            account = Account.objects.get(id=account_id)
        except Account.DoesNotExist:
            return HttpResponse(status=404)
        account.balance += amount
        Transaction.objects.create(account=account, amount=amount)
        account.save()

    return HttpResponse(200)


@csrf_exempt
def non_atomic_no_delay(request: HttpRequest) -> HttpResponse:
    """
    Effectively the same as atomic_no_delay. Just like how atomic_long_delay
    and non_atomic_long_delay are effectively the same.
    """
    try:
        account_id = int(request.POST["account"])
        amount = int(request.POST["amount"])
        if amount < 0:
            raise ValueError
    except (KeyError, ValueError):
        return HttpResponse(status=400)

    try:
        account = Account.objects.get(id=account_id)
    except Account.DoesNotExist:
        return HttpResponse(status=404)
    account.balance += amount
    Transaction.objects.create(account=account, amount=amount)
    account.save()

    return HttpResponse(200)


@csrf_exempt
def row_locking_atomic_long_delay(request: HttpRequest) -> HttpResponse:
    """
    Here we demonstrate that even if we did have a long delay between when we
    read v/s when we commit the transaction, if we lock the row we are
    operating on, then race conditions will never occur.

    This isn't a silver bullet since by locking the row we're preventing every
    other endpoint depending on it from working (even the read-only endpoints).
    This could cause serious issues if the row is needed in multiple endpoints.

    select_for_update cannot be used outside of an SQL transaction (so atomic
    is necessary).

    Race conditions are now impossible.
    """
    try:
        account_id = int(request.POST["account"])
        amount = int(request.POST["amount"])
        if amount < 0:
            raise ValueError
    except (KeyError, ValueError):
        return HttpResponse(status=400)

    with transaction.atomic():
        try:
            account = Account.objects.select_for_update().get(id=account_id)
        except Account.DoesNotExist:
            return HttpResponse(status=404)
        random_delay()
        account.balance += amount
        Transaction.objects.create(account=account, amount=amount)
        account.save()

    return HttpResponse(200)
