package controllers

import (
	"airdrop/models"
	"encoding/json"
	"net/http"
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

var client []*websocket.Conn

func (this *WebSocketController) HandleWs() {

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

func broadcastMarket() {

	var (
		token     models.TokenPrice
		tokenList []*models.TokenPrice
	)

	token.Query().OrderBy("id").Limit(8, 0).All(&tokenList)

	data := TransferData{
		Type: "market",
		Data: marketMsg{
			Time: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	broadcastWebSocket(data)
}

func marketWs() {
	client = make([]*websocket.Conn, 0, 50)

	t := time.NewTimer(1 * time.Second)

	for {
		select {
		case <-t.C:
			t.Reset(1 * time.Minute)
			broadcastMarket()
		}
	}
}

func init() {
	go marketWs()
}
