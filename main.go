package main

import (
	"flag"

	"github.com/teten-nugraha/mikro-backend/logging"
	route2 "github.com/teten-nugraha/mikro-backend/route"
)

func main() {
	flag.Parse()
	arg := flag.Args()

	logging.InitializeLogging("mikro-backend.log")

	route := route2.Init(arg)

	route.Logger.Fatal(route.Start(":7979"))
}
