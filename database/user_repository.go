package database

import (
	"backend_time_manager/entity"
	"errors"
	"github.com/google/uuid"
)

func FindUserById(id string) (entity.User, error) {
	query, err := Db.Query("SELECT * FROM TBL_USERS WHERE ID_USER = $1", id)
	if err != nil {
		return entity.User{}, err
	}

	defer query.Close()

	var users []entity.User

	for query.Next() {
		var user entity.User
		if err := query.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Name, &user.Password, &user.Status); err != nil {
			return entity.User{}, err
		}
		if len(users) > 0 {
			return users[0], errors.New("More than one user found")
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return entity.User{}, errors.New("User not found")
	}

	return users[0], nil
}

func FindUserByEmail(email string) (entity.User, error) {
	query, err := Db.Query("SELECT * FROM TBL_USERS WHERE DS_EMAIL = $1", email)
	if err != nil {
		return entity.User{}, err
	}

	defer query.Close()

	var users []entity.User

	for query.Next() {
		var user entity.User
		if err := query.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Name, &user.Password, &user.Status); err != nil {
			return entity.User{}, err
		}
		if len(users) > 0 {
			return users[0], errors.New("More than one user found")
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return entity.User{}, errors.New("User not found")
	}

	return users[0], nil
}

func SaveUser(user entity.User) (entity.User, error) {
	if user.Id == "" {
		user.Id = uuid.New().String()
	}
	return user, nil
}
