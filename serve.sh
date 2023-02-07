#!/bin/sh

ng serve &
gin --port 8080 --path . --build ./src/server/ --i --all &

wait