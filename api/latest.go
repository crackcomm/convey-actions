package api

import "net/http"
import "encoding/json"

type LatestHandler ApiHandler

func (h *LatestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Executor.Latest == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(h.Executor.Latest)
}
