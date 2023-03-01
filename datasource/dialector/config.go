package dialector

type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Params   string

	// DSN will override all above configs
	// Example:
	//  - Mysql: "user1@tcp(127.0.0.1:3306)/demo"
	// 	- Postgres: "host=127.0.0.1 user=user1 password=secret1 dbname=demo port=5432"
	//  - Sqlite:
	// 		- "sqlite.db"
	//  	- "file::memory:?cache=shared"
	Dsn string
}
