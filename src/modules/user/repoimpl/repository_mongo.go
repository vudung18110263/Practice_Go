package repoimpl

import (
	"context"
	"fmt"

	"github.com/vudung18110263/Practice_Go/src/modules/user/model"
	repo "github.com/vudung18110263/Practice_Go/src/modules/user/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepoMongo struct {
	db *mongo.Database
}
type MyError struct {
	a string
}

func (m *MyError) Error() string {
	return m.a
}
func ErrDo(b string) error {
	return &MyError{a: b}
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
func (r *UserRepoMongo) FindAll(offset, limit int) ([]model.User, error) {

	var users []model.User
	findOptions := options.Find()
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

	cur, err := r.db.Collection("user").Find(context.Background(), bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem model.User

		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		users = append(users, elem)
	}

	return users, nil
}

func (r *UserRepoMongo) Update(user model.User) error {
	objID, err := primitive.ObjectIDFromHex(user.Id.Hex())
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"password": user.Password,
		},
	}

	userUpdate, err := r.db.Collection("user").UpdateByID(context.Background(), objID, update)
	if err != nil {
		return err
	}
	fmt.Println(userUpdate)
	// var user_ model.User
	// err := userUpdate.Decode(user_)

	return nil
}
func (r *UserRepoMongo) Delete(id primitive.ObjectID) error {
	result, err := r.db.Collection("user").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return nil
}
func (r *UserRepoMongo) Find(id string) (*model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cur, err := r.db.Collection("user").Find(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}

	var result model.User
	for cur.Next(context.TODO()) {

		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}
func (r *UserRepoMongo) FindByName(name string) (*model.User, error) {
	cur, err := r.db.Collection("user").Find(context.Background(), bson.M{"name": name})
	if err != nil {
		return nil, err
	}

	var result model.User
	for cur.Next(context.TODO()) {

		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}
func (r *UserRepoMongo) IsUser(name, password string) (*model.User, bool) {
	cur, err := r.db.Collection("user").Find(context.Background(), bson.M{"name": name, "password": password})
	if err != nil {
		return nil, false
	}
	var result model.User
	//fmt.Println(cur)
	if cur.Next(context.TODO()) {
		err := cur.Decode(&result)
		//fmt.Println(result, "-", err)
		if err != nil {
			return nil, false
		}
		return &result, true
	}
	return nil, false
}
func (r *UserRepoMongo) IsUserNameExist(name string) error {
	count, err := r.db.Collection("user").CountDocuments(context.TODO(), bson.M{"name": name})
	if err != nil {
		return err
	}
	if count != 0 {
		err := ErrDo("user exist")
		return err
	}
	return nil
}
