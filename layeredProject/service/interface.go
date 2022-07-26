package service

import (
	"context"
	"layeredProject/entities"
)

type AuthorService interface {
	PostAuthor(context.Context, entities.Author) (entities.Author, error)
	PutAuthor(context.Context, entities.Author) (entities.Author, error)
	DeleteAuthor(context.Context, int) (int, error)
}

type BookService interface {
	GetAllBook(context.Context, string, string) ([]entities.Book, error)
	GetBookByID(context.Context, int) (entities.Book, error)
	PostBook(ctx context.Context, book *entities.Book) (entities.Book, error)
	PutBook(ctx context.Context, book *entities.Book, id int) (entities.Book, error)
	DeleteBook(context.Context, int) (int, error)
}
