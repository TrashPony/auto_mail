package webSocket

import (
	"log"
	"github.com/gorilla/websocket"
	"../sendEmailService"

	"strconv"
	"time"
)

// пайп доп. читать в документации

func Reader(ws *websocket.Conn)  {
	for {
		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			break
		}
		if msg.Event == "Send"{
			result := sendEmailService.SendEmailService()

			for client := range users {
				if ws == client.ws {
					for email, res := range result {
						resp := Response{Event: msg.Event, Email: email, Path: strconv.Itoa(res), ID: client.id}
						Pipe <- resp
						time.Sleep(300 * time.Millisecond)
					}
				}
			}

		}

		if msg.Event == "Config"{
			for client := range users {
				if ws == client.ws {
					emailAndPath, err := sendEmailService.ParseConfigFile()
					if err != nil {
						println(err)
					}
					for email, path := range emailAndPath {
						resp := Response{Event: msg.Event, Email: email, Path: path, ID: client.id}
						Pipe <- resp
					}
				}
			}
		}
		if msg.Event == "Del"{
			success := sendEmailService.Delete(msg.Email)
			if success {
				for client := range users {
					if ws == client.ws {
						resp := Response{Event: msg.Event, Email: msg.Email, ID: client.id}
						Pipe <- resp
					}
				}
			}
		}
		if msg.Event == "Add"{
			success := sendEmailService.Write(msg.Email, msg.Path)
			if success {
				for client := range users {
					if ws == client.ws {
						resp := Response{Event: msg.Event, ID: client.id}
						Pipe <- resp
					}
				}
			}
		}
	}
}

func ReposeSender() {
	for {
		resp := <-Pipe
		for client := range users {
			if client.id == resp.ID {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					client.ws.Close()
					delete(users, client)
				}
			}
		}
	}
}

type Message struct {
	Event    string `json:"event"`
	Email    string `json:"email"`
	Path     string `json:"path"`
	ID 		 string `json:"id"`
}

type Response struct {
	Event  string `json:"event"`
	Email  string `json:"email"`
	Path   string `json:"path"`
	ID 	   string `json:"id"`
}