package main

import (
	"os"

	"github.com/nansuri/godeep-base/infrastructure/auth"
	clients "github.com/nansuri/godeep-base/infrastructure/clients/osp"
	"github.com/nansuri/godeep-base/infrastructure/persistence"
	"github.com/nansuri/godeep-base/interfaces/handlers/base"
	"github.com/nansuri/godeep-base/interfaces/middleware"
	routers "github.com/nansuri/godeep-base/interfaces/routers/base"
	"github.com/nansuri/godeep-base/utils/core"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	//Initiate default env loader
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("=== Loading Config ERROR ===")
	}

	logrus.Info("Loading Configs.....\n")
	viper.AddConfigPath("configs")

	core.Intiatior()
}

func main() {
	app_mode := os.Getenv("APP_ENV")

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//redis details
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	// define all services here
	services, dbErr := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if dbErr != nil {
		logrus.Error("Database error: ", dbErr.Error())
	}
	defer services.Close()
	services.Automigrate()

	_, cacheErr := auth.NewRedisDB(redis_host, redis_port, redis_password)
	if cacheErr != nil {
		logrus.Error("Redis error: ", cacheErr)
	}

	userService := clients.NewOspUserService()

	// put all services here
	baseService := base.NewBase(services.Base, userService)

	if app_mode == "" {
		app_mode = "debug"
	}
	gin.SetMode(app_mode)
	r := gin.New()
	r.Use(middleware.CORSHeaderSet(), middleware.TracerInterceptor(), middleware.LoggerMiddleman())
	v1 := r.Group("/v1")

	// Routers will be here
	routers.BaseRouter(v1, baseService)

	//Starting the application
	app_port := os.Getenv("APP_PORT") //using heroku host
	if app_port == "" {
		app_port = "8088" //localhost
	}
	logrus.Fatal(r.Run(":" + app_port))

}
