package bookservice

import (
	"errors"
	"layeredProject/datastore"
	"layeredProject/entities"
	"strings"
)

type BookService struct {
	bookService datastore.BookStorer
}

func New(s datastore.BookStorer) BookService {
	return BookService{s}
}

func (bs BookService) GetAllBook(title, includeAuthor string) []entities.Book {
	books := bs.bookService.GetAllBook(title, includeAuthor)
	return books
}

func (bs BookService) GetBookByID(id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, nil
	}

	book := bs.bookService.GetBookByID(id)

	return book, nil
}

func (bs BookService) PostBook(book *entities.Book) (entities.Book, error) {
	if book.Title == "" || book.AuthorID < 0 || checkPublication(book.Publication) {
		return entities.Book{}, nil
	}

	id, err := bs.bookService.PostBook(book)
	if err != nil || id == -1 {
		return entities.Book{}, err
	}

	book.BookID = id

	return *book, nil
}

func (bs BookService) PutBook(book *entities.Book, id int) (entities.Book, error) {
	if book.Title == "" || book.AuthorID <= 0 || checkPublication(book.Publication) || book.Publication == "" {
		return entities.Book{}, errors.New("it's error")
	}

	i, err := bs.bookService.PutBook(book, id)
	if err != nil || id <= -1 {
		return entities.Book{}, err
	}

	book.BookID = i

	return *book, nil
}

func (bs BookService) DeleteBook(id int) (int, error) {
	if id < 0 {
		return -1, nil
	}

	i, err := bs.bookService.DeleteBook(id)
	if err != nil || i == -1 {
		return -1, err
	}

	return i, nil
}

func checkPublication(publication string) bool {
	_ = strings.ToLower(publication)

	return !(publication == "penguin" || publication == "scholastic" || publication == "arihant")
}
