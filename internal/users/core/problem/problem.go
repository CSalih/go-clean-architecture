package problem

// Problem is a machine-readable detail of error in a HTTP response
// See: https://www.rfc-editor.org/rfc/rfc7807.
type Problem struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
	Title  string `json:"title"`
}
