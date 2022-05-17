package user

import (
	"context"

	"github.com/eifzed/ares/internal/model/user"
	tx "github.com/eifzed/ares/lib/database/mongodb/transaction"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2/bson"
)

func (db *userDB) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	user := &user.User{}
	err := db.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *userDB) InsertUser(ctx context.Context, user *user.User) error {
	session := tx.GetSessionFromContext(ctx)

	err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		var collection *mongo.Collection
		if session != nil {
			collection = tx.GetCollectionFromSession(session, "users")
		} else {
			collection = db.DB.Collection("users")
		}
		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}
		user.UserID = result.InsertedID.(primitive.ObjectID)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (db *userDB) CheckUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	count, err := db.DB.Collection("users").CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *userDB) UpdateUserRoles(ctx context.Context, userID primitive.ObjectID, newRoles []user.UserRole) error {
	// TODO: handle failed update when no data found, prevent upsert
	session := tx.GetSessionFromContext(ctx)

	err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		var collection *mongo.Collection
		if session != nil {
			collection = tx.GetCollectionFromSession(session, "users")
		} else {
			collection = db.DB.Collection("users")
		}
		_, err := collection.UpdateByID(ctx, userID, bson.M{"$addToSet": bson.M{"roles": bson.M{"$each": newRoles}}})

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
