ARG TAG=latest

FROM mariadb:${TAG}

## Update system and install cron
RUN apt-get update
RUN apt-get install -y cron

## Create necessary file and directory structure
RUN mkdir -p /backup
RUN chmod 777 /backup
RUN touch /mysql_backup.log
RUN chmod 666 /mysql_backup.log

## Copy scripts
ADD scripts /scripts
RUN chmod 755 /scripts/*
