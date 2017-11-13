package router

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/e-capture/ECMConfiguration/controllers/ImportconfigController"
	"github.com/e-capture/ECMRetencionDocumental/controllers/RetencionController"
)

var _port string
var _serviceName string

func init()  {
	if err := godotenv.Load(); err != nil {
		panic(err)
		log.Fatal("Error loading .env file")
	}
	_port = os.Getenv("APP_PORT")
	_serviceName = os.Getenv("APP_SERVICE_NAME")
}

func StartService()  {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/api/v1/import", ImportconfigController.Import)
	e.POST("/api/v1/importconfig", ImportconfigController.ImportConfig)
	e.GET("/api/v1/roles/index", RetencionController.Index)
	e.Logger.Fatal(e.Start(":"+_port))
}






