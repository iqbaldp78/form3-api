package apps

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iqbaldp78/form3-api/apps/handler"
	"github.com/iqbaldp78/form3-api/apps/repository"
	"github.com/iqbaldp78/form3-api/apps/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func setRouterHandler(router *gin.Engine, db *sqlx.DB) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	accountRepo := repository.NewAccountRepo(db)
	log.Printf("accountRepo %+v \n", accountRepo)
	au := usecase.NewAccountUsecase(db, accountRepo, timeoutContext)
	handler.NewAccountHandler(router, au)
}
