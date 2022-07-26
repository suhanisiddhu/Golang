package bookhttp

import (
	"errors"
	"layeredProject/entities"
)

type mockService struct{}

func (m mockService) GetAllBook(s, s2 string) []entities.Book {
	if s == "" && s2 == "" {
		i := 1

		return []entities.Book{{BookID: i, AuthorID: 1, Title: "jk", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{}},
			{BookID: i + 1, AuthorID: 1, Title: "jk", Publication: "penguin", PublishedDate: "25/04/2000",
				Author: entities.Author{}}}
	}

	if s == "book two" && s2 == "true" {
		i := 2

		return []entities.Book{
			{BookID: i, AuthorID: 1, Title: "jk", Publication: "penguin", PublishedDate: "25/04/2000",
				Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}}
	}

	return []entities.Book{}
}

func (m mockService) GetBookByID(i int) (entities.Book, error) {
	if i == 1 {
		return entities.Book{BookID: 1, AuthorID: 1, Title: "jk", Publication: "penguin", PublishedDate: "25/04/2000",
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}, nil
	}

	return entities.Book{}, nil
}

func (m mockService) PostBook(book *entities.Book) (entities.Book, error) {
	i := 4
	if book.BookID == i {
		return *book, nil
	}

	return entities.Book{}, nil
}

func (m mockService) PutBook(book *entities.Book, id int) (entities.Book, error) {
	i := 4
	if book.BookID == i {
		return *book, nil
	}

	return entities.Book{}, nil
}

func (m mockService) DeleteBook(i int) (int, error) {
	if i <= 0 {
		return i, errors.New("error")
	}

	return i, nil
}
