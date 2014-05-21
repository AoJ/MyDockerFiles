# Docker redis

## Usage

### in Dockerfile
    FROM aooj/redis:latest

### build
    git clone https://github.com/AoJ/docker-redis.git
    make build
    
### and start it
    make run

### or build, start and attach
    make debug

    
## Changelog
- 1.1 update to aooj/base:1.5
- 1.0 first realese
