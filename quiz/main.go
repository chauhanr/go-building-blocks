package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var limit int
var timeout int
var csvPath string

func init() {
	flag.IntVar(&limit, "limit", 10, "Limit flag sets the number of questions that the user wants to answer.")
	flag.IntVar(&timeout, "to", 30, "-to defines the timeout per question that the user needs to answer the quiz")
	flag.StringVar(&csvPath, "path", "problem.csv", "-path us used to determine the path of the csv")
}

func main() {
	flag.Parse()
	fmt.Printf("Limit value is %d\n", limit)
	fmt.Printf("Timeout vlaue is %d\n", timeout)

	recs, err := ReadProblemCsv(csvPath)
	if err != nil {
		return
	}
	fmt.Printf("records %v\n", recs)

}

func ReadProblemCsv(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading the problem.csv file %s\n", err)
		return [][]string{}, err
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
