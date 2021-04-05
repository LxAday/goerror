package goerror

import (
	"fmt"
	"runtime"
	"strings"
)

//GoError
type GoError interface {
	Error() string
	Where() string
}

//Error
type Error struct {
	error string
	where string
}

//Error  error message
func (e *Error) Error() string {
	return e.error
}

//Where error from
func (e *Error) Where() string {
	return e.where
}

//New create an error
func New(error string) *Error {
	return &Error{
		error: error,
		where: caller(1),
	}
}

//Wrap nested error
func Wrap(errs error, msg string) *Error {
	err := &Error{}
	switch v := errs.(type) {
	case nil:
		return nil
	case GoError:
		err.error = v.Error()
		err.where = v.Where()
	default:
		err.error = errs.Error()
		err.where = caller(1)
	}
	var es []string
	if err.error != "" {
		es = append(es, err.error)
	}
	if msg != "" {
		es = append(es, msg)
	}
	if len(es) > 0 {
		err.error = strings.Join(es, ":")
	}
	return err
}

//caller
func caller(calldepth int) string {
	_, file, line, _ := runtime.Caller(calldepth + 1)
	return fmt.Sprintf("%s:%d", file, line)
}
