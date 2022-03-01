package problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func getAnsweredQuestions(questions []Question) int{
	valid_answers := 0
	var response string
	for _, question := range questions {
		fmt.Println("What ", question.question, ", sir?")
		fmt.Println("Please insert your answer: ")
		fmt.Scanln(&response)
		if question.answer == response {
			valid_answers++
		}
	}
  return valid_answers
}

func readProblems(file string) []Question {
	// open file
	questions := []Question{}
	f, err := os.Open(file)
	if err != nil {
		fmt.Print("Error opening the file")
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		array := strings.Split(scanner.Text(), ",")
		questions = append(questions, Question{question: array[0], answer: array[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return questions
}
