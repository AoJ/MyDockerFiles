# AooJ/repo ver 1.0

FROM darron/chef-omnibus

MAINTAINER AooJ "aooj@n13.cz"

RUN echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y install dialog net-tools lynx nano wget git openssh-server

RUN mkdir /var/run/sshd

# ssh
RUN mkdir -p /root/.ssh
ADD id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

# git
RUN mkdir /src
RUN cd /src; git clone git@github.com:zapakatel/zapakatel.git
# RUN mkdir -p /src/zapakatel & cd /src/zapakatel
# RUN git init & git config core.sparsecheckout true
# RUN echo www/nette/ >> .git/info/sparse-checkout
# RUN git remote add origin git@github.com:zapakatel/zapakatel.git

EXPOSE 22

CMD /usr/sbin/sshd

# docker build -t aooj/repo:1.0 .
# docker run -t -i aooj/repo:1.0 -v /home/ci/src:/src:rw /bin/bash
