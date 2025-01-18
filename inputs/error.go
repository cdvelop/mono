package inputs

import "strings"

// chainedText implementa la interfaz error
type chainedText struct {
	Message string
	Code    int
}

// Error implementa el método Error() de la interfaz error
func (c chainedText) Error() string {
	return c.Message
}

type myError struct {
	builder strings.Builder
}

var Error myError
