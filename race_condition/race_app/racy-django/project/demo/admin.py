from django.contrib import admin

from .models import Account, Transaction


class TransactionInline(admin.TabularInline):
    extra = 0
    model = Transaction


class AccountAdmin(admin.ModelAdmin):
    inlines = [TransactionInline]


admin.site.register(Transaction)
admin.site.register(Account, AccountAdmin)
