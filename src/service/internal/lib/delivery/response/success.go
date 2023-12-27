package response

import (
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
)

type Success struct {
	Data any `json:"data"`
}

func NewSuccess(r *http.Request, data any) *http.Response {
	resp := r.Respond().Code(status.OK)
	if data != nil {
		resp, _ = resp.JSON(Success{Data: data})
	}

	return resp
}
