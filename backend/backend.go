package main 

import "fmt"

var answers = map[string]string {
	"question1" : "System.out.println()",
	"question2" : "Chocolate is brown",
	"question3" : "Chicken tastes good",
	"question4" : "lol is laugh out loud",
}

func main () {
	
	level1()

}

func level1()  {
	var userInput string 
	var choice string
	fmt.Println("Welcome to Level 1")
	fmt.Println("This level will test your knowledge of input/output statements")
	fmt.Println ("Please drag and drop the words from the word bank to the appropriate blank spaces in the code")

	fmt.Scanln(&userInput) 

	check := userInput == answers["question1"]

	if check != true {
		choice = endGame()
		if choice != "No"{
			return
		} else {
			level1()
		}

	} else {
		fmt.Println("Correct") //output what we want 
		level2()
	}

}

func level2() {
	var userInput string 
	var choice string
	fmt.Println("Welcome to Level 2")
	fmt.Println("This level will test your knowledge of primitive data types and strings")
	fmt.Println ("Please drag and drop the words from the word bank to the appropriate blank spaces in the code")

	fmt.Scanln(&userInput) 

	check := userInput == answers["question2"]

	if check != true {
		choice = endGame()
		if choice != "yes"{
			return
		} else {
			level2()
		}

	} else {
		fmt.Println("Correct") //output what we want 
		level3()
	}
}

func level3() {
	var userInput string 
	var choice string
	fmt.Println("Welcome to Level 3")
	fmt.Println("This level will test your knowledge of conditional statements")
	fmt.Println ("Please drag and drop the words from the word bank to the appropriate blank spaces in the code")

	fmt.Scanln(&userInput)

	

	check := userInput == answers["question3"]

	if check != true {
		choice = endGame()
		if choice != "yes"{
			return
		} else {
			level3()
		}

	} else {
		fmt.Println("Correct") //output what we want 
		level4()
	}
}

func level4() {
	var userInput string 
	var choice string
	fmt.Println("Welcome to Level 1")
	fmt.Println("This level will test your knowledge of the input/output statements")
	fmt.Println ("Please drag and drop the words from the word bank to the appropriate blank spaces in the code")

	fmt.Scanln(&userInput) 


	check := userInput == answers["question4"]

	if check != true {
		choice = endGame()
		if choice != "yes"{
			return
		} else {
			level4()
		}

	} else {
		fmt.Println("Correct") //output what we want 
		return
	}
}
func endGame() string {
	var choice string
	fmt.Println("Incorrect Answer. Would you like to try again? Enter yes or no")
	fmt.Println("You got this!")
	fmt.Scanln(&choice)

	return choice
}
