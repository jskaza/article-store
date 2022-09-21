package main

import "github.com/jskaza/open-journal/app/routes"

func main() {
	routes.SetupRoutes("./app/ui/views/*.html", "./app/ui/public/css", "./app/ui/public/js")
}
