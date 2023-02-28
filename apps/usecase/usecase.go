package usecase

import "context"

type IAccountUsecase interface {
	Fetch(ctx context.Context)
}
