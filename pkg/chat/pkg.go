package chat

import (
	"github.com/gorilla/websocket"
	"github.com/jasonsoft/chatroom/pkg/chat/service"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// service
	_svc service.ChatServicer

	// handler
	_chatAPIHandler *ChatHandlder
)

func Initialize() {
	_svc = service.NewChatService()

	_chatAPIHandler = NewChatHandler(_svc)
}
