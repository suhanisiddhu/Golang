package authorhttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"io"
	"github.com/gorilla/mux"
	"layeredProject/entities"
	"layeredProject/service"

)

type Handler struct {
	authorService service.AuthorService
}

func New(a service.AuthorService) Handler {
	return Handler{a}
}

func (h Handler) PostAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var author entities.Author

	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &author)

	if err != nil {
		log.Printf("failed %v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	_, err = h.authorService.PostAuthor(ctx, author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("successfully posted!"))
}

func (h Handler) PutAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var author entities.Author

	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &author)

	if err != nil {
		log.Printf("failed %v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	_, err = h.authorService.PutAuthor(ctx, author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "successfully posted!")
}

func (h Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])

	if err != nil || i < 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid id"))

		return
	}

	res, err := h.authorService.DeleteAuthor(ctx, i)
	if err != nil || res == 0 {
		_, _ = w.Write([]byte("can't delete"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted!"))
}
