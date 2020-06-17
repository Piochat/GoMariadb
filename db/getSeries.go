package db

import (
	"log"

	"github.com/Piochat/GoMariadb/models"
)

//GetSeries Returns the series in the database
func GetSeries() []models.Serie {
	const table = "series"

	var serie []models.Serie
	var aux models.Serie

	rows, err := DbCon.Query("SELECT * FROM " + table + " WHERE 1=1")
	if err != nil {
		log.Println("Error Query Out/Select [File getSeries]", err.Error())
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&aux.ID, &aux.Name, &aux.Episodes, &aux.Status, &aux.Type, &aux.StudioID)
		if err != nil {
			panic(err.Error())
		}
		serie = append(serie, aux)
	}

	return serie
}

//GetSerieByID Find by ID and Returns the serie in the database
func GetSerieByID(ID int) models.Serie {
	const table = "series"
	const idTable = "id_serie"

	var aux models.Serie

	rows, err := DbCon.Query("SELECT * FROM "+table+" WHERE "+idTable+"= ?", ID)
	if err != nil {
		log.Println("Error Query Out/Select [File getSeries]", err.Error())
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&aux.ID, &aux.Name, &aux.Episodes, &aux.Status, &aux.Type, &aux.StudioID)
		if err != nil {
			panic(err.Error())
		}
	}

	return aux
}
