package bookservice

import (
	"errors"
	"layeredProject/entities"
	"log"
	"reflect"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expected []entities.Book
	}{
		{"all books details", "", "", []entities.Book{{1,
			1, "jk rowling", "penguin", "25/04/1990", entities.Author{}},
			{1, 1, "jk rowling", "penguin", "25/04/1990", entities.Author{}}},
		},
		{"getting book with author and given title", "jk", "true", []entities.Book{
			{2, 1, "jk", "penguin", "25/04/1990", entities.Author{1, "suhani",
				"siddhu", "25/04/2001", "roli"}}},
		},
	}

	for _, tc := range Testcases {
		b := New(mockStore{})
		book := b.GetAllBook(tc.title, tc.includeAuthor)

		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID int

		expectedBody entities.Book
		expectedErr  error
	}{
		{"getting book by id",
			1, entities.Book{1, 1, "jk rowling", "penguin",
				"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, nil},

		{"invalid id", -1, entities.Book{}, errors.New("invalid id")},
	}

	for _, tc := range Testcases {
		b := New(mockStore{})
		book, err := b.GetBookByID(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBody) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{desc: "already exists", body: entities.Book{BookID: 1, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "24/04/1990", Author: entities.Author{AuthorID: 1, FirstName: "suhani",
				LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}},
		{desc: "invalid bookID", body: entities.Book{BookID: -1, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "24/04/1990", Author: entities.Author{AuthorID: 1, FirstName: "suhani",
				LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})

		book, err := b.PostBook(&tc.body)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBook) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{desc: "inserting  book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "jk", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu",
				DOB: "25/04/2001", PenName: "roli"}}, expectedBook: entities.Book{BookID: 4, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1,
				FirstName: "suhani", LastName: "suhani", DOB: "25/04/2001", PenName: "roli"}}},

		{desc: "updating book", body: entities.Book{BookID: 3, AuthorID: 1, Title: "jk rowling", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu",
				DOB: "25/04/2001", PenName: "roli"}}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})

		book, err := b.PostBook(&tc.body)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBook) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID int

		expected int
	}{
		{"valid id", 1, 1},
		{"invalid id", -1, -1},
	}

	for _, tc := range testcases {
		b := New(mockStore{})

		id, err := b.DeleteBook(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(id, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}

	}

}
