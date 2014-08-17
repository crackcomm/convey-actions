package api

import "time"
import "strconv"
import "net/http"

type PollHandler ApiHandler

func (h *PollHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
	if err != nil || timeout > 180000 || timeout < 0 {
		timeout = 60000 // default timeout is 60 seconds
	}

	if h.Executor.Changed {
		h.Executor.Changed = false
		if h.Executor.Latest == nil {
			h.Executor.Run()
		}
		w.Write([]byte("idle"))
		return
	}

	select {
	case <-h.Executor.Events():
		w.Write([]byte("idle"))
	case <-time.After(time.Duration(timeout) * time.Millisecond):
		w.Write([]byte("idle"))
	}
}
