package db

import (
	"github.com/jjauzion/ws-backend/graph/model"
)

type DatabaseHandler interface {
	new() error
	Info() string
	Bootstrap() error
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id string) (model.User, error)
	CreateUser(user model.User) error
	CreateTask(task model.Task) error
}
