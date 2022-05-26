package errs

import "github.com/pkg/errors"

var (
	InvalidPropertyErr  = errors.New("invalid property")
	InvalidParameterErr = errors.New("invalid parameter")
)
