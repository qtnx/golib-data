package golibdataTestUtil

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"testing"
)

var orm *gorm.DB

func EnableDatabaseTestUtil() fx.Option {
	return fx.Invoke(func(db *gorm.DB) {
		orm = db
	})
}

func truncateTable(table string) {
	if err := orm.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", table)).Error; err != nil {
		log.Fatalf("Could not truncate table [%s], err [%v]", table, err)
	} else {
		log.Infof("Truncated table [%s]", table)
	}
}

func TruncateTables(tables ...string) {
	for _, table := range tables {
		truncateTable(table)
	}
}

func Insert(model interface{}) {
	if err := orm.Create(model).Error; err != nil {
		log.Fatalf("Could not create seed data, model: [%+v], err: [%v]", model, err)
	}
}

func CountWithoutQuery(table string) int64 {
	var count int64
	orm.Table(table).Count(&count)
	return count
}

func CountWithQuery(table string, conditions map[string]interface{}) int64 {
	var count int64
	orm.Table(table).Where(conditions).Count(&count)
	return count
}

// AssertDatabaseCount assert database has number of row without query
func AssertDatabaseCount(t *testing.T, table string, expected int64) {
	count := CountWithoutQuery(table)
	require.Equal(t, expected, count)
}

// AssertDatabaseHas assert database has more than a row with query params
func AssertDatabaseHas(t *testing.T, table string, conditions map[string]interface{}) {
	count := CountWithQuery(table, conditions)
	require.GreaterOrEqual(t, count, int64(1), "Record not found in database with query:", conditions)
}

// DB return gorm instance
func DB() *gorm.DB {
	return orm
}
