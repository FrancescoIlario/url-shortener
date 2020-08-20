#!/bin/bash

if [ $# -eq 0 ]
  then
    echo "No arguments supplied: provide a link to shorten"
fi

curl $(curl -X POST -s localhost:8080/shorten/anon -d "{ \"Url\": \"$1\"}" | jq -r .Url)