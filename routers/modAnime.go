package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Piochat/GoMariadb/db"

	"github.com/Piochat/GoMariadb/models"
)

//ModifySerie Modify Serie information
func ModifySerie(w http.ResponseWriter, r *http.Request) {
	var serie models.Serie

	idSerie := r.URL.Query().Get("idSerie")

	if len(idSerie) < 1 {
		http.Error(w, "Error ID Empty", http.StatusBadRequest)
		return
	}

	valueIDSerie, err := strconv.Atoi(idSerie)
	if err != nil {
		http.Error(w, "The id is not a number", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&serie)
	if err != nil {
		http.Error(w, "Data Wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.UpdateSerie(serie, valueIDSerie)
	if err != nil {
		http.Error(w, "Error trying to update information "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
