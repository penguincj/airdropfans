let app = undefined
let ws = undefined

var tempapp = new Vue({
    el: '#tempapp',
    data: {
        temperature: 0
    },

    methods: {
        increase_temp: function (event) {
            this.temperature = this.temperature + 1
            sendEvent("add-temperature")
        },
        decrease_temp: function (event) {
            if (this.temperature > 0) {
                this.temperature = this.temperature - 1
                sendEvent("del-temperature")
            }
        },
    },
});

ws = new WebSocket('ws://airdropfans.com/' + 'temp' + window.location.pathname)
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

//�ӷ�������ȡ��Ϣ���д���
function handleMessage(d) {
    data = JSON.parse(d)
    if (data.type === "temperature") {
        //��Ϣ
        tempapp.temperature = data.data
        console.log(data.data)
    }
}

function sendEvent(data) {
    let msg = {
        type: "temperature",
        data: data,
    }
    console.log(msg)
    ws.send(JSON.stringify(msg));
}