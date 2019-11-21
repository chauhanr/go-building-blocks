package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var limit int
var timeout int
var csvPath string

type QElement struct {
	Id       int
	Question string
	Answer   string
	Timeout  int
}

func init() {
	flag.IntVar(&limit, "limit", 10, "Limit flag sets the number of questions that the user wants to answer.")
	flag.IntVar(&timeout, "to", 30, "-to defines the timeout per question that the user needs to answer the quiz")
	flag.StringVar(&csvPath, "path", "problem.csv", "-path us used to determine the path of the csv")
}

func main() {
	flag.Parse()
	elements, err := ReadProblemCsv(csvPath)
	if err != nil {
		return
	}
	//fmt.Printf("records %v\n", elements)
	s := AskQuestions(elements)
	fmt.Printf("Total Score for Quiz: %d\n", s)
}

func ReadProblemCsv(path string) ([]QElement, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading the problem.csv file %s\n", err)
		return []QElement{}, err
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return []QElement{}, err
	}
	var qElements []QElement

	for index, rec := range records {
		q := rec[0]
		a := rec[1]
		qElement := QElement{Id: index + 1, Question: q, Answer: a, Timeout: timeout}
		qElements = append(qElements, qElement)
	}
	return qElements, nil
}

/*AskQuestions will ask the questions using go routine and take care of the timeuse as well*/
func AskQuestions(qSet []QElement) int {
	score := 0
	done := make(chan interface{})
	defer close(done)
	for s := range ask(done, qRepeat(done, qSet)) {
		score = score + s
	}
	return score
}

func qRepeat(done <-chan interface{}, ques []QElement) <-chan QElement {
	qStream := make(chan QElement)
	go func() {
		defer close(qStream)
		for {
			for i := 0; i < limit; i++ {
				select {
				case <-done:
					return
				case qStream <- ques[i]:
				}
			}
		}
	}()
	return qStream
}

func ask(done <-chan interface{}, qs <-chan QElement) <-chan int {
	qscore := make(chan int)
	reader := bufio.NewReader(os.Stdin)

	go func() {
		defer close(qscore)
		for i := 0; i < limit; i++ {
			select {
			case <-done:
				return
			case q := <-qs:
				// now the question needs to be asked.
				fmt.Printf("\n%d. %s: ", q.Id, q.Question)
				text, _ := reader.ReadString('\n')
				if runtime.GOOS == "windows" {
					text = strings.TrimRight(text, "\r\n")
				} else {
					text = strings.TrimRight(text, "\n")
				}
				if strings.Compare(text, q.Answer) == 0 {
					qscore <- 1
				} else {
					//fmt.Printf("Answer %s is wrong the answer must be %s\n", text, q.Answer)
					qscore <- 0
				}
			}
		}
	}()
	return qscore
}
