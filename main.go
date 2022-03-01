package main

import (
	"fmt"

	"./problems"
)

func main() {
	questions := problems.ReadProblems("problems.csv")
  valid_answers := problems.GetAnsweredQuestions(questions)
	fmt.Println("Number of questions", len(questions))
	fmt.Println("Question correctly answered: ", valid_answers)
}
