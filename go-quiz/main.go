package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filePath := flag.String("csv", "problems.csv", "a csv file with 'question,answer'")
	timerDuration := flag.Int("timer", 30, "timer duration in secs")
	flag.Parse()
	csvfile, err := os.Open(*filePath)

	if err != nil {
		log.Fatalln("could not open csv file ", err)
	}
	questions := csv.NewReader(csvfile)
	allQuestions, err := questions.ReadAll()
	if err != nil {
		log.Fatalln("could not read content in the csv file")
	}
	problemList := make([]problem, len(allQuestions))
	for i, line := range allQuestions {
		problemList[i] = problem{q: line[0], a: line[1]}
	}
	correct := 0
	timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)

	for _, p := range problemList {
		fmt.Printf("%v ?", p.q)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			fmt.Println(fmt.Sprintf("Answered %d out of %d questions", correct, len(problemList)))
			return
		case ans := <-answerChan:
			if ans == p.a {
				correct++
			}
		}
	}

	fmt.Printf("Answered %d questions correctly out of %d", correct, len(allQuestions))
}

type problem struct {
	q string
	a string
}
