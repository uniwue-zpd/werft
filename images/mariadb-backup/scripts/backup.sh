#!/bin/sh
DATE=$(date +%F)
mariadb-dump --user=$DB_USER --password=$DB_PASSWORD --lock-tables --databases $DB_NAME > /backup/$DATE/db_dump.sql