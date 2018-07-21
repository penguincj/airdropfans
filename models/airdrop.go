package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AirdropInfo struct {
	Id              int64
	Name            string    `orm:"size(500)"`
	Description     string    `orm:"size(2000)"`
	StepGuide       string    `orm:"size(2000)"`
	IncentivePlan   string    `orm:"size(2000)"` // 激励说明
	Cid             int64     //分类id
	Cname           string    //分类名称非数据库字段
	Addtime         time.Time `orm:"auto_now_add;type(datetime)"`
	EstimateValue   string
	TokensPerClaim  string
	MaxParticipants string
	TotalValue      string
	Logo            string
	Photo           string `orm:"size(500)"` //海报
	Iphoto          string
	StartTime       string
	EndTime         string
	Permanent       bool   // 永久有效
	Website         string `orm:"size(200)"`
	AirdropAddr     string `orm:"size(500)"`
	Platform        string `orm:"size(100)"`
	ReqTelegram     int64
	ReqTwitter      int64
	ReqMedium       int64
	ReqFacebook     int64
	ReqBitcointalk  int64
	ReqEmail        int64
	PreSaleDate     string
	BitcointalkAddr string
	BountyAddr      string
	IcoPrice        string
	Ticker          string
	TotalSupply     string
	WhitepaperAddr  string
	TwitterAddr     string
	FacebookAddr    string
	TelegramAddr    string
	MediumAddr      string
	GithubAddr      string
	Status          int64
	Isnew           int64
	Temperature     int64
	SeoTitle        string `orm:"size(500)"`  //seo标题
	SeoKeywords     string `orm:"size(500)"`  //seo关键字
	SeoDescription  string `orm:"size(1000)"` //seo说明
	Views           int64  //浏览量
	Vote            int64
	DisplayStatus   string
}

func (p *AirdropInfo) TableName() string {
	return "airdrop_info"
}

func (p *AirdropInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(p)
}

func (p *AirdropInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(p); err != nil {
		return err
	}
	return nil
}

func (p *AirdropInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *AirdropInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *AirdropInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(p); err != nil {
		return err
	}
	return nil
}
