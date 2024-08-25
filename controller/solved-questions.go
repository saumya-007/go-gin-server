package apis

import (
	// note: used for server side logging
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	SolvedQuestionsDdAccess "github.com/saumya-007/go-gin-server/db-access"
	SolvedQuestionsEntity "github.com/saumya-007/go-gin-server/entity"
)

// note: gin context will have all of the information about a request
func GetSolvedQuestions(c *gin.Context) {

	solvedQuestionsList := SolvedQuestionsDdAccess.GetSolvedQuestions()

	c.IndentedJSON(http.StatusOK, solvedQuestionsList)
}

func GetSolvedQuestionsById(c *gin.Context) {
	solvedQuestionId := c.Param("id")

	solvedQuestion := SolvedQuestionsDdAccess.GetSolvedQuestionById(solvedQuestionId)

	c.IndentedJSON(http.StatusOK, solvedQuestion)
}

func AddSolvedQuestion(c *gin.Context) {
	var questionDetails SolvedQuestionsEntity.QuestionsDetails

	// note: Bind Json takes in a pointer, it directly does not return the error message and status instead the BindJson method does that for us.
	if err := c.BindJSON(&questionDetails); err != nil {
		return
	}

	questionDetailFromDb := SolvedQuestionsDdAccess.GetSolvedQuestionByLink(questionDetails.QuestionLink)

	if questionDetailFromDb != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Question with the same link already added",
		})
		return
	}

	questionDetails.IsDeleted = false

	insertedId := SolvedQuestionsDdAccess.AddSolvedQuestion(questionDetails)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"id": insertedId,
	})
}

func UpdateSolvedQuestion(c *gin.Context) {
	solvedQuestionId := c.Param("id")

	var questionDetails SolvedQuestionsEntity.QuestionsDetails
	if err := c.BindJSON(&questionDetails); err != nil {
		return
	}

	questionDetailFromDb := SolvedQuestionsDdAccess.GetSolvedQuestionByLink(questionDetails.QuestionLink)

	if questionDetailFromDb != nil && questionDetailFromDb["_id"] != solvedQuestionId {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Question with the same link already added",
		})
		return
	}

	updatedCount := SolvedQuestionsDdAccess.UpdateSolvedQuestion(solvedQuestionId, questionDetails)

	c.IndentedJSON(http.StatusOK, gin.H{
		"updated_count": updatedCount,
	})

}

func DeleteSolvedQuestion(c *gin.Context) {
	solvedQuestionId := c.Param("id")

	questionDetailFromDb := SolvedQuestionsDdAccess.GetSolvedQuestionById(solvedQuestionId)

	if questionDetailFromDb == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Question details not found",
		})
		return
	}

	deletedCount := SolvedQuestionsDdAccess.SoftDeleteSolvedQuestion(solvedQuestionId)

	log.Printf("Deleted Count : %v/n", deletedCount)

	c.IndentedJSON(http.StatusOK, gin.H{
		"deleted_count": deletedCount,
	})
}
