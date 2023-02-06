package problem

// Problem represents the required details of error in a HTTP response.
type Problem struct {
	// Type is a URI reference [RFC3986] that identifies the problem type.
	Type string `json:"type"`
	// Title is a short, human-readable summary of the problem type.
	Title string `json:"title"`
	// Status is the HTTP status code [RFC7231].
	Status int `json:"status"`
}

// Details is a machine-readable detail of error in a HTTP response.
//
// See: https://www.rfc-editor.org/rfc/rfc7807.
type Details interface {
	Error() string
	GetProblem() Problem
}

// Error returns the Title of error otherwise "Unknown problem".
func (p Problem) Error() string {
	if p.Title == "" {
		return "Unknown problem"
	}
	return p.Title
}

// GetProblem returns the details of error.
func (p Problem) GetProblem() Problem {
	return p
}
