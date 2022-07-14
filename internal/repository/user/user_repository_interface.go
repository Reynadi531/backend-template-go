package user

import "backend-template-go/internal/entities/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	UserExists(email string) (bool, error)
	Delete(id string) error
}
