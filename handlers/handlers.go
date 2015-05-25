package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterHandlers interface {
	GetPathPrefix() string
	InitRouter(r *mux.Router)
}

func InitHandlers() {
	handlers := []RouterHandlers{MainHandlers{}}

	r := mux.NewRouter()
	for _, handler := range handlers {
		pathPrefix := handler.GetPathPrefix()
		if pathPrefix != "" {
			s := r.PathPrefix(pathPrefix).Subrouter()
			handler.InitRouter(s)
		} else {
			handler.InitRouter(r)
		}
	}

	http.HandleFunc("/static/", func(rw http.ResponseWriter, res *http.Request) {
		http.ServeFile(rw, res, res.URL.Path[1:])
	})

	// The following register router, so the router will be enabled
	http.Handle("/", r)
}
