package problem

import "github.com/CSalih/go-clean-architecture/internal/common/problem"

type UserNotFoundProblem struct {
	problem.Problem
}

func NewUserNotFoundProblem() problem.Details {
	return problem.Problem{
		Type:   "https://example.com/problems/user-not-found",
		Title:  "User not found",
		Status: 404,
	}
}
