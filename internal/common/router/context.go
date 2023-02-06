package router

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

func (c *Context) Json(code int, data interface{}) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	_, err = c.Writer.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) ProblemJson(err error) error {
	data := ProblemFromError(err)
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.Writer.Header().Set("Content-Type", "application/problem+json")
	c.Writer.WriteHeader(data.Status)
	_, err = c.Writer.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}

func ProblemFromError(err error) problem.Problem {
	switch prob := err.(type) {
	case problem.Details:
		return prob.GetProblem()
	default:
		return problem.Problem{
			Type:   "https://example.com/problems/internal-server-error",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
		}
	}
}
