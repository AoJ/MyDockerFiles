# Project: dad - data access database
# Author: AooJ <aooj@n13.cz>
# Date: 2014
# usage:
#       make build      - build new image from Dockerfile
#       make run        - debug run already created image by tag
#       make debug      - build and run bash in docker, dont run /run/start.sh


NAME=aooj/dad
VERSION=1.0


build:
	docker build -t $(NAME):$(VERSION) .


run:
	docker run -t -p 4000:3306 -p 22 -i $(NAME):$(VERSION)


debug: build
	docker run -p 3306 -p 22 -t -i $(NAME):$(VERSION) /bin/bash
