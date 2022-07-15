#!/bin/bash
set -e


python scripts/wait_for_mysql.py
python manage.py makemigrations
python manage.py migrate
python manage.py collectstatic --noinput
gunicorn project.wsgi:application --workers $GUNICORN_WORKERS --bind 0.0.0.0:8000
