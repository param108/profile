.PHONY: build clean version

SHELL=/bin/bash

version:
	git branch --show-current > server/instance/version.txt
	git for-each-ref --points-at `git branch --show-current` | head -1 | cut -d' ' -f 1 >> server/instance/version.txt

build:	version
	- rm -rf build
	mkdir -p build
	go build -o build ./...

test:	version
	if [ -n "`goimports -l .`" ]; then echo "lint failure"; exit 1; fi
	go test -v ./...

clean:
	rm server/instance/version.txt
	rm -rf build

schema:
	pg_dump -s -x -O -d profile -h localhost -U profile --no-tablespace --no-comments  --no-publications --no-security-labels --no-subscriptions > db/schema.sql


db: 	build
	sudo su postgres /bin/bash -c psql < db/scripts/createdb.sql
	build/server migrate --migrationsPath db/migrations

db-destroy:
	sudo su postgres /bin/bash -c psql < db/scripts/destroydb.sql

db-ci:	build
	cp ci.env test.env
	cp ci.env .env
	PGPASSWORD=${PGPASSWORD} psql -h ${PGHOST} -U ${PGUSER} < db/scripts/createdb.sql
	build/server migrate --migrationsPath db/migrations
