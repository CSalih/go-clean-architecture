package problem

type UserStatusRequiredProblem struct {
	Problem
}

func (p UserStatusRequiredProblem) NewUserStatusRequiredProblem() *UserStatusRequiredProblem {
	return &UserStatusRequiredProblem{
		Problem{
			Type:   "https://example.com/problems/user-status-required",
			Title:  "User status is required",
			Status: 400,
		},
	}
}

func (p UserStatusRequiredProblem) Error() string {
	if p.Title == "" {
		return "User status is required"
	}
	return p.Title
}
