package db

import (
	"github.com/Piochat/GoMariadb/models"
)

//DelStudio Delete the studio in the database
func DelStudio(studio models.Studio) (int64, error) {
	const table = "studios"
	const idTable string = "id_studio"

	result, err := DbCon.Exec("DELETE FROM "+table+" WHERE "+idTable+"=?", studio.ID)
	if err != nil {
		println(result)
		return -1, err
	}
	return result.RowsAffected()
}
