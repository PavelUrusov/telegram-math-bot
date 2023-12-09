package main

import (
	"flag"
	"fmt"
	"os"
	"telegram-math-bot/internal/app"
)

var appid app.WolframAppID
var apikey app.TelegramAPIKey
var debug bool

func main() {
	application := app.NewApp(appid, apikey, debug)
	application.Run()
}

func init() {
	appidPtr := flag.String("appid", "", "Application ID")
	apikeyPtr := flag.String("apikey", "", "API Key")
	d := flag.Bool("debug", false, "debug mode")

	flag.Parse()

	if *appidPtr == "" || *apikeyPtr == "" {
		fmt.Println("both appid and apikey must be provided")
		os.Exit(1)
	}

	appid = app.WolframAppID(*appidPtr)
	debug = *d
	apikey = app.TelegramAPIKey(*apikeyPtr)
}
