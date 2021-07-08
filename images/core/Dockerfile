# TODO: Change on release to Dockerhub address
FROM smwbase

# Install core extensions via composer
RUN composer require mediawiki/page-forms "^5.2.1"
RUN composer require mediawiki/semantic-cite "~2.1"
RUN composer require mediawiki/semantic-compound-queries "~2.1"
RUN composer require mediawiki/semantic-result-formats "~3.1"
RUN composer require mediawiki/maps "~8.0"

# Download and extract core extensions which can't be installed via composer
# Should be migrated to composer instalation as soon as available
RUN wget -O /tmp/SemanticDrilldown.zip https://github.com/wikimedia/mediawiki-extensions-SemanticDrilldown/archive/REL1_35.zip \
    && unzip -q /tmp/SemanticDrilldown.zip -d /var/www/html/extensions && rm /tmp/SemanticDrilldown.zip
