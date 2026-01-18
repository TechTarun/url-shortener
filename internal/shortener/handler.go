package shortener

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Shorten(res http.ResponseWriter, req *http.Request) {
	var reqBody struct {
		URL             string `json:"url"`
		CustomShortCode string `json:"short_code"`
	}

	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortCode, err := h.service.Shorten(reqBody.URL, reqBody.CustomShortCode)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(map[string]string{"short": shortCode})
}

func (h *Handler) Resolve(res http.ResponseWriter, req *http.Request) {
	shortCode := strings.TrimPrefix(req.URL.Path, "/")
	if shortCode == "" {
		http.NotFound(res, req)
		return
	}

	longUrl, err := h.service.Resolve(shortCode)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if longUrl == "" {
		http.NotFound(res, req)
		return
	}

	http.Redirect(res, req, longUrl, http.StatusFound)
}

func (h *Handler) Test(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(map[string]string{"message": "testinggggg"})
}
