package problem

import "github.com/CSalih/go-clean-architecture/internal/common/problem"

type UserStatusRequiredProblem struct {
	problem.Problem
}

func NewUserStatusRequiredProblem() problem.Details {
	return problem.Problem{
		Type:   "https://example.com/problems/user-status-required",
		Title:  "User status is required",
		Status: 400,
	}
}
