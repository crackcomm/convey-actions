package api

import "net/http"

type ExecuteHandler ApiHandler

func (h *ExecuteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Executor.Run()
	w.WriteHeader(http.StatusOK)
}
