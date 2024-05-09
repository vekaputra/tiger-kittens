package error

import (
	"errors"

	pkgerrors "github.com/pkg/errors"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
)

func ErrWithStackTrace(err error) error {
	if err == nil {
		return nil
	}

	type stackTracer interface {
		StackTrace() pkgerrors.StackTrace
	}

	_, ok := err.(stackTracer)
	if !ok {
		return pkgerrors.WithStack(err)
	}

	return err
}

func ErrCause(err error) error {
	if err == nil {
		return nil
	}

	var causeErr customerror.CustomError
	switch {
	case errors.As(pkgerrors.Cause(err), &causeErr):
		return causeErr
	default:
		return customerror.ErrorInternalServer
	}
}
