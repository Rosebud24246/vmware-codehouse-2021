package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var quiz_questions []Question
var value int

func GetNextId() int {
	value = nextId
	nextId++
	return value
}

func GetQuestion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": quiz_questions})
}

func GetNextQuestion() string {
	var questionBank = [4] string {"Please enter", "a ", " tired ", " b***h "}
	return questionBank[value]
}

func GetNextAnswer() string {
	var answerBank = [4] string {"Just", "A", "Little", "Depression"}
	return answerBank [value]
}

func PostQuestion(c *gin.Context) {
	var item Question
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.Id = GetNextId()
	quiz_questions = append(quiz_questions, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(item.Id))
}

func DeleteQuestion(c *gin.Context) {
	idString := c.Param("id")

	if id, err := strconv.Atoi(idString); err == nil {
		for index := range quiz_questions {
			if quiz_questions[index].Id == id {
				quiz_questions = append(quiz_questions[:index], quiz_questions[index+1:]...)
				c.Writer.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
}

func main() {
	quiz_questions = append(quiz_questions, Question{
		Id: 				GetNextId(), 
		Question_Statement: GetNextQuestion(), 
		Answer:				GetNextAnswer()})

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
	r.GET("/api/todos", GetQuestion)
	r.POST("/api/todos", PostQuestion)
	r.DELETE("/api/todos/:id", DeleteQuestion)
	r.Run(":8090")
}
