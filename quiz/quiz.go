package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var csvQuizFile = flag.String("csv", "problems.csv", "csv file with format 'question,answer'")

type problem struct {
	q string
	a string
}

func main() {
	flag.Parse()

	f, err := os.Open(*csvQuizFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseProblems(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s \n", i+1, p.q)
		var answer string
		_, _ = fmt.Scanf("%s", &answer)
		if p.a == answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d!\n", correct, len(problems))
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
