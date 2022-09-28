package golibdataTestUtil

import "go.uber.org/fx"

func DatabaseTestUtilOpt() fx.Option {
	return fx.Provide(NewDatabaseTestUtil)
}

func TruncateTablesOpt(tables ...string) fx.Option {
	return fx.Invoke(func(util *DatabaseTestUtil) {
		for _, table := range tables {
			util.TruncateTable(table)
		}
	})
}
