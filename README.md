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
options := []fx.Option{
    // When you want to use redis
    golib.RedisOpt(),
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


