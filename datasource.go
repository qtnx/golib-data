package golibdata

import (
	"database/sql"
	"github.com/pkg/errors"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-data/datasource"
	"gitlab.id.vin/vincart/golib-data/datasource/dialector"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func DatasourceOpt() fx.Option {
	return fx.Options(
		fx.Provide(NewDatasource),
		fx.Provide(newDialResolver),
		golib.ProvideHealthChecker(datasource.NewHealthChecker),
		golib.ProvideProps(datasource.NewProperties),
		ProvideDatasourceDialStrategy(dialector.NewMysql),
		ProvideDatasourceDialStrategy(dialector.NewPostgres),
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
