package dbaccess

import (
	"context"
	"fmt"
	"log"

	SolvedQuestionsEntity "github.com/saumya-007/go-gin-server/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSolvedQuestions() []primitive.M {

	// note: keep it empty if we don't want to apply any filter
	filter := bson.M{
		"is_deleted": false,
	}
	cnxt := context.Background()

	cursor, err := SolvedQuestionCollection.Find(cnxt, filter)
	if err != nil {
		log.Fatal()
	}
	defer cursor.Close(cnxt)

	var solvedQuestionsFromDb []primitive.M

	for cursor.Next(cnxt) {
		var solvedQuestionDetails bson.M
		if err := cursor.Decode(&solvedQuestionDetails); err != nil {
			log.Fatal(err)
		}
		solvedQuestionsFromDb = append(solvedQuestionsFromDb, solvedQuestionDetails)
	}

	return solvedQuestionsFromDb
}

func GetSolvedQuestionByLink(questionLink string) primitive.M {
	filter := bson.M{
		"question_link": questionLink,
	}

	var solvedQuestionFromDb bson.M

	cursor := SolvedQuestionCollection.FindOne(context.Background(), filter)
	if err := cursor.Decode(&solvedQuestionFromDb); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}

	// Convert ObjectId to string
	if id, ok := solvedQuestionFromDb["_id"].(primitive.ObjectID); ok {
		solvedQuestionFromDb["_id"] = id.Hex()
	}

	return solvedQuestionFromDb
}

func GetSolvedQuestionById(solvedQuestionId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(solvedQuestionId)
	filter := bson.M{"_id": id}

	var solvedQuestionFromDb bson.M

	cursor := SolvedQuestionCollection.FindOne(context.Background(), filter)

	fmt.Println(cursor, filter)

	if err := cursor.Decode(&solvedQuestionFromDb); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}

	// Convert ObjectId to string
	if id, ok := solvedQuestionFromDb["_id"].(primitive.ObjectID); ok {
		solvedQuestionFromDb["_id"] = id.Hex()
	}

	return solvedQuestionFromDb
}

func AddSolvedQuestion(questionDetails SolvedQuestionsEntity.QuestionsDetails) primitive.ObjectID {
	inserted, err := SolvedQuestionCollection.InsertOne(context.Background(), questionDetails)
	if err != nil {
		log.Fatal(err)
	}

	return inserted.InsertedID.(primitive.ObjectID)
}

func UpdateSolvedQuestion(solvedQuestionId string, questionDetails SolvedQuestionsEntity.QuestionsDetails) int64 {
	// note: this is because mongo has _id internally as everything inside of mongo is not json it's technically bson
	id, _ := primitive.ObjectIDFromHex(solvedQuestionId)
	filter := bson.M{"_id": id}

	// note: converting the struct to BSON
	dataToUpdate, err := bson.Marshal(questionDetails)
	if err != nil {
		log.Fatal(err)
	}

	// note: convert BSON data to a map for the update operation
	var updateData bson.M
	if err := bson.Unmarshal(dataToUpdate, &updateData); err != nil {
		log.Fatal(err)
	}

	updates := bson.M{"$set": updateData}

	updateMetaData, err := SolvedQuestionCollection.UpdateOne(context.Background(), filter, updates)
	if err != nil {
		log.Fatal(err)
	}

	return updateMetaData.ModifiedCount
}

func HardDeleteSolvedQuestion(solvedQuestionId string) int64 {
	id, _ := primitive.ObjectIDFromHex(solvedQuestionId)
	filter := bson.M{"_id": id}

	deleteMetaData, err := SolvedQuestionCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return deleteMetaData.DeletedCount
}

func HardDeleteAllSolvedQuestion(solvedQuestionId string) int64 {
	filter := bson.M{}

	deleteMetaData, err := SolvedQuestionCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return deleteMetaData.DeletedCount
}

func SoftDeleteSolvedQuestion(solvedQuestionId string) int64 {
	id, _ := primitive.ObjectIDFromHex(solvedQuestionId)
	filter := bson.M{"_id": id}

	dataToUpdate := bson.M{"$set": bson.M{"is_deleted": true}}

	updateMetaData, err := SolvedQuestionCollection.UpdateOne(context.Background(), filter, dataToUpdate)

	if err != nil {
		log.Fatal(err)
	}

	return updateMetaData.ModifiedCount
}
