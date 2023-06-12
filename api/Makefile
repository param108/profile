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
	cp test.env .env
	go test -v ./...

clean:
	rm server/instance/version.txt
	rm -rf build

schema:
	pg_dump -s -x -O -d profile -h localhost -U profile --no-tablespace --no-comments  --no-publications --no-security-labels --no-subscriptions > db/schema.sql


db: 	build
	cp db/scripts/createdb.sql /tmp/createdb.sql
	pushd /tmp; sudo su postgres /bin/bash -c psql < /tmp/createdb.sql; popd
	build/server migrate --migrationsPath db/migrations

db-destroy:
	cp db/scripts/destroydb.sql /tmp/
	pushd /tmp; sudo su postgres /bin/bash -c psql < /tmp/destroydb.sql; popd

db-ci:	build
	cp ci.env test.env
	cp ci.env .env
	PGPASSWORD=${PGPASSWORD} psql -h ${PGHOST} -U ${PGUSER} < db/scripts/createdb.sql
	build/server migrate --migrationsPath db/migrations

migrate:
	./build/server migrate --migrationsPath db/migrations

dev:	build
	cp dev.env .env
	PGPASSWORD="profile" psql -h "localhost" -U "profile" < db/scripts/dev_init.sql
	./build/server serve
