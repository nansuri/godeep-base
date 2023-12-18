package core

import (
	"fmt"
	"os"

	enum "github.com/nansuri/godeep-base/domain/enum/base"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Intiatior() {
	switch os.Getenv("APP_ENV") {
	case enum.EnvDevelopment:
		viper.SetConfigName(enum.EnvDevelopment)
	case enum.EnvStaging:
		viper.SetConfigName(enum.EnvStaging)
	case enum.EnvProduction:
		viper.SetConfigName(enum.EnvProduction)
	default:
		viper.SetConfigName(enum.EnvStaging)
	}

	viper.SetConfigType("yaml")

	readConfErr := viper.ReadInConfig()
	if readConfErr != nil {
		logrus.Error(readConfErr.Error())
		logrus.Fatal("=== Loading Config ERROR ===")
		return
	}
	fmt.Println("====== 🅂 🅃 🄰 🅁 🅃 🄸 🄽 🄶 ======\n\n" +
		"░█████╗░░██████╗██████╗░░░░░░░██████╗░░█████╗░░██████╗███████╗\n" +
		"██╔══██╗██╔════╝██╔══██╗░░░░░░██╔══██╗██╔══██╗██╔════╝██╔════╝\n" +
		"██║░░██║╚█████╗░██████╔╝█████╗██████╦╝███████║╚█████╗░█████╗░░\n" +
		"██║░░██║░╚═══██╗██╔═══╝░╚════╝██╔══██╗██╔══██║░╚═══██╗██╔══╝░░\n" +
		"╚█████╔╝██████╔╝██║░░░░░░░░░░░██████╦╝██║░░██║██████╔╝███████╗\n" +
		"░╚════╝░╚═════╝░╚═╝░░░░░░░░░░░╚═════╝░╚═╝░░╚═╝╚═════╝░╚══════╝\n" +
		os.Getenv("APP_NAME") + " Version: " + os.Getenv("APP_VERSION") + "-" + os.Getenv("APP_ENV") + "\n" +
		"Copyright© 2023 Nansuri.\n\n")
}
