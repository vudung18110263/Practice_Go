package repoimpl

import (
	"context"

	"github.com/vudung18110263/Practice_Go/src/modules/user/model"
	repo "github.com/vudung18110263/Practice_Go/src/modules/user/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoMongo struct {
	db *mongo.Database
}

func NewUserRepoMongo(db *mongo.Database) repo.UserRepository {
	return &UserRepoMongo{
		db: db,
	}
}

func (r *UserRepoMongo) Insert(user model.User) error {
	bbytes, _ := bson.Marshal(user)
	_, err := r.db.Collection("user").InsertOne(context.Background(), bbytes)
	return err
}
