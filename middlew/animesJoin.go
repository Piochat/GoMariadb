package middlew

import (
	"github.com/Piochat/GoMariadb/models"
)

//PassAnime To fill a slice with values from other slices
func PassAnime(series []models.Serie, studios []models.Studio) []models.Anime {
	var animes []models.Anime
	var aux models.Anime

	for i := 0; i < len(series); i++ {
		for j := 0; j < len(studios); j++ {
			aux.SerieAnime = series[i]
			if aux.SerieAnime.StudioID == studios[j].ID {
				aux.StudioAnime = studios[j]
			}
		}
		animes = append(animes, aux)
	}

	return animes
}

//PassOneAnime Generate a single anime with the received
func PassOneAnime(serie models.Serie, studio models.Studio) models.Anime {
	return models.Anime{
		StudioAnime: studio,
		SerieAnime:  serie,
	}
}
