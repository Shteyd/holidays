package getListHandler

import "github.com/Shteyd/holidays/src/service/internal/domain/entity"

type responseDTOList []responseDTO

type responseDTO struct {
	ID        int64  `json:"id"`
	UserAgent string `json:"user_agent"`
	ExpiresAt int64  `json:"expires_at"`
}

func newResponseDTO(sessionList []entity.Session) responseDTOList {
	data := make(responseDTOList, 0, len(sessionList))
	for _, session := range sessionList {
		data = append(data, responseDTO{
			ID:        session.ID,
			UserAgent: session.UserAgent,
			ExpiresAt: session.ExpiresAt.UTC().Unix(),
		})
	}

	return data
}
