package presenter

// The Presenter is responsible for presentation logic.
// It is used to present the result of a use case.
type Presenter interface {
	// OnSuccess transforms the result to a client specific format.
	OnSuccess(result interface{}) error
	// OnError transforms an error to a client specific format.
	OnError(err error) error
}
