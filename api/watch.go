package api

import "fmt"
import "net/http"

type WatchHandler ApiHandler

func (h *WatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Executor.Dirname)
}
