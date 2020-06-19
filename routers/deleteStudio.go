package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Piochat/GoMariadb/db"
)

//DeleteStudio delete the studio
func DeleteStudio(w http.ResponseWriter, r *http.Request) {
	idStudio := r.URL.Query().Get("idStudio")
	if len(idStudio) < 1 {

		log.Println("Error id empty [file deleteStudio]")
		http.Error(w, "Error id empty", http.StatusBadRequest)
		return
	}

	valueID, err := strconv.Atoi(idStudio)
	if err != nil {
		log.Println("Error id invalid [file deleteStudio]")
		http.Error(w, "Error id invalid "+err.Error(), http.StatusBadRequest)
		return
	}

	studio := db.GetStudioByID(valueID)
	_, err = db.DelStudio(studio)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
