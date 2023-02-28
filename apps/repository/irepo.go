package repository

import "context"

type AccountIRepository interface {
	FetchAccount(ctx context.Context) error
}
