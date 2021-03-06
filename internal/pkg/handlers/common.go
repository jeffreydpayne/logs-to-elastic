package handlers

import (
	"errors"
	"fmt"
)

// Lambda function return
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func NewResponse(oks int, total int) (Response, error) {
	var err error
	var ok bool

	f := total - oks

	if f <= 0 {
		err = error(nil)
		ok = true
	} else {
		err = errors.New(fmt.Sprintf("failed to send records: %d", f))
		ok = false
	}

	return Response{
		Message: fmt.Sprintf("sent records: %d", oks),
		Ok:      ok,
	}, err
}
