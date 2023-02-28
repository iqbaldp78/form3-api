package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaldp78/form3-api/apps/usecase"
)

// AccountHandler  represent the httphandler for account
type AccountHandler struct {
	AccUsecase usecase.IAccountUsecase
}

// NewAccountHandler will initialize the account/resources endpoint
func NewAccountHandler(router *gin.Engine, us usecase.IAccountUsecase) {
	handler := &AccountHandler{
		AccUsecase: us,
	}
	api := router.Group("accounts")
	api.GET("/", handler.FetchAccount) // used for get account

}

// FetchAccount will fetch the Account based on given params
func (a *AccountHandler) FetchAccount(c *gin.Context) {
	ctx := context.Background()
	a.AccUsecase.Fetch(ctx)
	c.JSON(http.StatusOK, "data")
}
