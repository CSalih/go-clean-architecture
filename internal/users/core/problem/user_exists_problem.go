package problem

type UserExistsProblem struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
	Title  string `json:"title"`
}

func (p UserExistsProblem) NewUserExistsProblem() *UserExistsProblem {
	return &UserExistsProblem{
		Type:   "https://example.com/problems/user-not-found",
		Title:  "Username already exist",
		Status: 400,
	}
}

func (p UserExistsProblem) Error() string {
	if p.Title == "" {
		return "Username already exist"
	}
	return p.Title
}
