package problems

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

func LoadProblems(path string) ([]Problem, error) {
	fileBytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(fileBytes)))

	problems := make([]Problem, 10)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		problems = append(problems, Problem{question: record[0], answer: record[1]})
	}

	return problems, nil
}
