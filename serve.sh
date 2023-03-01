!/bin/sh

ng serve &

gin --port 4200 --path . --build ./src/server/ --i --all &

wait