################################################################################################################
##### Base image which build on top of the official MediaWiki-Docker Image and installs Semantic MediaWiki #####
################################################################################################################
FROM mediawiki:1.35.3

LABEL vendor='Zentrum für Philologie und Digitalität "Kallimachos"' \
      maintainer='Maximilian Nöth <maximilian.noeth@uni-wuerzburg.de>'

RUN apt-get update
RUN apt-get install -y zip unzip wget

COPY --from=composer /usr/bin/composer /usr/bin/composer

### Install Semantic MediaWiki
WORKDIR cd /var/www/html
RUN composer require mediawiki/semantic-media-wiki "3.2.3"
