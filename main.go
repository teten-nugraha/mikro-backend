package main

import (
	route2 "github.com/teten-nugraha/mikro-backend/route"
	"os"
)

func main() {

	arg := os.Args[1]

	route := route2.Init(arg)

	route.Logger.Fatal(route.Start(":7979"))
}