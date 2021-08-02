package main

// type Todo struct {
// 	Id      int    `json:id`
// 	Question   string `json:"question"`
// 	Answer string `json:"answer"`
//
type Question struct {
	Id                int    `json:id`
	QuestionStatement string `json:"question_statement"`
	Answer            string `json:"answer"`
}

// type Answer struct {
// 	Id           int    `json:id`
// 	AnswerValue  string `json:"ans_value"`
// 	AnswerValueB string `json:"ans_value_b"`
// }
// type Right struct {
// 	Id               int    `json:id`
// 	RightAnswerValue string `json:"rt_ans_value"`
// }
