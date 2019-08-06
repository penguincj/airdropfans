package controllers

type LoveHandle struct {
	baseController
}

func (this *LoveHandle) LoveDetail() {
	this.TplName = "_love_detail.html"
}

func (this *LoveHandle) Question() {
	this.TplName = "_me_question.html"
}
