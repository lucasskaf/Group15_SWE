#!/bin/sh

ng serve &
<<<<<<< HEAD
gin --port 4201 --path . --build ./src/server/ --i --all &
=======
gin --port 4200 --path . --build ./src/server/ --i --all &
>>>>>>> origin/Lucas

wait