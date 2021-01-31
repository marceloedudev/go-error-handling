package main

import "go-error-handling/routes"

func main() {

	app := routes.AddRoutes()
	port := "3000"

	app.Run(":" + port)

}
