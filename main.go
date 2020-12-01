package main

import (
	"github.com/teten-nugraha/mikro-backend/logging"
	route2 "github.com/teten-nugraha/mikro-backend/route"
	"os"
)

func main() {

	arg := os.Args[1]

	logging.InitializeLogging("mikro-backend.log")

	route := route2.Init(arg)

	route.Logger.Fatal(route.Start(":7979"))
}