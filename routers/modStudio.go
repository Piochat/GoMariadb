package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Piochat/GoMariadb/db"
	"github.com/Piochat/GoMariadb/models"
)

//ModifyStudio Modify Studio information
func ModifyStudio(w http.ResponseWriter, r *http.Request) {
	var studio models.Studio

	idStudio := r.URL.Query().Get("idStudio")

	if len(idStudio) < 1 {
		http.Error(w, "Error ID Empty", http.StatusBadRequest)
		return
	}

	valueIDStudio, err := strconv.Atoi(idStudio)
	if err != nil {
		http.Error(w, "The id is not a number", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&studio)
	if err != nil {
		http.Error(w, "Data Wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.UpdateStudio(studio, valueIDStudio)
	if err != nil {
		http.Error(w, "Error trying to update information "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
