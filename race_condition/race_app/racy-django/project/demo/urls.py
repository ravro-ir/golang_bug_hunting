
from django.urls import path

from . import views

urlpatterns = [
    path("atomic_long_delay/", views.atomic_long_delay),
    path("non_atomic_long_delay/", views.non_atomic_long_delay),
    path("atomic_no_delay/", views.atomic_no_delay),
    path("non_atomic_no_delay/", views.non_atomic_no_delay),
    path("row_locking_atomic_long_delay/", views.row_locking_atomic_long_delay),
]