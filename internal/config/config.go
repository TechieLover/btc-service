package config

import (
	"os"
)

var (
	APP_NAME string
	PORT     string
	CURRENCY string
)

func init() {
	APP_NAME = os.Getenv("APP_NAME")
	PORT = os.Getenv("PORT")
	CURRENCY = os.Getenv("CURRENCY")
}
