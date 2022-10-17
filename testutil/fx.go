package golibdataTestUtil

import "go.uber.org/fx"

func DatabaseTestUtilOpt() fx.Option {
	return fx.Provide(EnableDatabaseTestUtil)
}

func TruncateTablesOpt(tables ...string) fx.Option {
	return fx.Invoke(func() {
		TruncateTables(tables...)
	})
}
