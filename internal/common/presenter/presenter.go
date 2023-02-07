package presenter

type Presenter interface {
	OnSuccess(result interface{}) error
	OnError(err error) error
}
