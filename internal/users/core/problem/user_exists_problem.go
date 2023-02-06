package problem

type UserExistsProblem struct {
	Problem
}

func NewUserExistsProblem() Details {
	return &UserExistsProblem{
		Problem: Problem{
			Type:   "https://example.com/problems/user-exists",
			Title:  "Username already exist",
			Status: 400,
		},
	}
}
