package main

import (
	db "webServer/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8889")
}
