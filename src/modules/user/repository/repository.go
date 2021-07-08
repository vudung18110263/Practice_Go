package repository

import (
	"github.com/vudung18110263/Practice_Go/src/modules/user/model"
)

type UserRepository interface {
	Insert(model.User) error
	// Update(string, *model.User) error
	// Delete(string) error
	//FindByID(string) (*model.UserPublic, error)
	//FindAll() (model.Users, error)
}
