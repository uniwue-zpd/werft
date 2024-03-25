ARG TAG=latest

FROM uniwuezpd/smw-base:${TAG}

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
## Data Transfer
WORKDIR /var/www/html/extensions
RUN git clone -b REL1_39 https://gerrit.wikimedia.org/r/mediawiki/extensions/DataTransfer
## Semantic Drilldown
WORKDIR /tmp/
RUN wget -O SemanticDrilldown.zip \
    https://github.com/gesinn-it-pub/mediawiki-extensions-SemanticDrilldown/archive/refs/tags/3.1.3.zip
RUN unzip -q SemanticDrilldown.zip -d /var/www/html/extensions
RUN mv /var/www/html/extensions/mediawiki-extensions-SemanticDrilldown-3.1.3 /var/www/html/extensions/SemanticDrilldown
RUN rm -r SemanticDrilldown.zip
## Arrays
RUN wget -O Arrays.tar.gz \
    https://extdist.wmflabs.org/dist/extensions/Arrays-REL1_39-097a761.tar.gz
RUN tar -xzf Arrays.tar.gz -C /var/www/html/extensions
RUN rm -r Arrays.tar.gz
## RegexFunctions
RUN wget -O RegexFunctions.tar.gz \
    https://extdist.wmflabs.org/dist/extensions/RegexFunctions-REL1_39-af31952.tar.gz
RUN tar -xzf /tmp/RegexFunctions.tar.gz -C /var/www/html/extensions
RUN rm -r RegexFunctions.tar.gz
## Echo (will be part of the MW core starting with v1.40)
RUN wget -O Echo.tar.gz \
  https://extdist.wmflabs.org/dist/extensions/Echo-REL1_39-15f9c05.tar.gz
RUN tar -xzf /tmp/Echo.tar.gz -C /var/www/html/extensions
RUN rm -r Echo.tar.gz
## LoginNotify (will be part of the MW core starting with v1.40)
RUN wget -O LoginNotify.tar.gz \
  https://extdist.wmflabs.org/dist/extensions/LoginNotify-REL1_39-2721431.tar.gz
RUN tar -xzf /tmp/LoginNotify.tar.gz -C /var/www/html/extensions
RUN rm -r LoginNotify.tar.gz

# Install core extensions via composer
WORKDIR /var/www/html
COPY composer.local.json .
RUN composer require --update-no-dev
