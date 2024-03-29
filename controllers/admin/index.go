package admin

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"airdrop/models"
)

type IndexHandle struct {
	baseController
}

type LoginHandle struct {
	baseController
}

func (this *IndexHandle) Main() {
	this.TplName = "admin/main.html"
}

func (this *IndexHandle) Left() {
	this.Data["username"] = this.username
	this.TplName = "admin/left.html"
}

func (this *IndexHandle) Right() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH
	this.Data["beegover"] = beego.VERSION
	this.Data["clientip"] = this.getClientIp()

	this.Data["systemver"] = beego.AppConfig.String("systemver")
	this.Data["developer"] = beego.AppConfig.String("developer")
	this.Data["servertime"] = this.FormatTime(time.Now(), "YYYY年MM月DD日 HH:mm:ss")

	var airdrop models.AirdropInfo
	count, _ := airdrop.Query().Count()
	this.Data["airdropcount"] = count
	this.TplName = "admin/right.html"
}

//login post
func (this *LoginHandle) Login() {

	if this.IsPost() {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		var info string
		if len(username) == 0 || len(password) == 0 {
			info = "请填写登录帐号和密码..."
		} else {
			var user models.UserInfo
			user.Username = username
			if user.Read("username") != nil || user.Password != models.Md5(password) {
				//if user.Read("username") != nil || user.Password != password {
				//fmt.Println("cj get password %s md5 %s", user.Password, password)
				info = "帐号或密码错误..."
			} else {
				user.Logintimes += 1
				user.Lastloginip = this.getClientIp()
				user.Lastlogintime = this.getTime()
				user.Update()
				authKey := models.Md5("samsong|" + user.Password)
				this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authKey)
				this.Redirect(this.admindir+"main", 302)
			}
		}
		this.Data["username"] = username
		this.Data["info"] = info
	}
	this.TplName = "admin/login.html"
}

//logout
func (this *LoginHandle) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Ctx.WriteString("<script>top.location.href='" + this.admindir + "'</script>")
}
