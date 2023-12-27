package request

import (
	"errors"

	"github.com/indigo-web/indigo/http"
)

type Key interface {
	CustomerKey
}

type CustomerKey struct {
}

var (
	ErrBrokenType = errors.New("broken type")
)

func GetKey[K Key, T any](req *http.Request) (T, error) {
	data, ok := req.Ctx.Value(K{}).(T)
	if !ok {
		return data, ErrBrokenType
	}

	return data, nil
}
