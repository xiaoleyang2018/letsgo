package models

//Models 为包名（传参时候决定）

import (
	"fmt"
)

//Mapper 数据库的表
var UsersModel = &Users{}

//struct这段是xorm官方给的模版
type Users struct {
	Id    int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name  string `json:"name" xorm:"not null VARCHAR(32)"`
	Email string `json:"email" xorm:"not null VARCHAR(64)"`
}

func (m *Users) GetId() (val int) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *Users) GetName() (val string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *Users) GetEmail() (val string) {
	if m == nil {
		return
	}
	return m.Email
}

func (m *Users) String() string {
	return fmt.Sprintf("%#v", m)
}

func (m *Users) TableName() string {
	return "users"
}

//增
//删
//该
//查
