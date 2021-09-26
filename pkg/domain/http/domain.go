package http

import (
	"github.com/ducketlab/mingo/http/response"
	"net/http"
)

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "domains")
}
