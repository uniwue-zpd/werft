################################################################################################################
##### Base image which build on top of the official MediaWiki-Docker Image and installs Semantic MediaWiki #####
################################################################################################################

# Latest stable release according to https://www.mediawiki.org/wiki/Version_lifecycle (as of 14.03.2024)
# LTS EOL November 2025
ARG TAG=lts

FROM mediawiki:${TAG}

LABEL vendor='Zentrum für Philologie und Digitalität "Kallimachos"' \
      maintainer='Maximilian Nöth <maximilian.noeth@uni-wuerzburg.de>'

RUN apt-get update
RUN apt-get install -y zip unzip wget cron

# Setup Composer
COPY --from=composer:2.1 /usr/bin/composer /usr/bin/composer
ENV COMPOSER_HOME=/tmp

# Enable Short URLs in Apache
RUN set -eux; \
	a2enmod rewrite; \
	{ \
		echo "<Directory /var/www/html>"; \
		echo "  RewriteEngine On"; \
		echo "  RewriteCond %{REQUEST_FILENAME} !-f"; \
		echo "  RewriteCond %{REQUEST_FILENAME} !-d"; \
		echo "  RewriteRule ^ %{DOCUMENT_ROOT}/index.php [L]"; \
		echo "</Directory>"; \
	} > "$APACHE_CONFDIR/conf-available/short-url.conf"; \
	a2enconf short-url

# Enable AllowEncodedSlashes for VisualEditor in Apache
RUN sed -i "s/<\/VirtualHost>/\tAllowEncodedSlashes NoDecode\n<\/VirtualHost>/" "$APACHE_CONFDIR/sites-available/000-default.conf"

# Set recommended PHP.ini settings
# See https://secure.php.net/manual/en/opcache.installation.php
RUN { \
		echo 'opcache.memory_consumption=128'; \
		echo 'opcache.interned_strings_buffer=8'; \
		echo 'opcache.max_accelerated_files=4000'; \
		echo 'opcache.revalidate_freq=60'; \
	} > /usr/local/etc/php/conf.d/opcache-recommended.inis

RUN touch /wiki-debug.log

WORKDIR /var/www/html

# Install base extensions
COPY composer.local.json .
RUN composer require --update-no-dev
