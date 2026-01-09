package shortener

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Shorten(res http.ResponseWriter, req *http.Request) {
	var reqBody struct {
		URL string `json:"url"`
	}

	json.NewDecoder(req.Body).Decode(&reqBody)

	shortCode, error := h.service.Shorten(reqBody.URL)
	if error != nil {
		http.Error(res, error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(map[string]string{"short": shortCode})
}

func (h *Handler) Resolve() {}

func (h *Handler) Test(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(map[string]string{"success": "true"})
}
