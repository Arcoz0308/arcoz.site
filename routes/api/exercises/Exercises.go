package exercises

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var languages = map[string]string{}

func Init(r *mux.Router) {
	loadLanguages()
	r.HandleFunc("/{lang}/random", getRandom).Methods("GET")
}
func getRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := mux.Vars(r)
	lang := languages[strings.ToLower(v["lang"])]
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		response := make(map[string]interface{})
		response["message"] = "Unknown language"
		response["code"] = 10001
		json.NewEncoder(w).Encode(response)
	}
}
