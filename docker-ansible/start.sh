#!/usr/bin/env bash

# fix multilayer docker container

apt-get install -y less sshpass vim inetutils-ping python2.7 python-software-properties python-setuptools build-essential python-crypto python-dev
apt-add-repository ppa:rquillo/ansible -y
apt-get update && apt-get upgrade -y
apt-get -y install ansible
apt-get clean
