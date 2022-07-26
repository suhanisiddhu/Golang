package authorhttp

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	//"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	//"io"
	"strconv"

	//"encoding/json"
	"gofr/entities"
	"gofr/service"
	"log"
	"net/http"
)

type Handler struct {
	authorService service.AuthorService
}

func New(a service.AuthorService) Handler {
	return Handler{a}
}

func (h Handler) PostAuthor(ctx *gofr.Context) (interface{}, error) {

	var author entities.Author

	if err := ctx.Bind(&author); err != nil {
		log.Print(err)

		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Reason:     "not readable",
		}
	}
	_, err := h.authorService.PostAuthor(ctx, author)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}

	}

	//c := ctx.Request().Context()
	return "posted", nil
}

func (h Handler) PutAuthor(ctx *gofr.Context) (interface{}, error) {

	var author entities.Author
	if err := ctx.Bind(&author); err != nil {
		log.Print(err)
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}
	}
	_, err := h.authorService.PutAuthor(ctx, author)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
		}

	}
	return "successfully posted!!", nil

}

func (h Handler) DeleteAuthor(ctx *gofr.Context) (interface{}, error) {
	inputID := ctx.PathParam("id")
	i, err := strconv.Atoi(inputID)

	if err != nil || i < 0 {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Reason:     "invalid id",
		}

	}

	_, err = h.authorService.DeleteAuthor(ctx, i)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Reason:     "can't delete",
		}

	}
	return "successfully deleted", &errors.Response{
		StatusCode: http.StatusNoContent,
	}
}
