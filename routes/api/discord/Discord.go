package discord

import (
	"arcoz.dev/api/routes/api"
	"arcoz.dev/api/utils"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Init(r *mux.Router) {
	r.HandleFunc("/users/{user}", getUser)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := &http.Client{}
	v := mux.Vars(r)
	user := v["user"]
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/"+user, nil)
	if err != nil {
		fmt.Print(err)
		api.SendInternalError(w)
		return
	}
	req.Header.Add("Authorization", "Bot "+utils.Config.Token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		api.SendInternalError(w)
		return
	}
	switch response.StatusCode {
	case http.StatusOK:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			api.SendInternalError(w)
			fmt.Print(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(body))
		return

	case http.StatusUnauthorized, http.StatusBadRequest:
		api.SendInternalError(w)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Print(string(body))
		return
	case http.StatusNotFound:

	case http.StatusForbidden:

	}
}
