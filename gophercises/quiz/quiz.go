package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

type Quiz struct {
	Problems  []Problem
	Score     int
	TimeLimit time.Duration
}

func NewQuiz(path string, timeLimit time.Duration, shuffle bool) (Quiz, error) {
	problems, err := FromCSV(path)
	if err != nil {
		return Quiz{}, nil
	}
	if shuffle {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}
	return Quiz{
		Problems:  problems,
		Score:     0,
		TimeLimit: timeLimit,
	}, nil
}

func FromCSV(path string) ([]Problem, error) {
	f, err := os.Open(path)
	if err != nil {
		return []Problem{}, fmt.Errorf("failed to open the CSV file: %s\n", path)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return []Problem{}, fmt.Errorf("failed to parse the CSV file: %s\n", path)
	}
	problems := make([]Problem, len(rows))
	for i, row := range rows {
		question, answer := row[0], row[1]
		problems[i] = Problem{
			Question: question,
			Answer:   strings.TrimSpace(answer),
		}
	}
	return problems, nil
}

func (q *Quiz) Play() {
	timer := time.NewTimer(q.TimeLimit)
	answerCh := make(chan string)
	for i, problem := range q.Problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.Question)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			return
		case answer := <-answerCh:
			if answer == problem.Answer {
				q.Score++
			}
		}
	}
}

func Main() int {
	shuffle := flag.Bool("shuffle", false, "shuffle questions?")
	path := flag.String("csv", "problems.csv", "a CSV file in the format of 'question:answer'")
	timeLimit := flag.Duration("limit", 30*time.Second, "time limit in seconds")
	flag.Parse()
	quiz, err := NewQuiz(*path, *timeLimit, *shuffle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	quiz.Play()
	fmt.Printf("\nYou answered %d out of %d questions correctly!", quiz.Score, len(quiz.Problems))
	return 0
}
