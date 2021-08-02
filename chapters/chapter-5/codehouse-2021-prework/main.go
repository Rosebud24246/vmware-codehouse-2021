package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var response Response
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
	var questionBank = [4]string{"Please enter", "a ", " tired ", " b***h "}
	return questionBank[value]
}

func GetNextAnswer() string {
	var answerBank = [4]string{"Just", "A", "Little", "Depression"}
	return answerBank[value]
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

func GetResponse() Response {
	var response Response
	response.Code = "0"
	response.Results = GetQuestions()

	return response
}

func GetQuestions() [5]Question {
	//Question 1 - Introduction
	var question1 Question
	question1.Category = "java"
	question1.Difficulty = "easy"
	question1.Type = "multiple"
	question1.Question = "Which statement prints out: \"We need more women in STEM\" in Java?"
	question1.CorrectAnswer = "System.out.println(\"We need more women in STEM\");"

	//Incorrect Answers
	var incorrect = [3]string{"fmt.Println(We need more women in STEM)", "print(We need more women in STEM)", "out.print(We need more women in STEM)"}
	question1.IncorrectAnswers = incorrect[:]

	//Question 2 - Variables and Formatting
	var question2 Question
	question2.Category = "java"
	question2.Difficulty = "easy"
	question2.Type = "multiple"
	question2.Question = "What data type would be best suitable to represent the 44.6%/ of underepresented people at Google?"
	question2.CorrectAnswer = "Double"

	var incorrect2 = [3]string{"String", "int", "boolean"}
	question2.IncorrectAnswers = incorrect2[:]

	//Questions 3 - Conditionals
	var question3 Question
	question3.Category = "java"
	question3.Difficulty = "mediuem"
	question3.Type = "multiple"
	question3.Question = "What conditional is used when you only want a certain action to occur if a condition is met?"
	question3.CorrectAnswer = "If"

	var incorrect3 = [3]string{"String", "int", "boolean"}
	question3.IncorrectAnswers = incorrect3[:]

	//Questions 4 -
	var question4 Question
	question4.Category = "java"
	question4.Difficulty = "mediuem"
	question4.Type = "multiple"
	question4.Question = "What conditional is used when you only want a certain action to occur if a condition is met?"
	question4.CorrectAnswer = "If"

	var incorrect4 = [3]string{"String", "int", "boolean"}
	question4.IncorrectAnswers = incorrect4[:]

	//Questions 5 -
	var question5 Question
	question5.Category = "java"
	question5.Difficulty = "mediuem"
	question5.Type = "multiple"
	question5.Question = "What conditional is used when you only want a certain action to occur if a condition is met?"
	question5.CorrectAnswer = "If"

	var incorrect5 = [3]string{"String", "int", "boolean"}
	question5.IncorrectAnswers = incorrect5[:]

	var questions = [5]Question{question1, question2, question3, question4, question5}

	return questions
}

func main() {
	response = GetResponse()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
	r.GET("/api/response", GetResponse)
	r.POST("/api/quiz_questions", PostQuestion)
	r.DELETE("/api/quiz_questions/:response_code", DeleteQuestion)
	r.Run(":8090")
}
