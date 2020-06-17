package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Piochat/GoMariadb/db"
	"github.com/Piochat/GoMariadb/middlew"
)

//GetAnAnime Controller to get an anime
func GetAnAnime(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		log.Println("Error id empty [file getAnAnime]")
		http.Error(w, "Error id empty", http.StatusBadRequest)
		return
	}

	serieID, _ := strconv.Atoi(ID)
	serie := db.GetSerieByID(serieID)
	studio := db.GetStudioByID(serie.StudioID)
	anime := middlew.PassOneAnime(serie, studio)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(anime)
}
