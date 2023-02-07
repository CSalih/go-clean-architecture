package problem

import "github.com/CSalih/go-clean-architecture/internal/common/problem"

type UsernameRequiredProblem struct {
	problem.Problem
}

func NewUsernameRequiredProblem() problem.Details {
	return problem.Problem{
		Type:   "https://example.com/problems/username-required",
		Title:  "Username is required",
		Status: 400,
	}
}
