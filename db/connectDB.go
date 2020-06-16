package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // para experimentar
)

const (
	user          = "piocha"
	pss           = "koke"
	dbname        = "animes"
	stringConnect = user + ":" + pss + "@/" + dbname
)

//DbCon Var to connect and use in other clases
var DbCon = Connection()

//Connection to database
func Connection() *sql.DB {
	db, err := sql.Open("mysql", stringConnect)
	if err != nil {
		log.Fatalln("Error pack db; file connectDB", err.Error())
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	println("Connecction Vo")
	return db
}
