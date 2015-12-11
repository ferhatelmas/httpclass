package httpclass

import "fmt"

// StatusClass denotes the class of a http status code
type StatusClass int

// HTTP status code class enum
const (
	Informational StatusClass = iota + 1
	Success
	Redirection
	ClientError
	ServerError
)

// Get returns the class of given http status code
func Get(code int) (StatusClass, error) {
	if code < 100 || code >= 600 {
		return Informational, fmt.Errorf("Wrong status code %d", code)
	}
	if code < 200 {
		return Informational, nil
	} else if code < 300 {
		return Success, nil
	} else if code < 400 {
		return Redirection, nil
	} else if code < 500 {
		return ClientError, nil
	}
	return ServerError, nil
}

// MustGet does what Get does but panics on error.
func MustGet(code int) StatusClass {
	status, err := Get(code)
	if err != nil {
		panic(err)
	}
	return status
}
