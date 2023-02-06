package problem

type UserExistsProblem struct {
	Problem
}

func (p UserExistsProblem) NewUserExistsProblem() *UserExistsProblem {
	return &UserExistsProblem{
		Problem{
			Type:   "https://example.com/problems/user-exists",
			Title:  "Username already exist",
			Status: 400,
		},
	}
}

func (p UserExistsProblem) Error() string {
	if p.Title == "" {
		return "Username already exist"
	}
	return p.Title
}
