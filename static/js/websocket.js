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
            alert("�ɹ�������������")
        }
        ws.onclose = function () {
            if (ws) {
                ws.close();
                ws = null;
            }
            alert("���ӷ�����-�ر�1")
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
            alert("���ӷ�����-�ر�2")
        }
    } catch (e) {
        alert(e.message)
    }
}

//�ӷ�������ȡ��Ϣ���д���
function handleMessage(d) {
    data = JSON.parse(d)
    if (data.type === "market") {
        //��Ϣ
        app.markets.push(data.data)
    }
}