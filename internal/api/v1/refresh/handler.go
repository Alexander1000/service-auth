package refresh

import "net/http"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
}