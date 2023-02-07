package presenter

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/users/core/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"net/http"
)

type jsonHttpPresenter struct {
	Writer            http.ResponseWriter
	SuccessStatusCode int
}

func NewJsonHttpPresenter(writer http.ResponseWriter, successStatusCode int) presenter.Presenter {
	return &jsonHttpPresenter{
		Writer:            writer,
		SuccessStatusCode: successStatusCode,
	}
}

func (p *jsonHttpPresenter) OnSuccess(data interface{}) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return p.OnError(err)

	}

	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(p.SuccessStatusCode)
	_, err = p.Writer.Write(jsonString)
	if err != nil {
		return p.OnError(err)
	}
	return nil
}

func (p *jsonHttpPresenter) OnError(err error) error {
	data := ProblemFromError(err)
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	p.Writer.Header().Set("Content-Type", "application/problem+json")
	p.Writer.WriteHeader(data.Status)
	_, err = p.Writer.Write(jsonString)
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
