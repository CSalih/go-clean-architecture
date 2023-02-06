package problem

type UserStatusRequiredProblem struct {
	Problem
}

func NewUserStatusRequiredProblem() Details {
	return Problem{
		Type:   "https://example.com/problems/user-status-required",
		Title:  "User status is required",
		Status: 400,
	}
}
