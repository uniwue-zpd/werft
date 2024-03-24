ARG TAG=latest

FROM mariadb:${TAG}

RUN apt-get update
RUN apt-get install -y cron

## Create backup directory
RUN mkdir -p /backup

## Copy scripts
ADD scripts /scripts

## Copy crontab
RUN echo "01 01 * * * sh /scripts/backup.sh" > /etc/cron.d/smw-cron
RUN chmod 0644 /etc/cron.d/smw-cron
RUN crontab /etc/cron.d/smw-cron
RUN touch /var/log/cron.log
