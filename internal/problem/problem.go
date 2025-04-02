package problem

type Problem struct {
	question string
	answer   string
}

func New(question string, answer string) Problem {
	return Problem{question: question, answer: answer}
}
