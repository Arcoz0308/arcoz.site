package routes

import (
	"arcoz.dev/api/routes/api"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	api.Init(r.PathPrefix("/api").Subrouter())

	return r
}
