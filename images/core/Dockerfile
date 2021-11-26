FROM uniwuezpd/smw_base:latest

# Install core extensions via composer
COPY composer.local.json .
RUN composer update --no-dev
RUN rm composer.local.json

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
## Semantic Drilldown
RUN wget -O /tmp/SemanticDrilldown.zip \
    https://github.com/wikimedia/mediawiki-extensions-SemanticDrilldown/archive/REL1_35.zip
RUN unzip -q /tmp/SemanticDrilldown.zip -d /var/www/html/extensions
## Data Transfer
WORKDIR /var/www/html/extensions
RUN git clone https://gerrit.wikimedia.org/r/mediawiki/extensions/DataTransfer
## Arrays
RUN wget -O /tmp/SemanticDrilldown.zip \
    https://extdist.wmflabs.org/dist/extensions/Arrays-REL1_35-8e95888.tar.gz
RUN tar -xzf Arrays-REL1_35-8e95888.tar.gz -C /var/www/mediawiki/extensions

# Clear tmp files
RUN rm -r /tmp/*
