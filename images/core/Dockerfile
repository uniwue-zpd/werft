ARG TAG=latest

FROM uniwuezpd/smw-base:${TAG}

LABEL vendor='Zentrum für Philologie und Digitalität "Kallimachos"' \
      maintainer='Maximilian Nöth <maximilian.noeth@uni-wuerzburg.de>'

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
WORKDIR /var/www/html/extensions

## Semantic Drilldown
WORKDIR /tmp/

# Install core extensions via composer
WORKDIR /var/www/html
COPY core.composer.local.json .
