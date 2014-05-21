# Docker zoot
#
# VERSION 1.0

FROM aooj/php54:latest

MAINTAINER AooJ "aoj@n13.cz"

RUN apt-get update && apt-get install -y subversion php5-xsl

# tmp cache, move to zoot.sh
WORKDIR /var/www
RUN svn co http://svn.zoot.cz/frontend/trunk --non-interactive --username tomas.hegenbart --password x82ayr6bd4

ADD files/zoot.sh /opt/run/zoot.sh
ADD files/php-config.conf /etc/php5/fpm/conf.d/99-php-config.conf

VOLUME /var/www

EXPOSE 80 443 22

# TODO
# disable short tag
# fpm/pool.d/... listen.owner = www-data
# listen.group = www-data
# listen.mode = 0660
