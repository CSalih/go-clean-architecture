package problem

type UserNotFoundProblem struct {
	Problem
}

func (p UserNotFoundProblem) NewUserNotFoundProblem() *UserNotFoundProblem {
	return &UserNotFoundProblem{
		Problem{
			Type:   "https://example.com/problems/user-not-found",
			Title:  "User not found",
			Status: 404,
		},
	}
}

func (p UserNotFoundProblem) Error() string {
	if p.Title == "" {
		return "User not found"
	}
	return p.Title
}
