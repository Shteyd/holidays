package model

import "github.com/Shteyd/holidays/src/service/internal/domain/entity"

type SessionList = []*Session

func ListToEntity(modelList SessionList) []entity.Session {
	entityList := make([]entity.Session, 0, len(modelList))
	for _, session := range modelList {
		entityList = append(entityList, session.ToEntity())
	}

	return entityList
}
