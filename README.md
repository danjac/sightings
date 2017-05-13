Database of UFO sightings. Go/React sample application.

Requires glide, nodejs, PostgreSQL

cd $GOPATH/src/github.com/danjac/sightings

Create database (TBD)

make

cd ui && npm start
cd .. && ./bin/importer source.csv
cd .. && ./bin/server
