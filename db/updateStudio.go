package db

import (
	"github.com/Piochat/GoMariadb/models"
)

//UpdateStudio Update the studios records using the id
func UpdateStudio(studio models.Studio, ID int) (int64, error) {
	const table = "studios"
	const idTable string = "id_studio"
	const condition string = "name_studio = ?"

	result, err := DbCon.Exec("UPDATE "+table+" SET "+condition+" Where "+idTable+"=?", studio.Name, ID)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
