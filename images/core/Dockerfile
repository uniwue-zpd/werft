FROM uniwuezpd/smw_base:latest

# Install core extensions via composer
COPY composer.local.json .
RUN composer update --no-dev
RUN rm composer.local.json

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
RUN wget -O /tmp/SemanticDrilldown.zip \
    https://github.com/wikimedia/mediawiki-extensions-SemanticDrilldown/archive/REL1_35.zip
RUN unzip -q /tmp/SemanticDrilldown.zip -d /var/www/html/extensions

# Clear tmp files
RUN rm -r /tmp/*
