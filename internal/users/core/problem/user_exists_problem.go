package problem

import "github.com/CSalih/go-clean-architecture/internal/common/problem"

type UserExistsProblem struct {
	problem.Problem
}

func NewUserExistsProblem() problem.Details {
	return &UserExistsProblem{
		Problem: problem.Problem{
			Type:   "https://example.com/problems/user-exists",
			Title:  "Username already exist",
			Status: 400,
		},
	}
}
