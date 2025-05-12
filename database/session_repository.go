package database

import (
	"backend_time_manager/entity"
	"errors"
	"github.com/google/uuid"
)

func FindSessionByUuid(id uuid.UUID) (entity.Session, error) {
	var sessions []entity.Session
	err := Db.Select(&sessions, "SELECT * FROM TBL_SESSIONS WHERE id_session = $1", id)
	if err != nil {
		return entity.Session{}, err
	}

	if sessions == nil || len(sessions) == 0 {
		return entity.Session{}, errors.New("Session not found")
	}

	if len(sessions) > 1 {
		return entity.Session{}, errors.New("Multiple sessions found")
	}

	return sessions[0], nil
}

func DeleteSession(id uuid.UUID) error {
	// TODO
	return nil
}
