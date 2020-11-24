package main

import route2 "github.com/teten-nugraha/mikro-backend/route"

func main() {

	route := route2.Init()

	route.Logger.Fatal(route.Start(":7979"))
}