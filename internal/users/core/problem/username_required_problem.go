package problem

type UsernameRequiredProblem struct {
	Problem
}

func (p UsernameRequiredProblem) NewUsernameRequiredProblem() *UsernameRequiredProblem {
	return &UsernameRequiredProblem{
		Problem{
			Type:   "https://example.com/problems/username-required",
			Title:  "Username is required",
			Status: 400,
		},
	}
}

func (p UsernameRequiredProblem) Error() string {
	if p.Title == "" {
		return "Username is required"
	}
	return p.Title
}
