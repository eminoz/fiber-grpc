package repo

import (
	"context"

	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	"gitbub.com/eminoz/graceful-fiber/server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	InsertUser(ctx context.Context, user *api.User) (*api.ResUser, error)
	GetUserById(id string) *api.ResUser
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
