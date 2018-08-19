package service

import (
	"errors"
	"sync"

	"github.com/jasonsoft/chatroom/pkg/chat/model"
	"github.com/jasonsoft/log"
)

type ChatServicer interface {
	RoomAdd(room *model.Room) error
	RoomGet(name string) (*model.Room, error)
}

type ChatService struct {
	rooms sync.Map
}

func NewChatService() ChatServicer {
	log.Debug("New service")
	room := model.NewRoom("default")
	svc := &ChatService{}
	if room != nil {
		svc.RoomAdd(room)
	}

	return svc
}

func (svc *ChatService) RoomAdd(room *model.Room) error {
	svc.rooms.Store(room.Name, room)
	log.Debug("roomadd:")
	return nil
}

func (svc *ChatService) RoomGet(name string) (*model.Room, error) {
	if target, ok := svc.rooms.Load(name); ok {
		log.Debug("find room")
		if room, ok := target.(*model.Room); ok {
			return room, nil
		}
	}

	return nil, errors.New("not found la")
}
