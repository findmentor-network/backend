package mongohelper

import (
	"github.com/pkg/errors"
)

var (
	ConnectionError = errors.New("Mongo connection cannot be established.")
	PingError       = errors.New("Mongo connection cannot be pinged")
)
