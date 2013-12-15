FROM darron/chef-omnibus

MAINTAINER AooJ "aooj@n13.cz"

RUN 		echo "deb http://archive.ubuntu.com/ubuntu/ precise universe" >> /etc/apt/sources.list
RUN 		apt-get update
RUN 		apt-get -y install dialog net-tools lynx nano wget
RUN 		apt-get install -y make gcc wget
RUN 		wget http://download.redis.io/releases/redis-2.8.2.tar.gz -O /tmp/redis.tar.gz
RUN 		(cd /tmp && tar zxf redis.tar.gz && cd redis-* && make install)
RUN 		rm -rf /tmp/*
EXPOSE		6379
# ENTRYPOINT	["/usr/local/bin/redis-server"]
# CMD 		["/redis/redis.conf"]



# docker build -t aooj/redis .
# docker run -i -t -d -p 6300:6379 aooj/redis /bin/bash
