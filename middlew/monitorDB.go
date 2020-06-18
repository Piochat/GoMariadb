package middlew

import (
	"net/http"

	"github.com/Piochat/GoMariadb/db"
)

//MonitorDB check ping with db
func MonitorDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.DbCon.Ping() != nil {
			http.Error(w, "Lost connection to the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
