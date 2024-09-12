#!/bin/bash
echo "[$(date +%T)] Running jobs in queue" >> /wiki-debug.log
/usr/local/bin/php /var/www/html/extensions/SemanticMediaWiki/maintenance/rebuildData.php -d 50
