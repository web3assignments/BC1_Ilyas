package chat

import (
	"net/http"
)

func JoinEndpoint(chatRoom *ChatRoom) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO
	}
}

func SendEndpoint(chatRoom *ChatRoom) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO
	}
}
