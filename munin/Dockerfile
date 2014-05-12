FROM darron/chef-omnibus

MAINTAINER AooJ "aooj@n13.cz"

RUN echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y install dialog net-tools lynx nano wget
RUN apt-get -y install python-software-properties
RUN add-apt-repository -y ppa:nginx/stable
RUN add-apt-repository -y ppa:ondrej/php5-oldstable
RUN apt-get update

RUN apt-get -y install nginx php5-fpm php5-mysql php-apc php5-imagick php5-imap php5-mcrypt git-core


RUN echo "cgi.fix_pathinfo = 0;" >> /etc/php5/fpm/php.ini
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

# configure
RUN sed -i '0,/worker_processes [0-9]*;/s//worker_processes 1;/' /etc/nginx/nginx.conf
RUN sed -i '0,/worker_connections [0-9]*;/s//worker_connections 50;/' /etc/nginx/nginx.conf


# munin
RUN mkdir -p            /var/run/munin
RUN apt-get install -y munin
ADD conf/munin.conf /etc/munin/munin.conf
RUN chown munin:munin   /var/run/munin
RUN chmod 0755          /var/run/munin

RUN rm /etc/nginx/sites-enabled/default
ADD conf/munin.nginx.conf /etc/nginx/sites-enabled/munin.conf

EXPOSE 80

CMD service php5-fpm start && nginx

# docker build -t aooj/munin:1.0 .
# docker run -t -i -p 80:80 aooj/munin:1.0 /bin/bash
