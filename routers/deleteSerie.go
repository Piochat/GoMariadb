package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Piochat/GoMariadb/db"
)

//DeleteSeries Delete the series
func DeleteSeries(w http.ResponseWriter, r *http.Request) {

	idSerie := r.URL.Query().Get("idSerie")
	if len(idSerie) < 1 {

		log.Println("Error id empty [file deleteSerie]")
		http.Error(w, "Error id empty", http.StatusBadRequest)
		return
	}

	valueID, err := strconv.Atoi(idSerie)
	if err != nil {
		log.Println("Error id invalid [file deleteSerie]")
		http.Error(w, "Error id invalid "+err.Error(), http.StatusBadRequest)
		return
	}

	serie := db.GetSerieByID(valueID)
	if serie.ID < 1 {
		log.Println("Error id invalid [file deleteSerie]")
		http.Error(w, "Error id invalid "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DelSerie(serie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
