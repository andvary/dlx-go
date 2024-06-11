package mydlx

import "fmt"

type InputError struct {
	msg string
}

func (e InputError) Error() string {
	return fmt.Sprintf("input error: %s", e.msg)
}

type CoverError struct {
	msg string
}

func (e CoverError) Error() string {
	return fmt.Sprintf("cover error: %s", e.msg)
}
