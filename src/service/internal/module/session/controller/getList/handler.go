package getListHandler

import (
	"context"
	"time"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/lib/delivery/request"
	"github.com/Shteyd/holidays/src/service/internal/lib/delivery/response"
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
	"github.com/indigo-web/indigo/router/inbuilt"
)

type sessionUsecase interface {
	GetSessionList(ctx context.Context, customerID int64) ([]entity.Session, error)
}

func New(usecase sessionUsecase, timeout time.Duration) inbuilt.Handler {
	return func(r *http.Request) *http.Response {
		customerID, err := request.GetKey[request.CustomerKey, int64](r)
		if err != nil {
			return response.NewError(r,
				response.WithStatus(status.InternalServerError),
				response.WithError(err),
			)
		}

		ctx, cancel := context.WithTimeout(r.Ctx, timeout)
		defer cancel()

		sessionList, err := usecase.GetSessionList(ctx, customerID)
		if err != nil {
			return response.NewError(r,
				response.WithStatus(status.InternalServerError),
				response.WithError(err),
			)
		}

		return response.NewSuccess(r, newResponseDTO(sessionList))
	}
}
