# Docker ansible
#
# VERSION 1.0

FROM aooj/base:latest

MAINTAINER AooJ "aoj@n13.cz"

# debug
RUN apt-get update && apt-get upgrade -y && apt-get clean

ADD files/install.sh /tmp/install.sh
RUN /tmp/install.sh
RUN rm /tmp/install.sh
RUN mkdir /var/ansible


ADD files/start.sh /opt/run/start.sh

VOLUME ["/var/ansible"]
