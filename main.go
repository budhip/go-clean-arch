package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
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
