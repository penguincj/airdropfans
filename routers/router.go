package routers

import (
	"github.com/astaxie/beego"

	"airdrop/controllers"
	"airdrop/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/detail", &controllers.IndexHandle{}, "*:OldDetail")
	beego.Router("/d/:id:int/", &controllers.IndexHandle{}, "*:Detail")
	beego.Router("/new", &controllers.IndexHandle{}, "*:New")

	//beego.Router("/meeting", &controllers.MeetingHandle{}, "*:Meeting")

	// WebSocket.
	beego.Router("/market", &controllers.WebSocketController{}, "*:Market")
	beego.Router("/temp/d/:id:int/", &controllers.WebSocketController{}, "*:Temperature")
	//beego.Router("/temperature", &controllers.WebSocketController{}, "*:Temperature")
	//beego.Router("/ws/test", &controllers.WebSocketController{}, "get:HandleWs")

	///admin
	admindir := beego.AppConfig.String("admindir")
	if len(admindir) == 0 {
		admindir = "admin"
	}
	ns := beego.NewNamespace(admindir,
		beego.NSRouter("/", &admin.LoginHandle{}, "*:Login"),
		beego.NSRouter("/logout", &admin.LoginHandle{}, "*:Logout"),
		beego.NSRouter("/main", &admin.IndexHandle{}, "*:Main"),
		beego.NSRouter("/left", &admin.IndexHandle{}, "*:Left"),
		beego.NSRouter("/right", &admin.IndexHandle{}, "*:Right"),

		beego.NSRouter("airdrop/add", &admin.AirdropHandle{}, "*:Add"),
		beego.NSRouter("airdrop/edit/:id:int/", &admin.AirdropHandle{}, "*:Edit"),
		beego.NSRouter("airdrop/delete/:id:int/", &admin.AirdropHandle{}, "*:Delete"),
		beego.NSRouter("airdrop/save", &admin.AirdropHandle{}, "post:Save"),
		beego.NSRouter("airdrop/list", &admin.AirdropHandle{}, "*:List"),
		beego.NSRouter("airdrop/list/:page:int/", &admin.AirdropHandle{}, "*:List"),

		beego.NSRouter("upload/add", &admin.UploadHandle{}, "*:UploadPage"),
		beego.NSRouter("upload/uploadsave", &admin.UploadHandle{}, "*:UploadSave"),

		beego.NSRouter("user/changepassword", &admin.UserHandle{}, "*:ChangePass"),
		beego.NSRouter("user/savepass", &admin.UserHandle{}, "*:SavePass"),
	)
	beego.AddNamespace(ns)
}
