#!/usr/bin/env bash

# add ssh public key from env
if [ "$SSH_PUBLIC_KEY" ]; then
 	echo "$SSH_PUBLIC_KEY" >> root/.ssh/authorized_keys
	echo "==> importing ssh public key from env"
fi
