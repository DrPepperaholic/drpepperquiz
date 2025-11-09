package main

import "fmt"

func main() {
	fmt.Println("Hello! What is your name?")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hey %v!, Welcome to the Dr Pepper Quiz! Note: Answers are Case Sensitive.\n", name)

	var likesdrpepper string
	fmt.Println("Do you like Dr Pepper? Yes or No.")
	fmt.Scan(&likesdrpepper)
	if likesdrpepper == "Yes" {
		fmt.Println("Nice! You are on the right quiz then!")
	} else {
		fmt.Println("Why are you taking the quiz then? Maybe give the Cola Quiz a try instead! :)")
	}

	fmt.Printf("What is better, Dr Pepper or Mountain Dew? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "Dr Pepper" {
		fmt.Println("Correct! We LOVE Dr Pepper!")
		fmt.Println("Thanks for taking the quiz, byeee!")
	} else {
		fmt.Println("You are kicked from the Dr Pepper Fan Club.")
		fmt.Println("Bye.")
	}
}
