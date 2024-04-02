#!/bin/sh
DATE=$(date +%F)
mkdir -p /backup/${DATE}
/var/www/html/maintenance php dumpBackup.php --full --quiet > /backup/${DATE}/dump.xml
