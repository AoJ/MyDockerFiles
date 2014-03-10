#!/usr/bin/env bash


curl -XPUT 'http://localhost:9200/_river/deal-manage-test/_meta' -d '{
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
        "name": "deal-manage-test",
        "type": "deal-manage"
    }
}'

curl -XPUT 'http://localhost:9200/_river/deal-manage-production/_meta' -d '{
    "type": "mongodb",
    "mongodb": {
        "servers": [
                {"host": "itchy.allin1.cz", "port": "28000" }
        ],
        "credentials" : [
                {"db": "admin", "user":"deal-manage-test", "password": "0uieAlWjQto1XIj" }
        ],
        "db": "deal-manage-production",
        "collection": "deals"
    },
    "index": {
        "name": "deal-manage-production",
        "type": "deal-manage"
    }
}'

