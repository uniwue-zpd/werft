FROM uniwuezpd/smw_base:latest

# Install core extensions via composer
COPY composer.local.json .
RUN composer update --no-dev
RUN rm composer.local.json

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
## Data Transfer
WORKDIR /var/www/html/extensions
RUN git clone https://gerrit.wikimedia.org/r/mediawiki/extensions/DataTransfer
## Semantic Drilldown
WORKDIR /tmp/
RUN wget -O SemanticDrilldown.zip \
    https://github.com/wikimedia/mediawiki-extensions-SemanticDrilldown/archive/REL1_35.zip
RUN unzip -q SemanticDrilldown.zip -d /var/www/html/extensions
## Arrays
RUN wget -O Arrays.tar.gz \
    https://extdist.wmflabs.org/dist/extensions/Arrays-REL1_35-8e95888.tar.gz
RUN tar -xzf Arrays.tar.gz -C /var/www/html/extensions
## RegexFunctions
RUN wget -O RegexFunctions.tar.gz \
    https://extdist.wmflabs.org/dist/extensions/RegexFunctions-REL1_35-3de8a80.tar.gz
RUN tar -xzf /tmp/RegexFunctions.tar.gz -C /var/www/html/extensions
## Variables
RUN wget -O /tmp/Variables.tar.gz \
    https://extdist.wmflabs.org/dist/extensions/Variables-REL1_35-e8c7edc.tar.gz
RUN tar -xzf /tmp/Variables.tar.gz -C /var/www/html/extensions


# Clear tmp files
RUN rm -r /tmp/*
