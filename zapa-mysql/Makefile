# Project: zapakatel worker with docker.io
# Author: AoJ <aoj@n13.cz>
# Date: 2014
# usage:
#	make build	- build new image from Dockerfile
#	make debug	- debug run already created image by tag
#	make try	- build and run in debug mode
#	make remove	- discarding current running production container
#	make deploy	- build new, discard old and run new in production
#	make logs	- display docker logs for production

NAME=aooj/zapa-mysql
ID=zapa-mysql
VERSION=1.0


build:
	docker build -t $(NAME):$(VERSION) .


debug:
	docker run -entrypoint="/bin/bash" -p 3306 -t -i $(NAME):$(VERSION) -c /bin/bash

remove:
	docker kill $(ID) 2>&1
	docker rm $(ID) 2>&1

deploy:
	docker run -p 3306 -t -i -d -name $(ID) $(NAME):$(VERSION)

production:
	docker run -d -i -t -p 3306:3306 -name $(ID)-data aooj/zapa-mysql-data:$(VERSION) /usr/bin/mysqld_safe
logs:
	docker logs $(ID)


try: build debug


.PHONY: build debug try

