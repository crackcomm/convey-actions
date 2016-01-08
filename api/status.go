package api

import (
	"fmt"
	"net/http"

	"github.com/crackcomm/convey-actions/executor"
)

type StatusHandler struct {
	*executor.Executor
}

func (h StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Executor.Status)
}
