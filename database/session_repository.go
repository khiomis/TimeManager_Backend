package database

import (
	"backend_time_manager/entity"
	"context"
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

func CreateSession(session entity.Session) (entity.Session, error) {
	var query = "INSERT INTO tbl_sessions (id_user, dt_expires_at) VALUES (:id_user,:dt_expires_at)"
	_, err := Db.NamedExecContext(context.Background(), query, session)

	if err != nil {
		return entity.Session{}, err
	}

	savedSession := entity.Session{}
	err = Db.Get(&savedSession, "SELECT * FROM TBL_SESSIONS WHERE id_user = $1 AND dt_expires_at = $2", session.IdUser, session.ExpireAt)

	if err != nil {
		return entity.Session{}, err
	}

	return savedSession, nil
}

func DeleteSession(id uuid.UUID) error {
	// TODO
	return nil
}
