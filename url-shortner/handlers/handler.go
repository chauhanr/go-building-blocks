package handlers

import "net/http"

func MapHandler(pathTourls map[string]string, fallback http.Handler) http.HandlerFunc {

	mapHandler := func(w http.ResponseWriter, r *http.Request) {

	}

	return mapHandler
}
