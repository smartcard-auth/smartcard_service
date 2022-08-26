package mongocontrol

import (
	"context"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoDriver struct {
	Ctx       context.Context
	MgoClient *mongo.Client
}

func (mgo *MgoDriver) GetDataOne(ctx context.Context, collectionName *mongo.Collection, query bson.M, opts *options.FindOneOptions) (*CardData, error) {
	var result CardData
	err := collectionName.FindOne(ctx, query, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Logger.Jrn.Printf("No objects in db which match current query = %v", query)
			return nil, err
		}
		log.Logger.Jrn.Printf("Error getting object = %v", err)
	}
	return &result, err
}

func (mgo *MgoDriver) GetDataMany(ctx context.Context, collectionName *mongo.Collection, query bson.M, opts *options.FindOptions) ([]*CardData, error) {
	var result []*CardData
	cursor, err := collectionName.Find(ctx, query, opts)
	if err != nil {
		log.Logger.Jrn.Printf("Error getting object array = %v", err)
		return nil, err
	}
	for cursor.Next(ctx) {
		var singleResult *CardData
		if err = cursor.Decode(&singleResult); err != nil {
			log.Logger.Jrn.Printf("Error getting cursor object = %v", err)
			return nil, err
		}
		result = append(result, singleResult)
	}
	return result, nil
}

func (mgo *MgoDriver) UpdateOne() {

}

func (mgo *MgoDriver) UpdateMany() {

}

func (mgo *MgoDriver) InsertOne() {

}

func (mgo *MgoDriver) InsertMany() {

}
