package main

type Question struct {
	Id      			 int    `json:"id"`
	Question_Statement   string `json:"question_statement"`
	Answer				 string `json:"answer"`
}
