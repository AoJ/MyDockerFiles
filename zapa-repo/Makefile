# Project: zapakatel repo with docker.io
# Author: AoJ <aoj@n13.cz>
# Date: 2014
# usage:
#	make build	- build new image from Dockerfile
#	make debug	- debug run already created image by tag
#	make run	- run production
#	make remove	- discard production
#	make try	- build and run in debug mode


AUTHOR=aooj
NAME=zapa-repo
VERSION=1.0
UID=$(AUTHOR)/$(NAME):$(VERSION)


build:
	docker build -t $(UID) .

run:
	docker run -d -t -i -p 22 -name $(NAME) $(UID)

remove:
	docker kill $(NAME)
	docker rm $(NAME)

debug:
	docker run -entrypoint="/bin/bash" -p 22 -t -i $(UID) -c /bin/bash


try: build debug


.PHONY: build run remove debug try

