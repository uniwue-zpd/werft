ARG TAG=latest

FROM uniwuezpd/smw-base:${TAG}

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
## Data Transfer
WORKDIR /var/www/html/extensions
RUN git clone -b REL1_39 https://gerrit.wikimedia.org/r/mediawiki/extensions/DataTransfer
RUN git clone -b REL1_39 https://github.com/wikimedia/mediawiki-extensions-Echo Echo
RUN git clone -b REL1_39 https://github.com/wikimedia/mediawiki-extensions-LoginNotify LoginNotify
RUN git clone -b REL1_39 https://github.com/wikimedia/mediawiki-extensions-RegexFunctions RegexFunctions
RUN git clone https://github.com/SemanticMediaWiki/SemanticDrilldown SemanticDrilldown
## Semantic Drilldown
WORKDIR /tmp/

# Install core extensions via composer
WORKDIR /var/www/html
COPY composer.local.json .
RUN composer require --update-no-dev
