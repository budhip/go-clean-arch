package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		result := map[string]interface{}{
			"message": "Hello, I'm from Go Postgre Clean Arch",
		}
		return c.JSON(http.StatusOK, result)
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}
