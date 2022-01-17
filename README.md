# Golib Data

Database solutions for Golang project. Includes:

* MySQL database
* Postgres database
* Redis database

### Setup instruction

Base setup, see [GoLib Instruction](https://gitlab.com/golibs-starter/golib/-/blob/develop/README.md)

Both `go get` and `go mod` are supported.

```shell
go get gitlab.com/golibs-starter/golib-data
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
	"database/sql"
	red "github.com/go-redis/redis/v8"
	"gitlab.com/golibs-starter/golib-data"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	_ = []fx.Option{
		// When you want to use redis
		golibdata.RedisOpt(),

		// When you want to use datasource
		golibdata.DatasourceOpt(),

		// Demo way to using redis
		fx.Provide(funcUseRedis),
		fx.Provide(funcUseOrm),
		fx.Provide(funcUseNativeDbConnection),
	}
}

func funcUseRedis(redisClient *red.Client) {
	// do something with redis client
}

func funcUseOrm(db *gorm.DB) {
	// do something with gorm
}

func funcUseNativeDbConnection(db *sql.DB) {
	// do something with the native database connection
}
```

### Configuration

```yaml
app:
    # Configuration available for golib.DatasourceOpt()
    datasource:
        driver: mysql # Support mysql, postgres
        host: localhost # Define the database host
        port: 3306 # Define the database port
        database: sample # Define the database name
        username: root # Define the username for authentication
        password: secret # Define the password for authentication
        params: parseTime=true # Extra params to add to the connection string

    # Configuration available for golib.RedisOpt()
    redis:
        host: localhost # Define redis host. Default: `localhost`
        port: 6379 # Define redis port. Default: `6379`
        database: 0 # Define redis database. Default: `0`
        username: "" # Define redis username
        password: "" # Define redis password
```
