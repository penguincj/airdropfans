package models

import (
	"github.com/astaxie/beego/orm"
)

type TokenPrice struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Pair         string  `json:"pair"`
	Open         float64 `json:"open"`
	Close        float64 `json:"close"`
	CloseS       string  `json:"closes"`
	FloatPercent string  `json:"floatpercent"`
	Fluctuation  string  `json:"fluctuation"`
}

func (t *TokenPrice) TableName() string {
	return "token_price"
}

func (t *TokenPrice) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(t)
}

func (t *TokenPrice) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

func (t *TokenPrice) Read(fields ...string) error {
	if err := orm.NewOrm().Read(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *TokenPrice) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *TokenPrice) Delete() error {
	if _, err := orm.NewOrm().Delete(t); err != nil {
		return err
	}
	return nil
}
