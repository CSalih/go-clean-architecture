package presenter

type Presenter interface {
	OnSuccess(result interface{})
	OnError(err error)
}
