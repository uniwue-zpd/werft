FROM uniwuezpd/smw_core:latest

MAINTAINER Maximilian Nöth maximilian.noeth@uni-wuerzburg.de

# Install additional dependencies (if necessary)
## Install additional extension with composer
COPY composer.local.json .
RUN composer update --no-dev
RUN rm composer.local.json
## Install extensions with git


### Update composer
RUN composer update

WORKDIR /var/www/html/
