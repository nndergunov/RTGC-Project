package app

import (
	"errors"
	"fmt"

	"github.com/nndergunov/RTGC-Project/server/pkg/app/allrooms"
	"github.com/nndergunov/RTGC-Project/server/pkg/domain"
)

// static errors.
var errUnknownAction = errors.New("action not supported")

type Router struct {
	roomList *allrooms.AllRooms
}

func (r *Router) Init(rl *allrooms.AllRooms) {
	r.roomList = rl
}

// ActionHandler sends request to the correct handler.
func (r *Router) ActionHandler(id, action, roomName, userName, text string) (*domain.Message, error) {
	switch action {
	case "register":
		return nil, nil

	case "join":
		err := r.roomList.Join(id, userName, roomName)
		if err != nil {
			return nil, fmt.Errorf("join handler: %w", err)
		}

		joinMessage := fmt.Sprintf("user '%s' joined the room '%s'", userName, roomName)

		return r.ActionHandler("SERVER", "send", roomName, "SERVER", joinMessage)

	case "leave":
		userName, err := r.roomList.Leave(id, roomName, text)
		if err != nil {
			return nil, fmt.Errorf("leave handler: %w", err)
		}

		leaveMessage := fmt.Sprintf("user '%s' left the room '%s'", userName, roomName)

		return r.ActionHandler("SERVER", "send", roomName, "SERVER", leaveMessage)

	case "send":
		msg, err := r.roomList.Send(id, roomName, text)
		if err != nil {
			return nil, fmt.Errorf("send handler: %w", err)
		}

		return msg, nil

	default:
		return nil, fmt.Errorf("%w : '%s'", errUnknownAction, action)
	}
}
