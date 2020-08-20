#!/bin/bash

if [ $# -eq 0 ]
  then
    echo "No arguments supplied: provide a link to shorten"
fi

curl -X GET -s localhost:8080/metrics/$1 | jq