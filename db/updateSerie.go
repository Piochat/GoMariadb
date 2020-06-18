package db

import (
	"github.com/Piochat/GoMariadb/models"
)

//UpdateSerie Update the series records using the id
func UpdateSerie(serie models.Serie, ID int) (int64, error) {
	const table = "series"
	const idTable string = "id_serie"
	const fields string = "name_serie = ?, episodes_serie = ?, status_serie = ?, type_serie = ?, studio_id = ?"

	result, err := DbCon.Exec("UPDATE "+table+" SET "+fields+" Where "+idTable+"=?", serie.Name, serie.Episodes, serie.Status, serie.Type, serie.StudioID, ID)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
