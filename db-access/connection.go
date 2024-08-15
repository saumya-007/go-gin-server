package dbaccess

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	DbDetailsEntity "github.com/saumya-007/go-gin-server/entity"
)

var MongoClient *mongo.Client

// Solved Question Collection
var SolvedQuestionCollection *mongo.Collection
var SolvedQuestionsDbDetails DbDetailsEntity.DbDetails

// Other Collection ...

func init() {
	// note: init all this will run before main
	SolvedQuestionsDbDetails.DbName = "DSA"
	SolvedQuestionsDbDetails.CollectionName = "SolvedQuestions"
}

func ConnectMongo() {
	connectionString := "mongodb://localhost:27017/"

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	fmt.Println("===> Connected to mongodb <===")
	InitAllCollections()
	fmt.Println("===> All Collections Initiated <===")
}

func InitAllCollections() {
	SolvedQuestionCollection = MongoClient.Database(SolvedQuestionsDbDetails.DbName).Collection(SolvedQuestionsDbDetails.CollectionName)
}
