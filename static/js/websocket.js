let app = undefined
let ws = undefined
appinit()
wsinit()
function appinit() {
    app = new Vue({
        el: '#market',
        data: {
            markets: [],
        },
    })
}

function wsinit() {
    ws = new WebSocket("ws://' + window.location.host + '/ws/")
    /*
     * ws = new WebSocket("ws://localhost:8080/chat")
    */
    try {
        ws.onopen = function () {
            alert("成功连接至服务器")
        }
        ws.onclose = function () {
            if (ws) {
                ws.close();
                ws = null;
            }
            alert("连接服务器-关闭1")
        }
        ws.onmessage = function (ret) {
            console.log(ret.data);
            handleMessage(ret.data)
        }
        ws.onerror = function () {
            if (ws) {
                ws.close()
                ws = null
            }
            alert("连接服务器-关闭2")
        }
    } catch (e) {
        alert(e.message)
    }
}

//从服务器获取消息进行处理
function handleMessage(d) {
    data = JSON.parse(d)
    if (data.type === "market") {
        //消息
        app.markets.push(data.data)
    }
}
