# /usr/bin/env python3

import os
from time import sleep

from MySQLdb import _mysql
from MySQLdb._exceptions import OperationalError

NUM_TRIES = 100
POLL_INTERVAL = 5

if __name__ == "__main__":
    while NUM_TRIES:
        try:
            print(f"Waiting for MySQL for {POLL_INTERVAL} seconds ({NUM_TRIES} tries remaining)")
            _mysql.connect(
                host=os.environ["MYSQL_HOST"],
                user=os.environ["MYSQL_USER"],
                passwd=os.environ["MYSQL_PASSWORD"],
                db=os.environ["MYSQL_DATABASE"],
            )
            print("MySQL is now ready.")
            exit(0)
        except OperationalError:
            sleep(POLL_INTERVAL)
            NUM_TRIES -= 1
    print("Failed to wait for MySQL.")
    exit(1)
