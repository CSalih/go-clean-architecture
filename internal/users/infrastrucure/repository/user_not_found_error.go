package repository

type UserNotFoundError struct {
	message string
}

func (err UserNotFoundError) Error() string {
	if err.message == "" {
		return "User not found"
	}
	return err.message
}
