package apps

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaldp78/form3-api/apps/lib/db"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

/*NewEngine used for create new engine*/
func NewEngine() *gin.Engine {
	// setConfig(config)
	router := gin.Default()
	db.New(
		viper.GetString("database.name"),
		viper.GetString("database.user"),
		viper.GetString("database.pass"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"))
	setRouterHandler(router, db.GetDB())
	setErrorHandler(router)
	return router
}

/*setErrorHandler used for handling 404 request*/
func setErrorHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		msg := map[string]interface{}{"message": "Page not found", "status": http.StatusNotFound}
		c.AbortWithStatusJSON(http.StatusNotFound, msg)
	})
}
