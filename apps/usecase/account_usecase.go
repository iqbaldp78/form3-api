package usecase

import (
	"context"
	"time"

	"github.com/iqbaldp78/form3-api/apps/repository"
	"github.com/jmoiron/sqlx"
)

type accountUsecase struct {
	Db             *sqlx.DB
	contextTimeout time.Duration
	accountRepo    repository.AccountIRepository
}

func NewAccountUsecase(db *sqlx.DB, accountRepo repository.AccountIRepository, timeout time.Duration) IAccountUsecase {
	return &accountUsecase{
		Db:             db,
		accountRepo:    accountRepo,
		contextTimeout: timeout,
	}
}

func (a *accountUsecase) Fetch(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	a.accountRepo.FetchAccount(ctx)

	// res, nextCursor, err = a.articleRepo.Fetch(ctx, cursor, num)
	// if err != nil {
	// 	return nil, "", err
	// }

	// res, err = a.fillAuthorDetails(ctx, res)
	// if err != nil {
	// 	nextCursor = ""
	// }
	// return
}
