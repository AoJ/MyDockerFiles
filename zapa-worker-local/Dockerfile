# AooJ/worker-local ver 1.2

FROM ubuntu:latest

MAINTAINER AooJ "aooj@n13.cz"
echo 1.2

RUN echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y install dialog net-tools lynx nano wget git openssh-server python-software-properties supervisor
RUN add-apt-repository -y ppa:nginx/stable
RUN add-apt-repository -y ppa:ondrej/php5-oldstable
RUN apt-get update
RUN apt-get upgrade -y

RUN apt-get -y install nginx php5-fpm php5-mysql php-apc php5-imagick php5-imap php5-mcrypt php5-curl php5-cli php5-cgi git-core
ADD nginx_default_host /etc/nginx/sites-available/default
RUN echo "cgi.fix_pathinfo = 0;" >> /etc/php5/fpm/php.ini
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

# configure
RUN sed -i '0,/worker_processes [0-9]*;/s//worker_processes 1;/' /etc/nginx/nginx.conf
RUN sed -i '0,/worker_connections [0-9]*;/s//worker_connections 50;/' /etc/nginx/nginx.conf

# supervisor
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# ssh
RUN mkdir /var/run/sshd
RUN mkdir -p /root/.ssh
ADD authorized_keys root/.ssh/authorized_keys
RUN chown root:root root/.ssh/authorized_keys
# for more keys put all keys to folder and write cmd to start.sh like "for a in ./folder/* do cat $a >> authorized_keys"


EXPOSE 22 80

ADD start.sh /opt/run/start.sh
ENTRYPOINT ["/opt/run/start.sh"]

# debug
# docker build -t aooj/zapa-worker:1.0 .
# docker run -volumes-from zapa-repo -t -i aooj/zapa-worker:1.0 /bin/bash
# ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no root@localhost -p 49279 -v
#  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no root@localhost -p 49325 /opt/deploy/deploy.sh beta deploy2

# production
# docker run -volumes-from zapa-repo -p 80 -p 22 -d -t -i -name zapa_deploy2 aooj/zapa-worker:1.0 beta deploy2
