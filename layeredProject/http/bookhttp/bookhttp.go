package bookhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"io"
	"github.com/gorilla/mux"
	"layeredProject/entities"
	"layeredProject/service"

)

type Handler struct {
	bookH service.BookService
}

func New(bookS service.BookService) Handler {
	return Handler{bookS}
}

func (h Handler) GetAllBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	x := r.URL.Query()
	title := x.Get("title")
	includeAuthor := x.Get("includeAuthor")
	books, err := h.bookH.GetAllBook(ctx, title, includeAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(books)
	if err != nil {
		_, _ = w.Write([]byte("can't encode"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (h Handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 0 {
		_, _ = w.Write([]byte("invalid id"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	book, err := h.bookH.GetBookByID(ctx, id)
	if err != nil {
		_, _ = w.Write([]byte("book does not exist"))
		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		_, _ = w.Write([]byte("can't encode"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (h Handler) PostBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("can't read!"))

		return
	}

	var book entities.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		_, _ = w.Write([]byte("can't decode "))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	book1, err := h.bookH.PostBook(ctx, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	data, err := json.Marshal(book1)
	if err != nil {
		log.Print(err)

		_, _ = w.Write([]byte("can't read!"))

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (h Handler) PutBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		_, _ = w.Write([]byte("can't read!"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var book entities.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		_, _ = w.Write([]byte("can't decode "))
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	book, err = h.bookH.PutBook(ctx, &book, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (h Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid id"))

		return
	}

	_, err = h.bookH.DeleteBook(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("can't delete"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted"))
}
