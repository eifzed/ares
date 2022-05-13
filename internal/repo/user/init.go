package user

import "go.mongodb.org/mongo-driver/mongo"

type userDB struct {
	DB *mongo.Database
}

type Options struct {
	DB *mongo.Database
}

func GetNewUserDB(option *Options) *userDB {
	if option == nil || option.DB == nil {
		return nil
	}
	return &userDB{
		DB: option.DB,
	}
}
