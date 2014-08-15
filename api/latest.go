package api

import "net/http"
import "encoding/json"

type LatestHandler ApiHandler

func (h *LatestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Executor.Run())
}
