#!/bin/sh
mkidr -p /bac
/var/www/html/maintenance php dumpBackup.php --full --quiet > dump.xml
