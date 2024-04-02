#!/bin/sh
DATE=$(date +%F)
mkdir -p /backup/${DATE}
mariadb-dump --user=$DB_USER --password=$DB_PASSWORD --lock-tables --databases $DB_NAME > /backup/${DATE}/dump.sql
