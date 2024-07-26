package chainableerror

import (
	"errors"
	"fmt"
)

type ChainableError struct {
	baseErr  error
	childErr error
}

func New(msg string) *ChainableError {
	return &ChainableError{
		baseErr: errors.New(msg),
	}
}

func (e *ChainableError) Error() string {
	if e.childErr == nil {
		return e.baseErr.Error()
	}
	return fmt.Sprintf("%s: %s", e.baseErr.Error(), e.childErr.Error())
}

func (e *ChainableError) Unwrap() error {
	return e.childErr
}

func (e *ChainableError) Is(err error) bool {
	if errors.Is(err, e.baseErr) {
		return true
	}
	return errors.Is(err, e.childErr)
}

func (e *ChainableError) Wrap(err error) *ChainableError {
	return &ChainableError{
		baseErr:  e.baseErr,
		childErr: err,
	}
}
