package db

import (
	"log"

	"github.com/Piochat/GoMariadb/models"
)

//SerieStudioID Varible para cargar el id de los studios en las series
//var SerieStudioID int

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

// func GetSeriesByID(ID int) {
// }
