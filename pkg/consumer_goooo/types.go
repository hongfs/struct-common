package consumer_goooo

import "errors"

var ErrSuccess = errors.New("success")

var ErrTimeout = errors.New("timeout")

var ErrRetry = errors.New("retry")
