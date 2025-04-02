package loader

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/TB516/gophercises-quiz-game/internal/problem"
)

func LoadProblems(path string) ([]problem.Problem, error) {
	fileBytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(fileBytes)))

	problems := make([]problem.Problem, 10)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		problems = append(problems, problem.New(record[0], record[1]))
	}

	return problems, nil
}
