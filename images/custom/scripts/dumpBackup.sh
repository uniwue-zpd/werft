#!/bin/sh
DATE=$(date +%F)
mkdir -p /backup/${DATE}
php /var/www/html/maintenance/dumpBackup.php --full --quiet > /backup/${DATE}/dump.xml
