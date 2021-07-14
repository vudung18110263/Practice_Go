package repository

import (
	"vudung-practice-go/src/modules/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Insert(model.User) error
	// Update(string, *model.User) error
	// Delete(string) error
	//FindByID(string) (*model.UserPublic, error)
	FindAll(offset, limit int) ([]model.User, error)
	Update(model.User) error
	Delete(primitive.ObjectID) error
	Find(string) (*model.User, error)
	IsUser(name, password string) (*model.User, bool)
	IsUserNameExist(name string) error
	FindByName(name string) (*model.User, error)
}
