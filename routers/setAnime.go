package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Piochat/GoMariadb/db"

	"github.com/Piochat/GoMariadb/models"
)

//AnimeInsert controller to Insert a Anime in DB
func AnimeInsert(w http.ResponseWriter, r *http.Request) {
	var t models.Anime
	var status bool

	//Receive the Body information and decode it in the Anime entity
	//We verify errors
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println("Error Routers [File setStudio]")
		http.Error(w, "Error DATA Wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	//We check the values inside Anime t
	if len(t.StudioAnime.Name) < 1 {
		http.Error(w, "Error Name Empty", http.StatusBadRequest)
		return
	}

	//Studio
	status, err = db.InsertStudio(t.StudioAnime)
	if err != nil {
		log.Println("Error Routers [File setStudio]")
		http.Error(w, "Error Inserting Studio  "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Error DATA Wrong (DIE)", http.StatusBadRequest)
		return
	}
	//End Studio

	//Serie
	status, err = db.InsertSerie(t.SerieAnime)
	if err != nil {
		log.Println("Error Routers [File setStudio]")
		http.Error(w, "Error Inserting Studio  "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Error DATA Wrong (DIE)", http.StatusBadRequest)
		return
	}
	//End Serie

	w.WriteHeader(http.StatusCreated)
}
