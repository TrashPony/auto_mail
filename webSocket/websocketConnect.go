package webSocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"math/rand"
	"strconv"
)

var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket
var users = make(map[Clients]bool) // тут будут храниться наши подключения
var Pipe = make(chan Response)


func ReadSocket(w http.ResponseWriter, r *http.Request, pool string) {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	id := rand.Intn(1000)
	users[Clients{ws, strconv.Itoa(id)}] = true // Регистрируем нового Клиента
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	Reader(ws)
}

type Clients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	id string
}

