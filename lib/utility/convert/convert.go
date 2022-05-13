package convert

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjectID(data string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(data)
}
