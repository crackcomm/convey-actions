package api

import "fmt"
import "net/http"

type StatusHandler ApiHandler

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Executor.Status)
}
