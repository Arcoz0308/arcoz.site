package api

import (
	"arcoz.dev/api/routes/api/exercises"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Init(r *mux.Router) {
	exercises.Init(r.PathPrefix("/exercises").Subrouter())
}
func SendInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	r := make(map[string]interface{})
	r["message"] = "an internal error happened"
	r["code"] = 0
	json.NewEncoder(w).Encode(r)
}
