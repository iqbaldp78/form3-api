package main

import (
	"fmt"

	"github.com/iqbaldp78/form3-api/apps"
	"github.com/spf13/viper"
)

func main() {
	engine := apps.NewEngine()
	engine.Run(fmt.Sprintf("%v", viper.GetString("server.address")))
}
