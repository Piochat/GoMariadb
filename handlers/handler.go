package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Piochat/GoMariadb/routers"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

//Controllers to endpoint
func Controllers() {
	router := mux.NewRouter()

	//paths
	router.HandleFunc("/animeInsert", routers.AnimeInsert).Methods("POST")
	router.HandleFunc("/animeRead", routers.GetAnimes).Methods("GET")
	router.HandleFunc("/getAnime", routers.GetAnAnime).Methods("GET")
	//end paths

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8484"
	}

	log.Println("___[PORT]___" + PORT)
	handler := cors.AllowAll().Handler(router)
	log.Fatalln(http.ListenAndServe(":"+PORT, handler))
}
