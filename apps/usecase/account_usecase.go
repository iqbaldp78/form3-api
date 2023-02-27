package usecase

import "github.com/jmoiron/sqlx"

type accountUsecase struct {
	Db *sqlx.DB
}

func NewUserMessagesRepo(db *sqlx.DB) repo.IRepoUserMessage {
	return &accountUsecase{
		Db: db,
	}
}
