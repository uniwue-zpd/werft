# TODO: Change on release to Dockerhub address
FROM smwcore

MAINTAINER Maximilian Nöth maximilian.noeth@uni-wuerzburg.de

# Install additional dependencies (if necessary)

# Install additional extension with composer

# Install extensions with git


### Update composer
RUN composer update

WORKDIR /var/www/html/

### Copy custom files into container
COPY apache.conf /etc/apache2/conf-enabled/99-wiki.conf
#COPY LocalSettings.php /var/www/html/LocalSettings.php
