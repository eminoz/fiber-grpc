package repo

import (
	"context"
	"encoding/json"

	"gitbub.com/eminoz/graceful-fiber/client/model"
	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	"gitbub.com/eminoz/graceful-fiber/server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	InsertUser(ctx context.Context, user *api.User) (*api.ResUser, error)
	GetUserById(id string) *api.ResUser
	DeleteUserById(id string) (bool, error)
	UpdateUserById(user *api.UpdateUser) (*api.ResUser, error)
}
type userRepo struct {
	DB *mongo.Database
	Cl *mongo.Collection
}

func NewUserController() UserRepo {
	var database = db.GetDatabase()
	return &userRepo{
		DB: database,
		Cl: database.Collection("user"),
	}
}

func (u userRepo) InsertUser(ctx context.Context, user *api.User) (*api.ResUser, error) {

	r, err := u.Cl.InsertOne(ctx, &user)

	if err != nil {
		panic(" error not nil in insert")
	}
	id := r.InsertedID
	var usr *api.ResUser

	filter := bson.D{{Key: "_id", Value: id}}
	u.Cl.FindOne(ctx, &filter).Decode(&usr)

	return usr, nil

}
func (u userRepo) GetUserById(id string) *api.ResUser {
	userID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userID}}
	var usr *api.ResUser
	u.Cl.FindOne(context.Background(), &filter).Decode(&usr)

	return usr
}
func (u userRepo) DeleteUserById(id string) (bool, error) {
	userId, _ := primitive.ObjectIDFromHex(id)
	_, err := u.Cl.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: userId}})
	if err != nil {
		return false, err
	}
	return true, nil
}
func (u userRepo) UpdateUserById(user *api.UpdateUser) (*api.ResUser, error) {
	id, _ := primitive.ObjectIDFromHex(user.Id)
	jsonBytes, _ := json.Marshal(user)
	var modeluser model.User
	json.Unmarshal(jsonBytes, &modeluser)

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: modeluser}}
	updated, err := u.Cl.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if updated.ModifiedCount == 1 {
		var usr *api.ResUser
		u.Cl.FindOne(context.Background(), &filter).Decode(&usr)
		return usr, nil
	}
	return &api.ResUser{}, nil

}
