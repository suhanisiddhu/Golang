package bookservice

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"layeredProject/datastore"
	"layeredProject/entities"

)

type BookService struct {
	bookService datastore.BookStorer
}

func New(s datastore.BookStorer) BookService {
	return BookService{s}
}

func (bs BookService) GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Book, error) {
	books, err := bs.bookService.GetAllBook(ctx, title, includeAuthor)
	if err != nil {
		return []entities.Book{}, err
	}

	return books, nil
}

func (bs BookService) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, nil
	}

	book, err := bs.bookService.GetBookByID(ctx, id)
	if err != nil {
		log.Println("error")
	}

	return book, nil
}

func (bs BookService) PostBook(ctx context.Context, book *entities.Book) (entities.Book, error) {
	correctPublication := checkPublication(book.Publication)
	if book.Title == "" || book.AuthorID < 0 || correctPublication || book.Publication == "" || book.BookID < 0 {
		return entities.Book{}, errors.New("invalid parameters")
	}

	if !isValidPublishedDate(book.PublishedDate) {
		return entities.Book{}, errors.New("error")
	}

	id, err := bs.bookService.PostBook(ctx, book)

	if err != nil || id == -1 {
		return entities.Book{}, err
	}

	book.BookID = id

	return *book, nil
}

func (bs BookService) PutBook(ctx context.Context, book *entities.Book, id int) (entities.Book, error) {
	if book.Title == "" || book.AuthorID <= 0 || checkPublication(book.Publication) || book.Publication == "" ||
		book.BookID < -1 || book.PublishedDate == "" {
		return entities.Book{}, errors.New("it's error")
	}

	if !isValidPublishedDate(book.PublishedDate) {
		return entities.Book{}, errors.New("error")
	}

	i, err := bs.bookService.PutBook(ctx, book, id)

	if err != nil || id <= 0 {
		return entities.Book{}, err
	}

	book.BookID = i

	return *book, nil
}

func (bs BookService) DeleteBook(ctx context.Context, id int) (int, error) {
	if id < 0 {
		return -1, errors.New("invalid")
	}

	i, err := bs.bookService.DeleteBook(ctx, id)
	if err != nil || i == 0 {
		return -1, errors.New("invalid")
	}

	return i, nil
}

func checkPublication(publication string) bool {
	_ = strings.ToLower(publication)

	return !(publication == "penguin" || publication == "scholastic" || publication == "arihant")
}
func isValidPublishedDate(date string) bool {
	split := strings.Split(date, "/")

	yearInstr := split[2]

	yearInint, err := strconv.Atoi(yearInstr)

	if err != nil {
		log.Printf("Cannot convert dob in integer : %v", yearInint)
	}

	if yearInint < 2022 && yearInint > 1880 {
		return true
	}

	return false
}
