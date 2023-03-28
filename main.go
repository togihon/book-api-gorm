package main

import "book-gorm/app/handler"

func main() {
	var PORT = ":8080"

	handler.StartServer().Run(PORT)
}
