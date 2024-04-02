#!/bin/sh
touch /mysql_backup.log
echo "${CRON_INTERVAL} sh /scripts/backup.sh" > /etc/cron.d/smw-cron
chmod 0644 /etc/cron.d/smw-cron
crontab /etc/cron.d/smw-cron
touch /var/log/cron.log

cron && tail -f /mysql_backup.log
