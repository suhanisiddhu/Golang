package bookhttp

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"gofr/entities"
	"gofr/service"
	"net/http"
	"strconv"
)

type Handler struct {
	bookH service.BookService
}

func New(bookS service.BookService) Handler {
	return Handler{bookS}
}

func (h Handler) GetAllBook(ctx *gofr.Context) (interface{}, error) {
	title := ctx.Request().URL.Query().Get("title")
	includeAuthor := ctx.Request().URL.Query().Get("includeAuthor")

	books, err := h.bookH.GetAllBook(ctx, title, includeAuthor)

	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}
	}
	return books, nil
}

func (h Handler) GetBookByID(ctx *gofr.Context) (interface{}, error) {
	inputID := ctx.PathParam("id")
	id, err := strconv.Atoi(inputID)

	if err != nil || id < 0 {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Reason:     "invalid id",
		}

	}

	book, err := h.bookH.GetBookByID(ctx, id)
	if err != nil {
		return nil, &errors.Response{
			Reason: "book does not exists",
		}

	}

	return book, nil
}

func (h Handler) PostBook(ctx *gofr.Context) (interface{}, error) {

	var book entities.Book
	if err := ctx.Bind(&book); err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}
	}

	book1, err := h.bookH.PostBook(ctx, &book)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}

	}
	return book1, nil
}

func (h Handler) PutBook(ctx *gofr.Context) (interface{}, error) {

	var book entities.Book
	if err := ctx.Bind(&book); err != nil {
		return nil, err
	}
	inputID := ctx.PathParam("id")
	id, err := strconv.Atoi(inputID)

	if err != nil {
		return nil, err
	}

	book, err = h.bookH.PutBook(ctx, &book, id)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}
	}

	return book, nil
}

func (h Handler) DeleteBook(ctx *gofr.Context) (interface{}, error) {

	inputID := ctx.PathParam("id")
	id, err := strconv.Atoi(inputID)

	if err != nil || id < 0 {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Reason:     "invalid id",
		}

	}

	_, err = h.bookH.DeleteBook(ctx, id)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Reason:     "cannot delete",
		}

	}
	return "successfully deleted! ", nil

}
