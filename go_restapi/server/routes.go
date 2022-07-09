package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {

	//index es la funcion definida el el handlers.go
	http.HandleFunc("/", index)

	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			getCountries(w, r)

		case http.MethodPost:
			addCountries(w, r)
		
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w,"Method not allowed")
			return
		
		}
		

	})

}
