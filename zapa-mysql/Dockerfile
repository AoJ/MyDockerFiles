FROM ubuntu:latest
 
RUN dpkg-divert --local --rename --add /sbin/initctl
RUN ln -s /bin/true /sbin/initctl
 
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt-get upgrade -y
 
RUN apt-get -y install mysql-client mysql-server
 
RUN sed -i -e"s/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/" /etc/mysql/my.cnf
 
ADD ./startup.sh /opt/startup.sh
 
EXPOSE 3306
 
CMD ["/bin/bash", "/opt/startup.sh"]

#sync mysql time mysqldump -uzapakateldebug -hlisa.allin1.cz -ptR9fruW9 zapakateldebug | mysql -hlocalhost zapakatel
#pass x2USQ8E8Nte5WX6v, zapakatel a6Tk8BcQZd4cQJNk
# image 1.0 8bf14a8e4f5cf53660b1ecafbf42560156e66ae9935accca58b415a34c6ba313
# docker commit -m "zapakatel db with lisa clone data, first release" -author="AooJ" zapa-mysql aooj/zapa-mysql-data:1.0

