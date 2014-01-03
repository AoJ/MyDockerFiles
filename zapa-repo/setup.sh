/bin/bash -c "cat <(crontab -l) <(echo \"*/1 * * * * git --git-dir=/src/zapakatel.git fetch --all >> /var/log/git-zapakatel-fetch.log 2>&1 \") | crontab -"
