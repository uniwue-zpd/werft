ARG TAG

FROM uniwuezpd/smw-core:${TAG}

MAINTAINER Maximilian Nöth maximilian.noeth@uni-wuerzburg.de


# Install additional dependencies (if necessary)
WORKDIR /tmp
## SemanticACL
RUN wget -O SemanticACL.tar.gz \
  https://extdist.wmflabs.org/dist/extensions/SemanticACL-REL1_39-1392e69.tar.gz
RUN tar -xzf SemanticACL.tar.gz -C /var/www/html/extensions
RUN rm -r SemanticACL.tar.gz
## Workflows
RUN wget -O Workflows.tar.gz \
    https://extdist.wmflabs.org/dist/extensions/Workflows-REL1_39-84de375.tar.gz
RUN tar -xzf Workflows.tar.gz -C /var/www/html/extensions
RUN rm -r Workflows.tar.gz

## Install extensions with git
WORKDIR /var/www/html/

## Copy crontab
COPY crontab /etc/cron.d/smw-cron
RUN chmod 0644 /etc/cron.d/smw-cron
RUN crontab /etc/cron.d/smw-cron
RUN touch /var/log/cron.log

### Update composer
## Install additional extension with composer
COPY composer.local.json .
RUN composer require --update-no-dev
