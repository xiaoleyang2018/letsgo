// Package ormx wraps ORM and provide some basic operation.
package ormx

import (
	"errors"
	"fmt"
	"os"

	"xorm.io/core"
	"github.com/go-xorm/xorm"
)

var orm *xorm.Engine

// list of DB errors
var (
	ErrNotExist = errors.New("not exist")
)

// DB configurations.
type DB struct {
	Type        string
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	TablePrefix string
	LogPath     string
}

// Init inits db info and setting db.
func Init(db DB, runMode string) (*xorm.Engine, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", db.User, db.Password, db.Host, db.Port, db.Name)

	println(url)
	var err error
	orm, err = xorm.NewEngine(db.Type, url)
	if err != nil {
		return nil, fmt.Errorf("init database error: %v", err)
	}

	println(orm)

	orm.SetLogLevel(core.LOG_DEBUG)
	orm.ShowSQL(true)
	orm.ShowExecTime(true)

	if runMode == "prod" {
		f, err := os.Create(db.LogPath)
		if err != nil {
			return nil, err
		}
		orm.SetLogger(xorm.NewSimpleLogger(f))
	}
	return orm, nil
}

// ORM returns initialized orm.
func ORM() *xorm.Engine {
	return orm
}

// GetByID return a obj by id.
func GetByID(id int64, obj interface{}) error {
	has, err := orm.Id(id).Get(obj)
	if err != nil {
		return err
	}
	if !has {
		return ErrNotExist
	}
	return nil
}

// SoftDeleteByID set record's is_deleted to true.
func SoftDeleteByID(id int, tab Table) error {
	tab.SetDeleted()
	_, err := orm.Id(id).Update(tab)
	return err
}

// DeleteByID delete a record from database.
func DeleteByID(id int, obj interface{}) error {
	_, err := orm.Id(id).Delete(obj)
	return err
}

// Create insert a record into database.
func Create(obj interface{}) error {
	_, err := orm.Insert(obj)
	return err
}
