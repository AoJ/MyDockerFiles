#!/usr/bin/env bash

#!/bin/bash

LOG=/var/log/all

touch $LOG

/usr/bin/supervisord >> $LOG &

sleep 3

/opt/deploy/deploy.sh $1 $2 $3 $4 $5 $6 $7 $8 $9 >> $LOG &

tail -f $LOG
# /var/log/supervisor/*
