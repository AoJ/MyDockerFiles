# Project: docker php54 worker
# Author: AooJ <aooj@n13.cz>
# Date: 2014
# usage:
#	make build	- build new image from Dockerfile
#	make run	- debug run already created image by tag
#	make debug	- build and run bash in docker, dont run /run/start.sh


NAME=aooj/php54-worker
VERSION=1.0


build:
	docker build -t $(NAME):$(VERSION) .


run:
	docker run -t -i $(NAME):$(VERSION) /bin/bash

debug: build
	docker run -t -i $(NAME):$(VERSION) /bin/bash

.PHONY: build run debug

