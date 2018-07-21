var app = new Vue({
  el: '#app',
  data: {
    sites: [
      { name: 'cj' },
      { name: 'Google' },
      { name: 'Taobao' }
    ]
  }
});


var marketapp = new Vue({
    el: '#marketapp',
    data: {
        markets: [],
    },
});

ws = new WebSocket("ws://localhost:80/market")
try {
    ws.onopen = function () {
    }
    ws.onclose = function () {
        if (ws) {
            ws.close();
            ws = null;
        }
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
    }
} catch (e) {
    alert(e.message)
}

//从服务器获取消息进行处理
function handleMessage(d) {
    data = JSON.parse(d)
    if (data.type === "market") {
        //消息
        marketapp.markets = []
        console.log(data.data.markets[0])
        for (var i = 0; i < data.data.markets.length; i++) {
            marketapp.markets.push(data.data.markets[i])
        }
    }
}

