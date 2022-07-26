package datastore

import (
	"context"
	"gofr/entities"
)

type AuthorStorer interface {
	PostAuthor(context.Context, entities.Author) (int, error)
	PutAuthor(context.Context, entities.Author) (int, error)
	DeleteAuthor(context.Context, int) (int, error)
}

type BookStorer interface {
	GetAllBook(ctx context.Context, string2 string, string3 string) ([]entities.Book, error)
	GetBookByID(context.Context, int) (entities.Book, error)
	PostBook(ctx context.Context, book *entities.Book) (int, error)
	PutBook(ctx context.Context, book *entities.Book, id int) (int, error)
	DeleteBook(context.Context, int) (int, error)
}
