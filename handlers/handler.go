package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Piochat/GoMariadb/middlew"

	"github.com/Piochat/GoMariadb/routers"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

//Controllers to endpoint
func Controllers() {
	router := mux.NewRouter()

	//paths
	router.HandleFunc("/animeInsert", middlew.MonitorDB(routers.AnimeInsert)).Methods("POST")
	router.HandleFunc("/animeRead", middlew.MonitorDB(routers.GetAnimes)).Methods("GET")
	router.HandleFunc("/getAnime", middlew.MonitorDB(routers.GetAnAnime)).Methods("GET")
	router.HandleFunc("/modSerie", middlew.MonitorDB(routers.ModifySerie)).Methods("PUT")
	router.HandleFunc("/modStudio", middlew.MonitorDB(routers.ModifyStudio)).Methods("PUT")
	router.HandleFunc("/delSerie", middlew.MonitorDB(routers.DeleteSeries)).Methods("DELETE")
	router.HandleFunc("/delStudio", middlew.MonitorDB(routers.DeleteStudio)).Methods("DELETE")
	//end paths

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8484"
	}

	log.Println("___[PORT]___" + PORT)
	handler := cors.AllowAll().Handler(router)
	log.Fatalln(http.ListenAndServe(":"+PORT, handler))
}
