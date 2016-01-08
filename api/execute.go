package api

import (
	"net/http"

	"github.com/crackcomm/convey-actions/executor"
)

type ExecuteHandler struct {
	*executor.Executor
}

func (h ExecuteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Executor.Run()
	w.WriteHeader(http.StatusOK)
}
