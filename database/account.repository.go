package database

import (
	"backend_time_manager/entity"
	"errors"
	"github.com/google/uuid"
)

func FindAccountById(id string) (entity.Account, error) {
	query, err := Db.Query("SELECT * FROM TBL_ACCOUNTS WHERE ID_ACCOUNT = $1", id)
	if err != nil {
		return entity.Account{}, err
	}

	defer query.Close()

	var accounts []entity.Account

	for query.Next() {
		var account entity.Account
		if err := query.Scan(&account.Id, &account.CreatedAt, &account.UpdatedAt, &account.Email, &account.Name, &account.Password, &account.Status); err != nil {
			return entity.Account{}, err
		}
		if len(accounts) > 0 {
			return accounts[0], errors.New("More than one account found")
		}
		accounts = append(accounts, account)
	}

	if len(accounts) == 0 {
		return entity.Account{}, errors.New("Account not found")
	}

	return accounts[0], nil
}

func FindAccountByEmail(email string) (entity.Account, error) {
	query, err := Db.Query("SELECT * FROM TBL_ACCOUNTS WHERE DS_EMAIL = $1", email)
	if err != nil {
		return entity.Account{}, err
	}

	defer query.Close()

	var accounts []entity.Account

	for query.Next() {
		var account entity.Account
		if err := query.Scan(&account.Id, &account.CreatedAt, &account.UpdatedAt, &account.Email, &account.Name, &account.Password, &account.Status); err != nil {
			return entity.Account{}, err
		}
		if len(accounts) > 0 {
			return accounts[0], errors.New("More than one account found")
		}
		accounts = append(accounts, account)
	}

	if len(accounts) == 0 {
		return entity.Account{}, errors.New("Account not found")
	}

	return accounts[0], nil
}

func SaveAccount(account entity.Account) (entity.Account, error) {
	if account.Id == "" {
		account.Id = uuid.New().String()
	}
	return account, nil
}
