package golibdata

import (
	"database/sql"
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-data/datasource"
	"gitlab.com/golibs-starter/golib-data/datasource/dialector"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// DatasourceOpt provide option to bootstrap datasource with all available strategies.
func DatasourceOpt() fx.Option {
	return fx.Options(
		DatasourceCommonOpt(),
		ProvideDatasourceDialStrategy(dialector.NewMysql),
		ProvideDatasourceDialStrategy(dialector.NewPostgres),
		ProvideDatasourceDialStrategy(dialector.NewSqlite),
	)
}

// StrategicDatasourceOpt provide option to bootstrap datasource with specified strategies.
func StrategicDatasourceOpt(strategyConstructors ...interface{}) fx.Option {
	opts := make([]fx.Option, 0)
	opts = append(opts, DatasourceCommonOpt())
	for _, strategyConstructor := range strategyConstructors {
		opts = append(opts, ProvideDatasourceDialStrategy(strategyConstructor))
	}
	return fx.Options(opts...)
}

func DatasourceCommonOpt() fx.Option {
	return fx.Options(
		fx.Provide(NewDatasource),
		fx.Provide(newDialResolver),
		golib.ProvideHealthChecker(datasource.NewHealthChecker),
		golib.ProvideInformer(datasource.NewInformer),
		golib.ProvideProps(datasource.NewProperties),
	)
}

func ProvideDatasourceDialStrategy(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "datasource_dial_strategy",
		Target: constructor,
	})
}

type DatasourceOut struct {
	fx.Out
	Connection    *gorm.DB
	SqlConnection *sql.DB
}

func NewDatasource(resolver *dialector.Resolver, properties *datasource.Properties) (DatasourceOut, error) {
	out := DatasourceOut{}
	connection, err := datasource.NewConnection(resolver, properties)
	if err != nil {
		return out, errors.WithMessage(err, "cannot init datasource")
	}
	sqlConnection, err := connection.DB()
	if err != nil {
		return out, errors.WithMessage(err, "cannot get sqlDb instance")
	}
	out.Connection = connection
	out.SqlConnection = sqlConnection
	return out, nil
}

type NewDialResolverIn struct {
	fx.In
	DialStrategies []dialector.Strategy `group:"datasource_dial_strategy"`
}

func newDialResolver(in NewDialResolverIn) *dialector.Resolver {
	return dialector.NewResolver(in.DialStrategies)
}
