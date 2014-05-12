build:
        docker build -t aooj/munin:1.0 .

run:
        docker run -t -i -p 85:80 aooj/munin:1.0 /bin/bash
        @echo "Ports mapping:"
        docker ps | grep "aooj/munin" | grep -o -P "\d*->\d*"
