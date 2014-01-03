# AooJ/repo ver 1.0

FROM ubuntu:latest

MAINTAINER AooJ "aooj@n13.cz"

RUN echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get -y install dialog net-tools lynx nano wget git openssh-server cron supervisor

RUN mkdir -p /var/run/sshd
RUN mkdir -p /var/log/supervisor

# ssh
RUN mkdir -p /root/.ssh
ADD id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

# supervisor
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# git
RUN mkdir /src
RUN cd /src; git clone --mirror git@github.com:zapakatel/zapakatel.git
# mirroring whole repo in background, require run cron with "cron -f"
ADD ./setup.sh /
RUN sh /setup.sh
# RUN "cat <(crontab -l) <(echo \"*/1 * * * * git --git-dir=/src/zapakatel.git fetch --all >> /var/log/git-zapakatel-fetch.log 2>&1 \") | crontab -"

# RUN mkdir -p /src/zapakatel & cd /src/zapakatel
# RUN git init & git config core.sparsecheckout true
# RUN echo www/nette/ >> .git/info/sparse-checkout
# RUN git remote add origin git@github.com:zapakatel/zapakatel.git

EXPOSE 22

VOLUME ["/src"]

CMD ["/usr/bin/supervisord"]

# production
# docker build -t aooj/repo:1.0 .
# docker run -d -p 22 -name zapa-repo aooj/repo:1.0
# docker logs #ID

# debug
# docker run -p 22 -t -i aooj/repo:1.0 /bin/bash
