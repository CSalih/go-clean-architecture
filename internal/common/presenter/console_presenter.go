package presenter

import (
	"fmt"
	"os"
)

type consolePresenter struct {
}

func NewConsolePresenter() Presenter {
	return &consolePresenter{}
}

func (p consolePresenter) OnSuccess(result interface{}) error {
	_, err := fmt.Printf("%+v\n", result)
	return err
}

func (p consolePresenter) OnError(err error) error {
	_, printErr := fmt.Fprintln(os.Stderr, err)
	return printErr
}
