package repository

import (
	"github.com/jmoiron/sqlx"
)

type AccountIrepo struct {
}

func NewAccountRepo(db *sqlx.DB) *AccountIrepo {
	return &AccountIrepo{}
}
