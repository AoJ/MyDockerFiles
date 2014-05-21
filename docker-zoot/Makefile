# Project: docker zoot
# Author: AooJ <aooj@n13.cz>
# Date: 2014
# usage:
#	make build	- build new image from Dockerfile
#	make run	- debug run already created image by tag
#	make debug	- build and run bash in docker, dont run /run/start.sh


NAME=aooj/zoot
VERSION=1.0


build:
	docker build -t $(NAME):$(VERSION) .

run:
	docker run --link mysql-zoot:mysql -p 800:80 -p 22 -t -i --name zoot $(NAME):$(VERSION)

debug: build
	docker run --link mysql-zoot:mysql -p 800:80 -p 22 -t -i --name zoot $(NAME):$(VERSION) /bin/bash	


.PHONY: build run debug ssh
