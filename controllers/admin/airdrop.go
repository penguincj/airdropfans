package admin

import (
	"fmt"
	"strconv"
	"strings"

	"airdrop/models"
)

type AirdropHandle struct {
	baseController
}

//add page
func (this *AirdropHandle) Add() {
	var (
		info models.AirdropInfo
	)
	this.Data["AdminDir"] = this.admindir
	this.Data["info"] = info
	this.TplName = "admin/airdropadd.html"
}

//edit page
func (this *AirdropHandle) Edit() {
	var (
		id   int64
		info models.AirdropInfo
	)

	idStr := this.Ctx.Input.Param(":id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	if id <= 0 {
		this.showmsg("数据错误，返回重试...")
	}
	info.Id = id
	err := info.Read()
	if err != nil {
		this.showmsg("数据不存在...")
	}

	this.Data["info"] = info
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/airdropadd.html"
}

func (this *AirdropHandle) Delete() {
	var (
		id   int64
		info models.AirdropInfo
	)

	idStr := this.Ctx.Input.Param(":id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	if id <= 0 {
		this.showmsg("数据错误，返回重试...")
	}
	info.Id = id
	err := info.Delete()
	if err != nil {
		this.showmsg("删除数据失败...")
	}

	// TODO delete img

	this.List()
}

//list page
func (this *AirdropHandle) List() {

	var (
		page     int64
		pagesize int64 = 12
		offset   int64
		list     []*models.AirdropInfo
		airdrop  models.AirdropInfo
		keyword  string
		pager    string
	)
	keyword = strings.TrimSpace(this.GetString("keyword"))
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := airdrop.Query()
	if len(keyword) > 0 {
		query = query.Filter("title__icontains", keyword)
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PageList(pagesize, page, count, false, this.admindir+"airdrop/list")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["keyword"] = keyword
	this.Data["count"] = count
	this.TplName = "admin/airdroplist.html"
}

//save post
func (this *AirdropHandle) Save() {
	var (
		id   int64
		info models.AirdropInfo
		err  error
	)

	info.Name = strings.TrimSpace(this.GetString("name"))
	info.AirdropAddr = strings.TrimSpace(this.GetString("airdrop-website"))
	info.StartTime = strings.TrimSpace(this.GetString("start-time"))
	info.EndTime = strings.TrimSpace(this.GetString("end-time"))
	info.WhitepaperAddr = strings.TrimSpace(this.GetString("whitepaper-addr"))
	info.Photo = strings.TrimSpace(this.GetString("photo"))
	info.SeoTitle = strings.TrimSpace(this.GetString("title"))
	info.SeoKeywords = strings.TrimSpace(this.GetString("keywords"))
	info.SeoDescription = strings.TrimSpace(this.GetString("description"))
	info.Iphoto = strings.TrimSpace(this.GetString("iphoto"))
	info.Platform = strings.TrimSpace(this.GetString("platform"))

	id, _ = this.GetInt64("id")
	info.Cid, _ = this.GetInt64("cid")
	info.Status, _ = this.GetInt64("status")
	info.Isnew, _ = this.GetInt64("isnew")

	info.Website = strings.TrimSpace(this.GetString("website"))

	info.ReqTelegram, _ = this.GetInt64("reqTelegram")
	info.ReqTwitter, _ = this.GetInt64("reqTwitter")
	info.ReqMedium, _ = this.GetInt64("reqMedium")
	info.ReqFacebook, _ = this.GetInt64("reqFacebook")
	info.ReqBitcointalk, _ = this.GetInt64("reqBitcointalk")
	info.ReqEmail, _ = this.GetInt64("reqEmail")

	info.EstimateValue = strings.TrimSpace(this.GetString("estimate-value"))
	info.TokensPerClaim = strings.TrimSpace(this.GetString("tokens-per-claim"))
	info.MaxParticipants = strings.TrimSpace(this.GetString("max-participants"))
	info.TotalValue = strings.TrimSpace(this.GetString("total-value"))

	info.Description = strings.TrimSpace(this.GetString("airdrop-desc"))
	info.StepGuide = strings.TrimSpace(this.GetString("stepGuide"))
	info.IncentivePlan = strings.TrimSpace(this.GetString("incentivePlan"))

	if len(info.Name) == 0 || len(info.Photo) == 0 || len(info.AirdropAddr) == 0 || len(info.EstimateValue) == 0 || len(info.Platform) == 0 || len(info.StartTime) == 0 || len(info.TotalValue) == 0 || len(info.TokensPerClaim) == 0 {
		this.showmsg("带*号的为必须填写的内容...", this.admindir+"airdrop/edit/"+fmt.Sprintf("%d", id))
	}

	if id > 0 {
		info.Id = id
		err = info.Update("name", "description", "cid", "estimate_value", "tokens_per_claim", "max_participants", "logo", "photo", "iphoto", "start_time", "end_time", "website", "permanent", "airdrop_addr", "total_value", "platform", "step_guide", "incentive_plan", "req_telegram", "req_twitter", "req_medium", "req_facebook", "req_bitcointalk", "req_email", "pre_sale_date", "bitcointalk_addr", "bounty_addr", "whitepaper_addr", "twitter_addr", "facebook_addr", "telegram_addr", "medium_addr", "github_addr", "status", "isnew", "seo_title", "seo_keywords", "seo_description", "views", "vote", "temperature")
	} else {
		err = info.Insert()
	}
	if err != nil {
		this.showmsg("保存出错，错误信息：" + err.Error())
	} else {
		this.showmsg("数据保存成功...", "/d/"+strconv.FormatInt(id, 10))
		/*
			this.showmsg("数据保存成功...", "/d/" + strconv.FormatInt(int64, id)this.admindir+"airdrop/add", )
		*/
	}
}
