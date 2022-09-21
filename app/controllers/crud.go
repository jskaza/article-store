package controllers

import (
	"context"
	"log"
	"os"

	"github.com/jskaza/open-journal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = os.Getenv("MONGODB")

func GetPapers(category string) ([]models.Paper, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	papersCollection := client.Database("OpenJournal").Collection("papers")
	cursor, err := papersCollection.Find(context.TODO(), models.Paper{Category: category})
	if err != nil {
		log.Fatal(err)
	}
	var papers []models.Paper
	if err = cursor.All(context.TODO(), &papers); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(papers); i++ {
		papers[i].HexID = papers[i].ID.Hex()
	}
	return papers, nil
}

func GetPaper(HexID string) (models.DisplayPaper, error) {
	ID, _ := primitive.ObjectIDFromHex(HexID)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	papersCollection := client.Database("OpenJournal").Collection("papers")

	var paper models.DisplayPaper
	papersCollection.FindOne(context.TODO(), models.Paper{ID: ID}).Decode(&paper)

	return paper, nil
}

func InsertPaper(paper models.Paper) (string, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	papersCollection := client.Database("OpenJournal").Collection("papers")

	res, err := papersCollection.InsertOne(context.TODO(), paper)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetName(id primitive.ObjectID) (string, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	usersCollection := client.Database("OpenJournal").Collection("users")

	var name string
	usersCollection.FindOne(context.TODO(), models.User{ID: id}).Decode(&name)

	return name, nil
}
