package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	dbConn "github.com/budhip/go-postgre-clean-arch/db"
	mWare "github.com/budhip/go-postgre-clean-arch/middleware"
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
	var stageEnv string

	err := godotenv.Load()
	if err != nil {
		// if stage not local
		fmt.Println("loading from os")
	}

	stageEnv = os.Getenv("GO_ENV")

	// connect to db
	db, errConn := dbConn.ConnectToDB(stageEnv)
	if errConn != nil {
		panic(errConn)
	}

	defer db.Close()

	fmt.Println("Successfully connected!")

	e := echo.New()
	customMware := mWare.InitMiddleware()
	e.Use(customMware.CheckAuthHeader)

	e.GET("/ping", func(c echo.Context) error {
		result := map[string]interface{}{
			"message": "Hello, I'm from Go Postgre Clean Arch",
		}
		return c.JSON(http.StatusOK, result)
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}
