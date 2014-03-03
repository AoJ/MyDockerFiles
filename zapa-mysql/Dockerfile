FROM ubuntu:latest
 
# RUN dpkg-divert --local --rename --add /sbin/initctl
# RUN ln -s /bin/true /sbin/initctl
 
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt-get upgrade -y
 
# a mounted file systems table to make MySQL happy
RUN cat /proc/mounts > /etc/mtab

RUN apt-get -y install pwgen openssh-server mysql-client mysql-server supervisor && apt-get clean

# supervisor
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf

 
# mysql
RUN sed -i -e"s/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/" /etc/mysql/my.cnf
RUN sed -i '/^datadir*/ s|/var/lib/mysql|/data/mysql|' /etc/mysql/my.cnf
RUN mkdir -p /data/mysql
RUN rm -Rf /var/lib/mysql
ADD mysql.cnf /etc/mysql/conf.d/mysql.cnf 

# ssh
RUN mkdir /var/run/sshd
RUN mkdir -p /root/.ssh
ADD authorized_keys root/.ssh/authorized_keys
RUN chown root:root root/.ssh/authorized_keys



ADD start.sh /opt/run/start.sh
ADD ./setup_mysql.sh /opt/run/setup_mysql.sh

 
EXPOSE 3306
 
CMD ["/opt/run/start.sh"]

# sync mysql time mysqldump -uzapakateldebug -hlisa.allin1.cz -ptR9fruW9 zapakateldebug | mysql -hlocalhost zapakatel
# pass x2USQ8E8Nte5WX6v, zapakatel a6Tk8BcQZd4cQJNk
# image 1.0 8bf14a8e4f5cf53660b1ecafbf42560156e66ae9935accca58b415a34c6ba313
# docker commit -m "zapakatel db with lisa clone data, first release" -author="AooJ" zapa-mysql aooj/zapa-mysql-data:1.0

