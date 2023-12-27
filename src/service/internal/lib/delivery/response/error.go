package response

import (
	"errors"

	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

type Fail struct {
	Error Error `json:"error"`
	Data  any   `json:"data"`
}

type EmptyFail struct {
	Error Error `json:"error"`
}

type options struct {
	status status.Code
	err    error
	data   any
}

type Option func(opts *options)

func WithStatus(status status.Code) Option {
	return func(opts *options) {
		opts.status = status
	}
}

func WithError(err error) Option {
	return func(opts *options) {
		opts.err = err
	}
}

func WithData(data any) Option {
	return func(opts *options) {
		opts.data = data
	}
}

func NewError(r *http.Request, opts ...Option) *http.Response {
	options := options{
		status: status.InternalServerError,
		err:    errors.New("internal server error"),
	}

	for _, opt := range opts {
		opt(&options)
	}

	resp := r.Respond().Code(options.status)

	if options.data == nil {
		resp, _ = resp.JSON(EmptyFail{
			Error: Error{
				Message:    options.err.Error(),
				StatusCode: int(options.status),
			},
		})
		return resp
	}

	resp, _ = resp.JSON(Fail{
		Error: Error{
			Message:    options.err.Error(),
			StatusCode: int(options.status),
		},
		Data: options.data,
	})
	return resp
}
