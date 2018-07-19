package models

import (
    "github.com/astaxie/beego/orm"
)

type TokenPrice struct {
    Id              int64
    Name         string
	Pair         string
    Open         float64
    Close        float64
    CloseS        string
    FloatPercent string
	Fluctuation  string
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

