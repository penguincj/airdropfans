package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

type ContactSubmitController struct {
    baseController
}

func (c *ContactSubmitController) Get() {
	c.Data["IsAbout"] = true
	c.Data["IsBlack"] = true
	name := c.GetString("user")
	beego.Info("get submit")
	beego.Info(name)
	fmt.Printf("get name %s\n", name)
	c.TplName = "submit-success.html"
}

func (c *ContactSubmitController) Post() {
	c.Data["IsAbout"] = true
	c.Data["IsBlack"] = true
	beego.Info("post submit")
	message := c.GetString("message")
	//beego.Info(name, email, company, phone_number, message)

	c.TplName = "submit-success.html"

	body := "相亲收到访客留言，详情如下： <br><br>" +
		"留言：" + message + "<br><br>"

	m := gomail.NewMessage()
	m.SetHeader("From", "liuchuanjun1990@163.com")
	//m.SetHeader("To", "liuchuanjun1990@163.com", "cora@example.com")
	m.SetHeader("To", "penguincj@163.com")
	//m.SetAddressHeader("Cc", "liuchuanjun@connetos.com", "刘传军")
	m.SetHeader("Subject", "相亲访客留言")
	m.SetBody("text/html", body)
	//m.Attach("static/img/city.jpg")

	d := gomail.NewDialer("smtp.163.com", 25, "liuchuanjun1990@163.com", "5qishiaini")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}    

