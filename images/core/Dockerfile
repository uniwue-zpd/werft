ARG TAG=latest

FROM uniwuezpd/smw-base:${TAG}

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
WORKDIR /var/www/html/extensions

## Semantic Drilldown
WORKDIR /tmp/

# Install core extensions via composer
WORKDIR /var/www/html
COPY core.composer.local.json .
