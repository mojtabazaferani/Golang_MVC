package main

import(

	"myapp/routes"
)

func main() {

	routeMain := routes.Route()

	routeMain.Run(":8080")
	
}