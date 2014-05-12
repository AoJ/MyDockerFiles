#!/usr/bin/env bash

curl -XPUT 'http://localhost:9200/_river/dm-test-dealgroups/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-test",
        "collection": "dealgroups"
    },
    "index": {
        "name": "dm-dealgroups",
        "type": "dm-dealgroups"
    }
}'


curl -XPUT 'http://localhost:9200/_river/dm-test-admins/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-test",
        "collection": "admins"
    },
    "index": {
        "name": "dm-admins",
        "type": "dm-admins"
    }
}'

curl -XPUT 'http://localhost:9200/_river/dm-test-admingroups/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-test",
        "collection": "admingroups"
    },
    "index": {
        "name": "dm-admingroups",
        "type": "dm-admingroups"
    }
}'

curl -XPUT 'http://localhost:9200/_river/dm-test-groups/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-test",
        "collection": "groups"
    },
    "index": {
        "name": "dm-groups",
        "type": "dm-groups"
    }
}'




curl -XPUT 'http://localhost:9200/_river/dm-test-users/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-test",
        "collection": "users"
    },
    "index": {
        "name": "dm-users",
        "type": "dm-users"
    }
}'

curl -XPUT 'http://localhost:9200/_river/dm-test-deals/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-test",
        "collection": "deals"
    },
    "index": {
        "name": "dm-deals",
        "type": "dm-deals"
    }
}'

