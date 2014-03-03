#!/usr/bin/env bash

LOG=/var/log/start
touch $LOG

/opt/run/setup_mysql.sh

/usr/bin/supervisord >> $LOG &

tailf $LOG /var/log/supervisor/*

