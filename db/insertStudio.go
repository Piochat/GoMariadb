package db

import (
	"log"

	"github.com/Piochat/GoMariadb/models"
	_ "github.com/go-sql-driver/mysql" // para experimentar
)

//IDstudio send id to anime routers
var IDstudio int

//InsertStudio returns the id of the inserted studios
func InsertStudio(studio models.Studio) (bool, error) {
	const table = "studios"
	const idTable string = "id_studio"
	const condition string = "name_studio"

	//Checking if the data is already in the table
	idStudio, value := checkStudio(table, condition, studio.Name, idTable)

	if !value {
		stmtIns, err := DbCon.Prepare("INSERT INTO " + table + "(" + condition + ") VALUES (?)")
		if err != nil {
			log.Println("Error Statement Insert [File insertStudio]", err.Error())
			panic(err.Error())
		}
		defer stmtIns.Close()

		_, err = stmtIns.Exec(studio.Name)
		if err != nil {
			log.Println("Error Statement Insert [File insertStudio]", err.Error())
			panic(err.Error())
		}
		idStudio, _ = checkStudio(table, condition, studio.Name, idTable)
	}

	IDstudio = idStudio
	return (IDstudio != 0), nil
}

//Returns the id of the anime found, if not found it returns a false boolean.
func checkStudio(table string, condition string, name string, idTable string) (int, bool) {
	stmtOut, err := DbCon.Prepare("SELECT " + idTable + " FROM " + table + " WHERE " + condition + "= ?")
	if err != nil {
		log.Println("Error Statement Out/Select [File insertStudio]", err.Error())
		panic(err.Error())
	}
	defer stmtOut.Close()

	var idStudio int
	err = stmtOut.QueryRow(name).Scan(&idStudio)
	if err != nil {
		log.Println("Error QueryRow [File insertStudio]", err.Error())
	}

	return idStudio, (idStudio != 0)
}
