# Project: zapakatel worker with docker.io and local folder
# Author: AoJ <aoj@n13.cz>
# Date: 2013
# usage:
#	make build	- build new image from Dockerfile
#	make debug	- debug run already created image by tag
#	make try	- build and run in debug mode
#	make $REV	- deploy in background repo revision


NAME=aooj/zapa-worker-local
VERSION=1.0
REPO=/home/aoj/src/repos/zapakatel
PREFIX=beta_
DEPLOY_TYPE=beta


build:
	docker build -t $(NAME):$(VERSION) .


debug:
	docker run -entrypoint="/bin/bash" -v $(REPO):/var/www -p 80 -p 22 -t -i $(NAME):$(VERSION) -c /bin/bash


try: build debug


.PHONY: build debug try

