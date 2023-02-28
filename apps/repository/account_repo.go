package repository

import (
	"context"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/iqbaldp78/form3-api/apps/lib/thirdparty"
	"github.com/jmoiron/sqlx"
)

type accountIrepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) AccountIRepository {
	return &accountIrepo{
		db: db,
	}
}

func (repo *accountIrepo) FetchAccount(ctx context.Context) error {
	// create request to fetch to form3-api container
	client := resty.New()
	data, err := thirdparty.NewInitForm3(client).RequestFetchData()
	if err != nil {
		return err
	}
	log.Printf("data %+v \n", data)
	return nil
}
