package loader

import (
	_ "embed"
	"encoding/csv"
	"io"
	"strings"

	"github.com/TB516/gophercises-quiz-game/internal/problem"
)

//go:embed assets/problems.csv
var problemsFile string

func LoadProblems() ([]problem.Problem, error) {
	reader := csv.NewReader(strings.NewReader(problemsFile))

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
