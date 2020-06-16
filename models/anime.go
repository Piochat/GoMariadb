package models

//Anime Structure to form the Model
type Anime struct {
	SerieAnime  Serie  `json:"serie"`
	StudioAnime Studio `json:"studio"`
}
