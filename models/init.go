package models

import (
	"github.com/go-xorm/xorm"

	"github.com/qinhao/letsgo/ormx"
)

var orm *xorm.Engine

// Init inits models.
func Init(db ormx.DB, runMode string) error {
	var err error
	orm, err = ormx.Init(db, runMode)
	if err != nil {
		return err
	}
	return nil
}

