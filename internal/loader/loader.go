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

func LoadProblems(ch chan []problem.Problem) {
	defer close(ch)

	reader := csv.NewReader(strings.NewReader(problemsFile))

	problems := make([]problem.Problem, 15)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			ch <- problems
			return
		} else if err != nil {
			ch <- nil
			return
		}

		problems = append(problems, problem.New(record[0], record[1]))
	}
}
