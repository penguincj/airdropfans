package controllers

type DappHandle struct {
	baseController
}

func (this *DappHandle) Casino() {
	this.TplName = "_casino.html"
}

func (this *DappHandle) Love() {
	this.TplName = "_love.html"
}

func (this *DappHandle) Insurance() {
	this.TplName = "_insurance.html"
}

func (this *DappHandle) Signin() {
	this.TplName = "_signin.html"
}

