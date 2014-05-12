FROM ubuntu

MAINTAINER AooJ "aooj@n13.cz"

RUN echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y install apt-utils dialog
RUN apt-get -y install net-tools apt-utils nano wget
RUN apt-get -y install python-software-properties
RUN add-apt-repository -y ppa:nginx/stable
RUN add-apt-repository -y ppa:ondrej/php5-oldstable
RUN apt-get update

RUN apt-get -y install nginx php5-fpm php5-mysql php-apc php5-imagick php5-imap php5-mcrypt git-core php5-curl

RUN wget -O /etc/nginx/sites-available/default https://gist.github.com/zapakatel/7612990/raw/a615983202764446b1c430808a603b627653a351/nginx-sites-available-default
RUN echo "cgi.fix_pathinfo = 0;" >> /etc/php5/fpm/php.ini
RUN echo "daemon off;" >> /etc/nginx/nginx.conf
# RUN mkdir /var/www
# RUN echo "<?php phpinfo(); ?>" > /var/www/index.php


# configure
RUN sed -i '0,/worker_processes [0-9]*;/s//worker_processes 1;/' /etc/nginx/nginx.conf
RUN sed -i '0,/worker_connections [0-9]*;/s//worker_connections 50;/' /etc/nginx/nginx.conf

EXPOSE 80

CMD service php5-fpm start && nginx

# docker build -t aooj/php:1.0 .
# docker run -t -p 80:80 -i  -v /home/aoj/src/repos:/var/www:rw aooj/php:1.0 /bin/bash
