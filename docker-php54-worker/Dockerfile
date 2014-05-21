# Docker php54 worker
#
# veresion 1.0

FROM aooj/php54:latest

MAINTAINER AooJ "aooj@n13.cz"


RUN mkdir /opt/install
ADD install/ /opt/install/

RUN chmod +x /opt/install/*.sh
RUN /opt/install/nodejs.sh

