# Project: zapakatel worker with docker.io
# Author: AoJ <aoj@n13.cz>
# Date: 2013
# usage:
#	make build	- build new image from Dockerfile
#	make debug	- debug run already created image by tag
#	make try	- build and run in debug mode
#	make $REV	- deploy in background repo revision


NAME=aooj/zapa-worker
VERSION=1.0
REPO=zapa-repo
PREFIX=beta_
DEPLOY_TYPE=beta


build:
	docker build -t $(NAME):$(VERSION) .


debug:
	docker run -entrypoint="/bin/bash" -volumes-from $(REPO) -p 80 -p 22 -t -i $(NAME):$(VERSION) -c /bin/bash


%:
	docker run -volumes-from $(REPO) -p 80 -p 22 -d -t -i -name $(PREFIX)$@ $(NAME):$(VERSION) $(DEPLOY_TYPE) $@
#	sleep 5
#	docker logs $(PREFIX)$@


try: build debug


.PHONY: build debug try

