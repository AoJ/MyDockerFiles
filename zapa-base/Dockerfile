# zapa/base ver 1.0

FROM ubuntu:latest

MAINTAINER AooJ "aooj@n13.cz"
echo 1.0

RUN echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y install dialog net-tools lynx nano wget openssh-server python-software-properties supervisor
RUN apt-get update
RUN apt-get upgrade -y


# ssh
RUN mkdir /var/run/sshd
RUN mkdir -p /root/.ssh
ADD ./keys /tmp
RUN mkdir -p /opt/run
ADD ssh.sh /opt/run/ssh.sh
# for more keys put all keys to folder and write cmd to start.sh like "for a in ./folder/* do cat $a >> authorized_keys"



# debug
# docker build -t aooj/zapa-worker:1.0 .
# docker run -volumes-from zapa-repo -t -i aooj/zapa-worker:1.0 /bin/bash
# ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no root@localhost -p 49279 -v
#  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no root@localhost -p 49325 /opt/deploy/deploy.sh beta deploy2

# production
# docker run -volumes-from zapa-repo -p 80 -p 22 -d -t -i -name zapa_deploy2 aooj/zapa-worker:1.0 beta deploy2



# supervisor
# ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
