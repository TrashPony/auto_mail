var sock;
function ConnectLobby() {
    sock = new WebSocket("ws://" + window.location.host + "/ws");
    console.log("Websocket - status: " + sock.readyState);

    sock.onopen = function(msg) {
        console.log("CONNECTION opened..." + this.readyState);
        Config();
    };
    sock.onmessage = function(msg) {
        console.log("message: " + msg.data);
        ReaderLobby(msg.data);
    };
    sock.onerror = function(msg) {
        console.log("Error occured sending..." + msg.data);
    };
    sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
    };
}

function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "Add" || event === "Del"){

        var container = document.getElementById('List');

        while (container.firstChild) {
            container.removeChild(container.firstChild);
        }
        Config();
    }

    if (event === "Send") {
        var configLine;
        var value;

        if(JSON.parse(jsonMessage).path === "0"){
            configLine = document.getElementById(JSON.parse(jsonMessage).email);
            configLine.className = "ConfigLine Passed";
            value = configLine.innerHTML;
            configLine.innerHTML = value + " PASSED";
        } else {
            if(JSON.parse(jsonMessage).path === "1"){
                configLine = document.getElementById(JSON.parse(jsonMessage).email);
                configLine.className = "ConfigLine Error";
                value = configLine.innerHTML;
                configLine.innerHTML = value + " ERR SEND";
            }
            if(JSON.parse(jsonMessage).path === "2"){
                configLine = document.getElementById(JSON.parse(jsonMessage).email);
                configLine.className = "ConfigLine Error";
                value = configLine.innerHTML;
                configLine.innerHTML = value + " ERR DIR";
            }
        }

        var button = document.getElementById("SendButton");

        button.value = "Отправить!";
        button.onclick = Send;
    }

    if (event === "Config"){

        var div = document.createElement('div');
        div.className = "ConfigLine";
        div.id = JSON.parse(jsonMessage).email;
        div.onclick = function () {
            Delete(this.id)
        };
        div.innerHTML = JSON.parse(jsonMessage).email + "   :   " + JSON.parse(jsonMessage).path;
        var parentElem = document.getElementById("List");
        parentElem.appendChild(div);
    }
}

function Delete(id) {
    var div = document.createElement('div');
    div.className = "ok";
    div.innerHTML = "Удалить";
    div.onclick = function () {
        Del(id);
    };
    var parentElem = document.getElementById(id);
    parentElem.appendChild(div);
}

function Add() {
    var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;

    var email = document.getElementById("Email").value;
    var passed = regex.test(email);

    if (passed) {
        sock.send(JSON.stringify({
            event: "Add",
            email: email,
            path: document.getElementById("Path").value
        }));
    } else {
        var div = document.createElement('div');
        div.className = "Error";
        div.innerHTML = "Неправильно введен Email";
        var parentElem = document.getElementById("main");
        parentElem.appendChild(div);
    }

    document.getElementById("Email").value = "";
    document.getElementById("Path").value = "";
}

function Del(id) {
    sock.send(JSON.stringify({
        event:"Del",
        email:id
    }));
}

function Config() {
    sock.send(JSON.stringify({
        event:"Config"
    }));
}

function Send() {
    var container = document.getElementById('List');

    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }

    Config();

    var button = document.getElementById("SendButton");

    button.value = "Ожидайте";
    button.onclick = null;

    sock.send(JSON.stringify({
        event:"Send"
    }));
}