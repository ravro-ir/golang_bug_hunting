from django.db import models


class Account(models.Model):
    balance = models.IntegerField(default=0)

    def __str__(self):
        return f"Account #{self.id}"

class Transaction(models.Model):
    account = models.ForeignKey(Account, on_delete=models.SET_NULL, null=True)
    amount = models.IntegerField(default=0)
    timestamp = models.DateTimeField(auto_now=True)

    def __str__(self):
        return f"<{self.account_id}, {self.amount}, {self.timestamp.isoformat()}>"
