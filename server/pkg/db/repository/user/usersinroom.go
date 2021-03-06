package user

import (
	"context"

	db2 "github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public/model"
)

type UsersInRoomRepository interface {
	CreateUsersInRoom(ctx context.Context, roomID int, userName, userID string) (id int, err error)
	ReadUsersInRoom(ctx context.Context, id int) (*model.Usersinroom, error)
	UpdateUsersInRoom(ctx context.Context, room *model.Usersinroom) error
	DeleteUsersInRoom(ctx context.Context, id int) error

	ListUsersInRoom(
		ctx context.Context,
		list *db2.ListOptions, criteria *db2.UsersInRoomCriteria) ([]*model.Usersinroom, error)
}
