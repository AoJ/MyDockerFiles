# Project: docker ansible
# Author: AooJ <aooj@n13.cz>
# Date: 2014
# usage:
#	make build	- build new image from Dockerfile
#	make run	- debug run already created image by tag
#	make debug	- build and run bash in docker, dont run /run/start.sh


NAME=aooj/ansible
VERSION=1.0


build:
	docker build -t $(NAME):$(VERSION) .


run:
	$PUB=~/.ssh/id_rsa.pub
	[ -f $PUB ] && docker run -p 22 -t -i -e "SSH_PUBLIC_KEY=$(cat ~/.ssh/id_rsa.pub)" $(NAME):$(VERSION)
        [ ! -f $PUB ] && docker run -p 22 -t -i  $(NAME):$(VERSION)

debug: build
	docker run -t -i $(NAME):$(VERSION) /bin/bash	


ssh:
	@ID=$$(docker ps | grep -F "$(NAME):$(VERSION)" | awk '{ print $$1 }') && \
		if test "$$ID" = ""; then echo "Container is not running."; exit 1; fi && \
		IP=$$(docker inspect $$ID | grep IPAddr | sed 's/.*: "//; s/".*//') && \
		echo "SSHing into $$IP" && \
		ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no root@localhost

.PHONY: build run debug ssh
