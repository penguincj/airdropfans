package main

import (
	_ "airdrop/routers"
	"github.com/astaxie/beego"
	//"airdrop/market"

)

func main() {
	//go market.RunMarket()
	beego.Run()
}

