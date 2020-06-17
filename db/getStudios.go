package db

import (
	"log"

	"github.com/Piochat/GoMariadb/models"
)

//GetStudios Returns the studios in the database
func GetStudios() []models.Studio {
	const table = "studios"
	var studio []models.Studio
	var aux models.Studio

	rows, err := DbCon.Query("SELECT * FROM " + table + " WHERE 1=1")
	if err != nil {
		log.Println("Error Query Out/Select [File getStudios]", err.Error())
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&aux.ID, &aux.Name)
		if err != nil {
			panic(err.Error())
		}
		studio = append(studio, aux)
	}

	return studio
}
