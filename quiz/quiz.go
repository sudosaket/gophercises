package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	csvQuizFile = flag.String("csv", "problems.csv", "csv file with format 'question,answer'")
	timeLimit   = flag.Duration("tl", 30, "time limit for the quiz in seconds")
)

type problem struct {
	q string
	a string
}

func main() {
	flag.Parse()

	f, err := os.Open(*csvQuizFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	problems := parseProblems(lines)
	done := make(chan bool)
	q := &quiz{
		correct:  0,
		problems: problems,
		done:     done,
	}
	go q.start()
	select {
	case <-done:
		fmt.Printf("You scored %d out of %d!\n", q.correct, len(problems))
	case <-time.After(*timeLimit):
		fmt.Printf("You scored %d out of %d!\n", q.correct, len(problems))
	}
}

type quiz struct {
	correct  int
	problems []problem
	done     chan bool
}

func (q *quiz) start() {
	for i, p := range q.problems {
		fmt.Printf("Problem #%d: %s \n", i+1, p.q)
		var answer string
		_, _ = fmt.Scanf("%s", &answer)
		if p.a == answer {
			q.correct++
		}
	}
	q.done <- true
}

func parseProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return problems
}
