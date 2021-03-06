package room

import (
	"context"

	db2 "github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/server/pkg/db/internal/rtgc/public/model"
)

type RoomsRepository interface {
	CreateRooms(ctx context.Context, name string) (id int, err error)
	ReadRooms(ctx context.Context, id int) (*model.Rooms, error)
	UpdateRooms(ctx context.Context, room *model.Rooms) error
	DeleteRooms(ctx context.Context, id int) error

	ListRooms(ctx context.Context, list *db2.ListOptions, criteria *db2.RoomsCriteria) ([]*model.Rooms, error)
}
