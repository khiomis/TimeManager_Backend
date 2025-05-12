package database

import (
	"backend_time_manager/entity"
	"context"
	"github.com/google/uuid"
)

func FindToken(id uuid.UUID, tokenCode string, tokenType entity.ValidationTokenType) (entity.ValidationToken, error) {
	var token entity.ValidationToken
	err := Db.Get(&token, "SELECT * FROM tbl_validation_tokens WHERE id_validation_token = $1 AND cd_validation_token = $2 AND tp_validation_token = $3", id, tokenCode, tokenType)
	if err != nil {
		return entity.ValidationToken{}, err
	}

	return token, nil
}

func InsertToken(token entity.ValidationToken) (entity.ValidationToken, error) {
	var err error

	var query = "INSERT INTO tbl_validation_tokens(cd_validation_token, dt_expire_at, tp_validation_token, id_user) values (:cd_validation_token, :dt_expire_at, :tp_validation_token, :id_user)"
	_, err = Db.NamedExecContext(context.Background(), query, token)

	if err != nil {
		return entity.ValidationToken{}, err
	}

	var tokenSaved entity.ValidationToken
	err = Db.Get(&tokenSaved, "SELECT * FROM tbl_validation_tokens WHERE cd_validation_token = $1 AND tp_validation_token = $2 AND dt_expire_at = $3", token.Code, token.Type, token.ExpireAt)

	if err != nil {
		return entity.ValidationToken{}, err
	}

	return tokenSaved, nil
}

func RemoveToken(id uuid.UUID) error {
	_, err := Db.Exec("DELETE FROM tbl_validation_tokens WHERE id_validation_token = $1", id)

	if err != nil {
		return err
	}
	return nil
}
