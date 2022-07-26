package authorhttp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"layeredProject/entities"
	"layeredProject/service"
	"log"
	"net/http"
	"strconv"
)

type authorHandler struct {
	authorService service.AuthorService
}

func New(a service.AuthorService) authorHandler {
	return authorHandler{a}
}

func (h authorHandler) PostAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &author)
	if err != nil {
		log.Printf("failed %v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	_, err = h.authorService.PostAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("successfully posted!"))
}

func (h authorHandler) PutAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &author)
	if err != nil {
		log.Printf("failed %v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	_, err = h.authorService.PutAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "successfully posted!")
}

func (h authorHandler) DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		_, _ = w.Write([]byte("invalid typed id"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.authorService.DeleteAuthor(i)
	if err != nil || res == 0 {
		_, _ = w.Write([]byte("can't delete"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted!"))
}
