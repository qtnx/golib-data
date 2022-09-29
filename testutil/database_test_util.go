package golibdataTestUtil

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"gitlab.com/golibs-starter/golib/log"
	"gorm.io/gorm"
	"testing"
)

type DatabaseTestUtil struct {
	db *gorm.DB
}

func NewDatabaseTestUtil(db *gorm.DB) *DatabaseTestUtil {
	return &DatabaseTestUtil{db: db}
}

func (d *DatabaseTestUtil) TruncateTable(table string) {
	if err := d.db.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", table)).Error; err != nil {
		log.Fatalf("Could not truncate table [%s], err [%v]", table, err)
	} else {
		log.Infof("Truncated table [%s]", table)
	}
}

func (d *DatabaseTestUtil) Insert(model interface{}) {
	if err := d.db.Create(model).Error; err != nil {
		log.Fatalf("Could not create seed data, model: [%+v], err: [%v]", model, err)
	}
}

func (d *DatabaseTestUtil) CountWithoutQuery(table string) int64 {
	var count int64
	d.db.Table(table).Count(&count)
	return count
}

func (d *DatabaseTestUtil) CountWithQuery(table string, conditions map[string]interface{}) int64 {
	var count int64
	d.db.Table(table).Where(conditions).Count(&count)
	return count
}

// AssertDatabaseCount assert database has number of row without query
func (d *DatabaseTestUtil) AssertDatabaseCount(t *testing.T, table string, expected int64) {
	count := d.CountWithoutQuery(table)
	require.Equal(t, expected, count)
}

// AssertDatabaseHas assert database has more than a row with query params
func (d *DatabaseTestUtil) AssertDatabaseHas(t *testing.T, table string, conditions map[string]interface{}) {
	count := d.CountWithQuery(table, conditions)
	require.GreaterOrEqual(t, count, int64(1), "Record not found in database with query:", conditions)
}
