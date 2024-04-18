package units

import (
	"errors"
)

var (
	ErrParse       = errors.New("Parse error")
	ErrUnknownUnit = errors.New("Unknown unit")
)
