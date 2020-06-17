package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Piochat/GoMariadb/middlew"

	"github.com/Piochat/GoMariadb/db"

	"github.com/Piochat/GoMariadb/models"
)

//GetAnimes Controlador para retornar un objeto anime
func GetAnimes(w http.ResponseWriter, r *http.Request) {
	var animes []models.Anime
	var series []models.Serie
	var studios []models.Studio

	series = db.GetSeries()
	studios = db.GetStudios()
	animes = middlew.PassAnime(series, studios)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(animes)
}
