package main

import "github.com/jskaza/article-store/app/routes"

func main() {
	routes.SetupRoutes("./app/ui/views/*.html", "./app/ui/public/css", "./app/ui/public/js", "./app/ui/public/favicon.ico")
}
