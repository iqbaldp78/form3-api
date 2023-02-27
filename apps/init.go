package apps

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaldp78/form3-api/apps/lib/db"
	"github.com/spf13/viper"
)

/*NewEngine used for create new engine*/
func NewEngine(config string) *gin.Engine {
	setConfig(config)
	router := gin.Default()
	db.New(
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"))
	setRouterHandler(router)
	setErrorHandler(router)
	return router
}

//setConfig used for setup application configuration
func setConfig(config string) {
	viper.SetConfigName(config)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("configuration file not found: %s", err))
	}
}

/*setErrorHandler used for handling 404 request*/
func setErrorHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		msg := map[string]interface{}{"message": "Page not found", "status": http.StatusNotFound}
		c.AbortWithStatusJSON(http.StatusNotFound, msg)
	})
}
