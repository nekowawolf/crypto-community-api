package module

import (
	"context"
	"fmt"
	"github.com/nekowawolf/crypto-community-api/config"
	"github.com/nekowawolf/crypto-community-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOneDoc(collection string, doc interface{}) interface{} {
    insertResult, err := config.Database.Collection(collection).InsertOne(context.TODO(), doc)
    if err != nil {
        fmt.Printf("InsertOneDoc error: %v\n", err)
        return nil
    }
    return insertResult.InsertedID
}

func InsertCryptoCommunity(name, platforms, category, imgURL, linkURL string) interface{} {
    newCrypto := model.CryptoCommunity{
        ID:        primitive.NewObjectID(),
        Name:      name,
        Platforms: platforms,
        Category:  category,
        ImgURL:    imgURL,
        LinkURL:   linkURL,
    }

    insertedID := InsertOneDoc("cryptoCommunity", newCrypto)
    if insertedID == nil {
        fmt.Println("Failed to insert crypto community")
        return nil
    }

    fmt.Printf("Inserted new crypto community with ID: %v\n", insertedID)
    return insertedID
}

func GetAllCryptoCommunity() ([]model.CryptoCommunity, error) {
	collection := config.Database.Collection("cryptoCommunity")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %v", err)
	}
	defer cursor.Close(context.TODO())

	var communities []model.CryptoCommunity
	if err = cursor.All(context.TODO(), &communities); err != nil {
		return nil, fmt.Errorf("error decoding data: %v", err)
	}

	return communities, nil
}

func GetCryptoCommunityByID(id primitive.ObjectID) (*model.CryptoCommunity, error) {
	collection := config.Database.Collection("cryptoCommunity")
	filter := bson.M{"_id": id}

	var result model.CryptoCommunity
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetCryptoCommunityByName(name string) ([]model.CryptoCommunity, error) {
	collection := config.Database.Collection("cryptoCommunity")

	filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data by name: %v", err)
	}
	defer cursor.Close(context.TODO())

	var communities []model.CryptoCommunity
	if err = cursor.All(context.TODO(), &communities); err != nil {
		return nil, fmt.Errorf("error decoding data: %v", err)
	}

	return communities, nil
}

func UpdateCryptoCommunityByID(id primitive.ObjectID, updateData model.CryptoCommunity) (*model.CryptoCommunity, error) {
	collection := config.Database.Collection("cryptoCommunity")

	update := bson.M{
		"$set": bson.M{
			"name":       updateData.Name,
			"platforms":  updateData.Platforms,
			"category":   updateData.Category,
			"img_url":    updateData.ImgURL,
			"link_url":   updateData.LinkURL,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("error updating document: %v", err)
	}

	return &updateData, nil
}

func DeleteCryptoCommunityByID(id primitive.ObjectID) error {
    collection := config.Database.Collection("cryptoCommunity")
    filter := bson.M{"_id": id}

    result, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return fmt.Errorf("error deleting crypto community for ID %s: %s", id.Hex(), err.Error())
    }

    if result.DeletedCount == 0 {
        return fmt.Errorf("no crypto community found with ID %s", id.Hex())
    }

    return nil
}