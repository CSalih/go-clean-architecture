package problem

type UsernameRequiredProblem struct {
	Problem
}

func NewUsernameRequiredProblem() Details {
	return Problem{
		Type:   "https://example.com/problems/username-required",
		Title:  "Username is required",
		Status: 400,
	}
}
