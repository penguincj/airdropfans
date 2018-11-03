package controllers

import (
	"strconv"
	"strings"
	"time"

	"airdrop/models"
)

//"encoding/json"

//"time"

type IndexHandle struct {
	baseController
}

func (this *IndexHandle) Index() {
	var (
		info      models.AirdropInfo
		list      []*models.AirdropInfo
		hotList   []*models.AirdropInfo
		token     models.TokenPrice
		tokenList []*models.TokenPrice
	)

	token.Query().OrderBy("id").Limit(8, 0).All(&tokenList)
	info.Query().OrderBy("-id").Limit(30, 0).All(&list)
	info.Query().OrderBy("-temperature").Limit(30, 0).All(&hotList)

	for _, airdrop := range list {
		duringHours := time.Now().Sub(airdrop.Addtime).Hours()
		if duringHours <= 548 {
			airdrop.DisplayStatus = "new-flag"
		}
	}

	this.Data["market"] = list
	this.Data["tokenList"] = tokenList
	this.Data["list"] = list
	this.Data["hotList"] = hotList
	this.TplName = "_index.html"
}

func (this *IndexHandle) Detail() {
	var (
		id   int64
		info *models.AirdropInfo = new(models.AirdropInfo)
		err  error
	)

	this.Ctx.Output.Header("Cache-Control", "public")

	idstr := this.Ctx.Input.Param(":id")
	id, err = strconv.ParseInt(idstr, 10, 64)

	if err != nil || id <= 0 {
		this.Abort("404")
		return
	}

	info.Id = id
	err = info.Read()
	if err != nil || info.Status < 0 {
		this.Abort("404")
		return
	}

	info.Views++
	info.Temperature = info.Views*3 + info.Vote*5

	info.Update("views", "temperature")

	duringHours := time.Now().Sub(info.Addtime).Hours()
	if duringHours <= 548 {
		info.DisplayStatus = "new-flag"
	}

	info.Description = strings.Replace(info.Description, "&nbsp;", "", -1)
	this.Data["info"] = info
	this.TplName = "_detail.html"
	//this.TplName = "_airdrop_detial.html"
}

func (this *IndexHandle) New() {
	this.TplName = "_new.html"
}
