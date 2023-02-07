package presenter

import (
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"net/http"
)

type jsonResponsePresenter struct {
	Writer            http.ResponseWriter
	SuccessStatusCode int
}

func NewJsonResponsePresenter(writer http.ResponseWriter, successStatusCode int) Presenter {
	return &jsonResponsePresenter{
		Writer:            writer,
		SuccessStatusCode: successStatusCode,
	}
}

func (p *jsonResponsePresenter) OnSuccess(data interface{}) error {
	ctx := &router.Context{
		Writer: p.Writer,
	}
	return ctx.Json(p.SuccessStatusCode, data)
}

func (p *jsonResponsePresenter) OnError(err error) error {
	ctx := &router.Context{
		Writer: p.Writer,
	}
	return ctx.ProblemJson(err)
}
