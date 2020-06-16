package db

import (
	"log"

	"github.com/Piochat/GoMariadb/models"
	_ "github.com/go-sql-driver/mysql" // para experimentar
)

//InsertSerie save on db
func InsertSerie(serie models.Serie) (bool, error) {
	const table = "series"
	const idTable string = "id_serie"
	const fields string = "name_serie, episodes_serie, status_serie, type_serie, studio_id"

	stmtIns, err := DbCon.Prepare("INSERT INTO " + table + "(" + fields + ") VALUES (?,?,?,?,?)")
	if err != nil {
		log.Println("Error Statement Insert [File insertSerie]", err.Error())
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(serie.Name, serie.Episodes, serie.Status, serie.Type, IDstudio)
	if err != nil {
		log.Println("Error Statement Insert [File insertSerie]", err.Error())
		panic(err.Error())
	}

	return (err == nil), err
}
