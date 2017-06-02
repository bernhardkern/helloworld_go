package rest

import (
	"github.com/bernhardkern/helloworld_go/data"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo data.PersonRepository
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	persons, err := h.repo.ReadAll();
	if err != nil {
		fmt.Fprintf(w, "error while reading all persons", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	person := &data.Person{}
	json.NewDecoder(r.Body).Decode(person)
	fmt.Println(person)

	err := h.repo.Create(*person)
	if err != nil {
		fmt.Println("error at create after post", err)
	}
}

func NewRestHandler(repo data.PersonRepository) Handler {
	return Handler{repo: repo}
}
