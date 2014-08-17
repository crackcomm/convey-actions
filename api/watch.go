package api

import "fmt"
import "net/http"

type WatchHandler ApiHandler

func (h *WatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.Executor.Dirname = first(r.URL.Query()["root"])
	}
	fmt.Fprint(w, h.Executor.Dirname)
	h.Executor.RunEvent()
}

func first(in []string) string {
	if len(in) <= 0 {
		return ""
	}
	return in[0]
}
