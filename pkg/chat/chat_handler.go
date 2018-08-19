package chat

import (
	"github.com/jasonsoft/chatroom/pkg/chat/model"
	"github.com/jasonsoft/chatroom/pkg/chat/service"
	"github.com/jasonsoft/napnap"
)

type ChatHandlder struct {
	svc service.ChatServicer
}

func NewChatHandler(svc service.ChatServicer) *ChatHandlder {
	return &ChatHandlder{
		svc: svc,
	}
}

func NewChatRouter() *napnap.Router {
	router := napnap.NewRouter()
	router.Get("/v1/rooms/default/join", _chatAPIHandler.roomJoinEndpoint)
	return router
}

func (h *ChatHandlder) roomJoinEndpoint(c *napnap.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	member := model.Member{
		Name:   "Jason",
		Avatar: "000",
	}
	client := model.NewClient(member, conn)

	room, err := h.svc.RoomGet("default")
	if err != nil {
		panic(err)
	}

	err = room.Join(client)
	if err != nil {
		panic(err)
	}
}
