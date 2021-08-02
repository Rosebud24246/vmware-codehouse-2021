package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var quiz_questions []Question

func GetNextId() int {
	value := nextId
	nextId++
	return value
}
func GetNextWord() string {

	var a = [3]string{"System", "println", "out"}
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// rand number for 0 2

	//rand -> num
	// word -> b[num]
	return a[rand.Intn(len(a))]
}

// func GetQuestions(numberOfQuestions int) []Question {
// 	var q []Question
// 	var questionCount = len(quiz_questions)
// 	for i := 0; i < len(quiz_questions); i++ {

// 		q = append(q, quiz_questions[rand.Intn(questionCount)])
// 	}
// 	return q
// }
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": quiz_questions})
}

func PostTodo(c *gin.Context) {
	var item Question
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.Id = GetNextId()
	quiz_questions = append(quiz_questions, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(item.Id))
}

func DeleteTodo(c *gin.Context) {
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

func addQuestion(question string, answer string) {
	quiz_questions = append(quiz_questions,
		Question{
			Id:                GetNextId(),
			QuestionStatement: question,
			Answer:            answer,
		})
}

func preloadQuestions() {
	addQuestion("System___.println()", "out")
}

func main() {

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
	r.GET("/api/todos", GetTodos)
	r.POST("/api/todos", PostTodo)
	r.DELETE("/api/todos/:id", DeleteTodo)
	r.Run(":8090")
}
