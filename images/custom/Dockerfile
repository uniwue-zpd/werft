ARG TAG

FROM uniwuezpd/smw-core:${TAG}

MAINTAINER Maximilian Nöth maximilian.noeth@uni-wuerzburg.de

# Install additional dependencies via archive download (if necessary)
WORKDIR /tmp

## Install extensions with git
WORKDIR /var/www/html/extensions
RUN git clone -b REL1_39 https://github.com/wikimedia/mediawiki-extensions-SemanticACL SemanticACL
RUN git clone -b REL1_39 https://github.com/wikimedia/mediawiki-extensions-Workflows Workflows

## Copy crontab
COPY crontab /etc/cron.d/smw-cron
RUN chmod 0644 /etc/cron.d/smw-cron
RUN crontab /etc/cron.d/smw-cron
RUN touch /var/log/cron.log

### Update composer
## Install additional extension with composer
COPY composer.local.json .
RUN composer require --update-no-dev
