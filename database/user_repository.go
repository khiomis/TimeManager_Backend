package database

import (
	"backend_time_manager/entity"
	"context"
	"errors"
	"time"
)

func FindUserById(id int64) (entity.User, error) {
	var users []entity.User
	err := Db.Select(&users, "SELECT * FROM TBL_USERS WHERE ID_USER = $1", id)
	if err != nil {
		return entity.User{}, err
	}

	if users == nil || len(users) == 0 {
		return entity.User{}, errors.New("User not found")
	}

	if len(users) > 1 {
		return entity.User{}, errors.New("Multiple users found")
	}

	return users[0], nil
}

func FindUserByEmail(email string) (entity.User, error) {
	var users []entity.User
	err := Db.Select(&users, "SELECT * FROM TBL_USERS WHERE DS_EMAIL = $1", email)
	if err != nil {
		return entity.User{}, err
	}

	if users == nil || len(users) == 0 {
		return entity.User{}, errors.New("User not found")
	}

	if len(users) > 1 {
		return entity.User{}, errors.New("Multiple users found")
	}

	return users[0], nil
}

func CheckEmailAlreadyInUseUser(email string) (bool, error) {
	var users []entity.User
	result, err := Db.Queryx("SELECT * FROM TBL_USERS WHERE DS_EMAIL = $1", email)
	if err != nil {
		return false, err
	}
	for result.Next() {
		var user entity.User
		err = result.StructScan(&user)
		if err != nil {
			return false, err
		}
		users = append(users, user)
	}
	result.Close()

	if users == nil || len(users) == 0 {
		return false, nil
	}

	return true, nil
}

func SaveUser(user entity.User) (entity.User, error) {
	isInsert := user.Id <= 0

	timeOperation := time.Now()

	if isInsert {
		user.CreatedAt = timeOperation
	}
	user.UpdatedAt = timeOperation

	var err error
	if isInsert {
		var query = "INSERT INTO TBL_USERS (DT_CREATED_AT, DT_UPDATED_AT, NM_USER, DS_EMAIL, DS_PASSWORD, TP_STATUS) " +
			"VALUES (current_timestamp,current_timestamp,:nm_user,:ds_email,:ds_password,:tp_status)"
		_, err = Db.NamedExecContext(context.Background(), query, user)
	} else {
		var query = "UPDATE TBL_USERS " +
			"SET DB_UPDATED_AT = :DB_UPDATED_AT, DS_EMAIL = :DS_EMAIL, NM_USER = :NM_USER, NM_PASSWORD = :NM_PASSWORD, TP_STATUS = :TP_STATUS" +
			"WHERE ID_USER = :ID_USER"
		_, err = Db.NamedExecContext(context.Background(), query, user)
	}

	if err != nil {
		return entity.User{}, err
	}

	return FindUserByEmail(user.Email)
}
