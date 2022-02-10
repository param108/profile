# Store

The `store` is an interface that encapsulates all data needs.

Under the store, there maybe different data storage components like postgres, redis etc.

The file `store_impl.go` is the glue layer that has the concrete object `StoreImpl` which
implements `store` and connects with storage component functionality.

# Component directories

Component specific code will be housed in component specific directories
like `postgres` `redis` etc.

# Models

Models will be stored in `api/models` directory so that it can be used by both
the store and the consumer (the app (say)).


