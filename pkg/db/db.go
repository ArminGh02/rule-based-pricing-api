package db

import (
	"context"
	"log"
	"pricingapi/pkg/model"
	"pricingapi/pkg/model/usertype"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func New(uri string) *DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panicln(err)
	}

	coll := client.Database("pricing_api").Collection("rules")

	defer log.Println("Connected to MongoDB.")

	return &DB{
		client: client,
		coll:   coll,
	}
}

func (db *DB) InsertRule(r *model.Rule) {
	_, err := db.coll.InsertOne(context.TODO(), r)
	if err != nil {
		log.Panicln(err)
	}
}

func (db *DB) FilterRulesByUserType(userType usertype.UserType) []*model.Rule {
	cur, err := db.coll.Find(context.TODO(), bson.M{"conditions": bson.M{"$in": bson.M{"user_type": userType}}})
	if err != nil {
		log.Panicln(err)
	}
	var rules []*model.Rule
	for cur.Next(context.TODO()) {
		var rule model.Rule
		err := cur.Decode(&rule)
		if err != nil {
			log.Panicln(err)
		}
		rules = append(rules, &rule)
	}
	return rules
}

func (db *DB) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Panicln(err)
	}
}
