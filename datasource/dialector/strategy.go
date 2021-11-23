package dialector

import "gorm.io/gorm"

type Strategy interface {
	Driver() string

	Open(cf *Config) (gorm.Dialector, error)
}
