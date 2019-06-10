#!/usr/bin/env bash

cd $(dirname $0)

for i in {1..100}
do
    curl -H "x_request_id:$i" http://127.0.0.1:8081/ &
done
