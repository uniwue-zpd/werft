################################################################################################################
##### Base image which build on top of the official MediaWiki-Docker Image and installs Semantic MediaWiki #####
################################################################################################################

# Latest stable release according to https://www.mediawiki.org/wiki/Version_lifecycle (as of 17.01.2022)
FROM mediawiki:1.35.5

LABEL vendor='Zentrum für Philologie und Digitalität "Kallimachos"' \
      maintainer='Maximilian Nöth <maximilian.noeth@uni-wuerzburg.de>'

RUN apt-get update
RUN apt-get install -y zip unzip wget

COPY --from=composer /usr/bin/composer /usr/bin/composer

COPY composer.local.json .
RUN composer update --no-dev
RUN rm composer.local.json
