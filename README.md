# Golib Data

Database solutions for Golang project. Includes:
* Redis database
* MySQL database

### Setup instruction

Base setup, see [GoLib Instruction](https://gitlab.id.vin/vincart/golib/-/blob/develop/README.md)

Both `go get` and `go mod` are supported.
```shell
go get gitlab.id.vin/vincart/golib-data
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
    red "github.com/go-redis/redis/v8"
    "gitlab.id.vin/vincart/golib-data"
    "go.uber.org/fx"
)

func funcUseRedis(redisClient *red.Client) {
    // do something with redis client
}

func main() {
    options := []fx.Option{
        // When you want to register redis
        golibdata.RedisOpt(),

        // Demo way to using redis
        fx.Provide(funcUseRedis),
    }
}
```

### Configuration

```yaml
# Configuration available for golib.RedisOpt()
app:
    redis:
      host: localhost # Defines redis host. Default: `localhost`
      port: 6379 # Defines redis port. Default: `6379`
      database: 0 # Defines redis database. Default: `0`
      password: "" # Defines redis password
```
