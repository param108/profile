# Database


The database is a local postgres instance on the production server. This instance will only serve
localhost.

The following variables are injected via environment variables.
1. username - `DB_USER`
2. password - `DB_PASS`
3. host     - `DB_HOST`
4. port     - `DB_PORT`
5. db name  - `DB_NAME`

## Setup

On Ubuntu:
```
make db
```

On Other OS:
1. create a local postgres db called `profile`
2. run the sql file `api/db/scripts/createdb.sql` as superuser.

## Migrations

We will use `https://github.com/golang-migrate/migrate` to run migrations.

`/profile/api/db/migrations/` will have the migrations. The format of the migration files
will be `xxx_<name of migration>.[up|down].sql`.

`xxx` will be a monotonically increasing number starting from `000`.

Each `up` migration must have its counterpart `down` migration.

## Schema file

Make sure you update the file `api/db/schema.sql` with the final schema **BEFORE** raising your
pull request.

```
make schema
```
should create it for you, using `pg_dump`.

## Migrate command

To trigger migrations, we will add a command `migrate` in the `cmd` directory.

`server migrate` 

without any other parameters it means a catchup migration to the latest migration file.

`server migrate -n` will migrate down by n

`server migrate +n` will migrate up by n

## Running migrations from cicd

To run migrations we need to 
1. stop old image
2. run the migrations
3. start the new image

The decision to run migrations or not will be based on the environment variable
`MIGRATE=[true|false]`

Default behaviour (if `MIGRATE` is not set) is to run migrations whenever `server.sh start` runs and we will migrate to the latest migration. If `MIGRATE=false` then we don't run migrations.

The decision to update the server binary will be based on the environment variable
`UPDATE=[true|false]`

Default behaviour (if `UPDATE` is not set) is to update the server binary whenever `server.sh start` runs and we will update to the latest image. If `UPDATE=false` then we let the current image run.

To do this, we will modify `server.sh` to do the necessary work.

## Reversing migrations or irregular migrations

An irregular migration is any migration where we don't migrate to the latest migration. This
assumes that the current image has all the necessary migrations. So we can use the existing binary
to perform the necessary migration.

1. Set `MIGRATE=false` and `UPDATE=false` in the production `API_ENV_CONFIG` secret.
2. push the new image change to master and wait for the deployment to finish.
   - as `MIGRATE` and `UPDATE` are set to `false` the `.env` and `db` files gets updated.
3. ssh into the server
4. manually stop the server using `systemctl` command
5. run the migrations manually using the command
   `/home/param/tribist/server migrate <+n|-n>`
   - you will need admin ssh credentials to do this.
6. manually start the server using `systemctl` command
   - you will need admin ssh credentials to do this.

## writer column

Every table will have a `writer` column which will be a `uuid`. This will be used primarily
by the tests to avoid clashes between store tests and end-to-end flow tests.

All teardown functions will only delete entries in the db based on the `writer` column, similarly
the asserts should be based on queries using the `writer` column. This way, even if two tests that 
modify tables run in parallel there will be no clash.
