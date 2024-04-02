#!/bin/bash
DATE=$(date +%F)
mkdir -p /backup/${DATE}
echo "[$(date +%T)] Dumping XML backup into /backup/${DATE}/dump.xml" >> /wiki-debug.log
/usr/local/bin/php /var/www/html/maintenance/dumpBackup.php --full --quiet > /backup/${DATE}/dump.xml 2>&1
