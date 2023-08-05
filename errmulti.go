package toolbelt

import (
	"fmt"
	"strings"
)

type ErrMultiple struct {
	Errors    []error
	Separator string
}

func ErrJoin(errs ...error) error {
	return NewErrMultiple(errs)
}

func NewErrMultiple(errs []error) *ErrMultiple {
	return &ErrMultiple{Errors: errs, Separator: " | "}
}

func NewErrMultipleCustom(errs []error, sep string) *ErrMultiple {
	return &ErrMultiple{Errors: errs, Separator: sep}
}

func (e *ErrMultiple) Add(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *ErrMultiple) Addf(msg string, args ...any) {
	e.Errors = append(e.Errors, fmt.Errorf(msg, args...))
}

func (e *ErrMultiple) First() error {
	if len(e.Errors) == 0 {
		return nil
	}
	return e.Errors[0]
}

func (e *ErrMultiple) Error() string {
	var errors []string
	for _, err := range e.Errors {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, e.Separator)
}
