package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()
	
	// r := mux.NewRouter()

	// routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// fmt.Println("server running localhost:5000")
	// http.ListenAndServe("localhost:5000", r)
}

