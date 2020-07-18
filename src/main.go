package main

import (
	"fmt"
	"strings"

	"git.darknebu.la/chaosdorf/freitagsfoo/src/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initLogging()

	pgdb := db.Connect()
	defer db.Disconnect(pgdb)

	initDB(pgdb)

	initHTTPServer()
}

func initConfig() {
	logrus.Info("Init config")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/go/src/app/")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// var conf structs.Configuration
	// err = viper.Unmarshal(&conf)
	// if err != nil {
	// 	log.Fatal("unable to decode into struct: %v", err)
	// }
}

func initLogging() {
	logrus.Info("Init logging")
	switch viper.GetString("verbose.level") {
	case "1":
		logrus.SetLevel(logrus.PanicLevel)
	case "2":
		logrus.SetLevel(logrus.FatalLevel)
	case "3":
		logrus.SetLevel(logrus.ErrorLevel)
	case "4":
		logrus.SetLevel(logrus.WarnLevel)
	case "5":
		logrus.SetLevel(logrus.InfoLevel)
	case "6":
		logrus.SetLevel(logrus.DebugLevel)
	case "7":
		logrus.SetLevel(logrus.TraceLevel)
	}
}
