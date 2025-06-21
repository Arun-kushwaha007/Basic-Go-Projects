package main
import "fmt"

func main() {
	 fmt.Println("Welcome to the Quiz Game!")
	 fmt.Println("In this game, you will answer a series of questions.")
	 fmt.Println("Let's get started!")
	 // Declare a variable to hold the score
	 fmt.Printf("Please enter your name:")
	 var name string
	 fmt.Scan(&name)
	 fmt.Printf("Hello, %s! Let's begin the quiz.\n", name)
	 fmt.Printf("Enter your age:")
	 var age uint
	 var num_question int = 3
	 fmt.Scan(&age)
	 if age < 10 {
		 fmt.Println("You must be at least 18 years old to play this game.") 
		 return
	 }
	 fmt.Println("Great! Let's start the quiz.")
	 // Declare a variable to hold the score
	 score := 0
	 // Question 1
	 fmt.Println("Question 1: What is the capital of France?")
	 fmt.Println("a) Berlin")
	 fmt.Println("b) Madrid")
	 fmt.Println("c) Paris")
	 var answer1 string
	 fmt.Scan(&answer1)
	 if answer1 == "c" || answer1 == "C" {
		 fmt.Println("Correct!")
		 score++
	 } else {
		 fmt.Println("Incorrect. The correct answer is c) Paris.")
	 }
	 // Question 2
	 fmt.Println("Question 2: What is the largest planet in our solar system?")
	 fmt.Println("a) Earth")
	 fmt.Println("b) Jupiter")
	 fmt.Println("c) Saturn")
	 var answer2 string
	 fmt.Scan(&answer2)
	 if answer2 == "b" || answer2 == "B" {

		 fmt.Println("Correct!")
		 score++
	 } else {
		 fmt.Println("Incorrect. The correct answer is b) Jupiter.")
	 }
	 // Question 3
	 fmt.Println("Question 3: What is better, the RTX 3080 or the RTX 3090?")
	 var answer3 string
	 var answer4 string

	 fmt.Scan(&answer3, &answer4)
	 
	 
	 if answer3 + " " + answer4 == "RTX 3090" || answer3 + " " + answer4 == "rtx 3090" {
		 fmt.Println("Correct!")
		 score++
	 } else {
		 fmt.Println("Incorrect. The correct answer is RTX 3090.")
	 }
	 // Display the final score
	 fmt.Printf("Quiz over! Your final score is: %d out of 3\n", score)
	 percent := (float64(score)/ float64(num_question))* 100
	 fmt.Printf("Your score percentage is: %v%%. \n", percent)
	 // Provide feedback based on the score
	 if score == 3 {
		 fmt.Println("Excellent job! You're a quiz master!")
	 }
	 if score == 2 {
		 fmt.Println("Good job! You have a solid understanding of the topics.")
	 }
	 if score == 1 {
		 fmt.Println("Not bad! But you can do better. Keep learning!")
	 }
	 if score == 0 {
		 fmt.Println("Better luck next time! Don't be discouraged, keep trying!")
	 }
	 // Thank the player for playing
	 fmt.Printf("Thank you for playing, %s! We hope you enjoyed the quiz.\n", name)
	 fmt.Println("Goodbye!")

}
