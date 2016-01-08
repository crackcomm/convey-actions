package api

import (
	"encoding/json"
	"net/http"

	"github.com/crackcomm/convey-actions/executor"
)

type LatestHandler struct {
	*executor.Executor
}

func (h LatestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Executor.Latest == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(h.Executor.Latest)
}
