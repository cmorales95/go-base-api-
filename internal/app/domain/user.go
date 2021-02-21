package domain

import (
	"github.com/google/uuid"
	"gitlab.com/crisdavidmm95.dev/aug7/internal/app/model"
)

type User interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(ID uuid.UUID) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetAll() (model.Users, error)
}
