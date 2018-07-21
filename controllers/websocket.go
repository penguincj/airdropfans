package controllers

import (
	"airdrop/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

//TransferData 数据传输结构体
type TransferData struct {
	Type string      `json:"type"` //数据类别 login表示登录  msg表消息 users用户列表
	Data interface{} `json:"data"` //数据
}

//Message 消息的结构体
type marketMsg struct {
	Markets []models.TokenPrice `json:"markets"`
	Time    string              `json:"time"` //消息发出的时间
}

var (
	client = make([]*websocket.Conn, 0, 50)
	event  = make(chan models.Event, 10)
)

func newEvent(ep models.EventType, id int64, conn *websocket.Conn, data TransferData) models.Event {
	return models.Event{ep, id, conn, int(time.Now().Unix()), data}
}

func (this *WebSocketController) Get() {
	fmt.Println("cj get")
	// Safe check.
	this.TplName = "_index.html"
	this.Data["IsWebSocket"] = true
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	client = append(client, ws)

}

func (this *WebSocketController) Market() {

	fmt.Println("cj HandleWs")

	this.TplName = "_index.html"

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	client = append(client, ws)

	data := getMarket()
	sendWebSocket(ws, data)

}

func sendWebSocket(ws *websocket.Conn, d TransferData) {
	data, err := json.Marshal(d)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	ws.WriteMessage(websocket.TextMessage, data)
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(d TransferData) {
	data, err := json.Marshal(d)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for i, ws := range client {
		if ws.WriteMessage(websocket.TextMessage, data) != nil {
			client = append(client[:i], client[i+1:]...)
		}
	}
}

func getMarket() TransferData {

	var (
		token     models.TokenPrice
		tokenList []*models.TokenPrice
	)

	token.Query().OrderBy("id").Limit(8, 0).All(&tokenList)

	var msg marketMsg
	msg.Time = time.Now().Format("2006-01-02 15:04:05")
	msg.Markets = make([]models.TokenPrice, 0, 8)

	for _, t := range tokenList {
		msg.Markets = append(msg.Markets, *t)
	}

	data := TransferData{
		Type: "market",
		Data: msg,
	}

	return data
}

func sendTemperature(ws *websocket.Conn, id int64) {
	var (
		info *models.AirdropInfo = new(models.AirdropInfo)
		err  error
	)

	info.Id = id
	err = info.Read()
	if err != nil || info.Status < 0 {
		return
	}

	data := TransferData{
		Type: "temperature",
		Data: info.Temperature,
	}

	sendWebSocket(ws, data)
}

func (this *WebSocketController) Temperature() {

	fmt.Println("cj tem")

	this.TplName = "_detail.html"

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idstr, 10, 64)

	if err != nil || id <= 0 {
		this.Abort("404")
		return
	}

	sendTemperature(ws, id)

	for {
		_, p, err := ws.ReadMessage()
		fmt.Println("receive cj")
		if err != nil {
			return
		}

		var data TransferData
		err = json.Unmarshal(p, &data)
		if err == nil {
			event <- newEvent(models.EVENT_TEMPERATURE, id, ws, data)
		}
	}
}

func handleTemperature(event models.Event) {
	var (
		info *models.AirdropInfo = new(models.AirdropInfo)
		err  error
	)

	info.Id = event.Id
	err = info.Read()
	if err != nil || info.Status < 0 {
		return
	}
	data := event.Data.(TransferData)
	str := data.Data.(string)
	if str == "add-temperature" {
		info.Vote++
		info.Temperature = info.Views*3 + info.Vote*5
		//info.Temperature = info.Temperature + 7
		err = info.Update("name", "description", "cid", "estimate_value", "tokens_per_claim", "max_participants", "logo", "photo", "iphoto", "start_time", "end_time", "website", "permanent", "airdrop_addr", "total_value", "platform", "step_guide", "incentive_plan", "req_telegram", "req_twitter", "req_medium", "req_facebook", "req_bitcointalk", "req_email", "pre_sale_date", "bitcointalk_addr", "bounty_addr", "whitepaper_addr", "twitter_addr", "facebook_addr", "telegram_addr", "medium_addr", "github_addr", "status", "isnew", "seo_title", "seo_keywords", "seo_description", "views", "vote", "temperature")
		if err != nil {
			fmt.Printf("update airdrop %d failed due to temperature", info.Id)
		}
	}
}

func handleEvent(event models.Event) {
	switch event.Type {
	case models.EVENT_TEMPERATURE:
		handleTemperature(event)
	}

}

func marketWs() {

	t := time.NewTimer(1 * time.Second)

	for {
		select {
		case <-t.C:
			t.Reset(1 * time.Minute)
			data := getMarket()
			broadcastWebSocket(data)
		case e := <-event:
			handleEvent(e)
		}
	}
}

func init() {
	go marketWs()
}
