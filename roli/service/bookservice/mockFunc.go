package bookservice

import "layeredProject/entities"

type mockStore struct{}

func (m mockStore) GetAllBook(string2, string3 string) []entities.Book {
	i := 1

	if string2 == "" && string3 == "" {
		return []entities.Book{{BookID: i,
			AuthorID: 1, Title: "jk", Publication: "arihant", PublishedDate: "24/04/1990",
			Author: entities.Author{}}, {BookID: i + 1, AuthorID: 1, Title: "jk", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{}}}
	}

	if string2 == "jk" && string3 == "true" {
		i := 2

		return []entities.Book{
			{BookID: i, AuthorID: 1, Title: "jk", Publication: "penguin", PublishedDate: "25/04/2000",
				Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001",
					PenName: "roli"}}}
	}

	return []entities.Book{}
}

func (m mockStore) GetBookByID(i int) entities.Book {
	if i == 1 {
		return entities.Book{BookID: 1, AuthorID: 1, Title: "jk", Publication: "penguin",
			PublishedDate: "24/04/1990", Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu",
				DOB: "25/04/2001", PenName: "roli"}}
	}

	return entities.Book{}
}

func (m mockStore) PostBook(book *entities.Book) (int, error) {
	i := 3
	if book.BookID == i {
		return book.BookID, nil
	}

	return -1, nil
}

func (m mockStore) PutBook(book *entities.Book, id int) (int, error) {
	i := 3
	if book.BookID == i {
		return book.BookID, nil
	}

	return -1, nil
}

func (m mockStore) DeleteBook(i int) (int, error) {
	if i < 0 {
		return -1, nil
	}

	return i, nil
}
