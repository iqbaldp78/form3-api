package main

import (
	"fmt"

	"github.com/iqbaldp78/form3-api/apps"
	"github.com/spf13/viper"
)

func main() {
	engine := apps.NewEngine("config")
	engine.Run(fmt.Sprintf(":%v", viper.GetInt("SERVER_PORT")))
}
