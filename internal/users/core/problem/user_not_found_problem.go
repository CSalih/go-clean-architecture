package problem

type UserNotFoundProblem struct {
	Problem
}

func NewUserNotFoundProblem() Details {
	return Problem{
		Type:   "https://example.com/problems/user-not-found",
		Title:  "User not found",
		Status: 404,
	}
}
