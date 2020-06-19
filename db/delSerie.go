package db

import (
	"github.com/Piochat/GoMariadb/models"
)

//DelSerie Delete the series in the database
func DelSerie(serie models.Serie) (int64, error) {
	const table = "series"
	const idTable string = "id_serie"

	result, err := DbCon.Exec("DELETE FROM "+table+" WHERE "+idTable+"= ?", serie.ID)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
