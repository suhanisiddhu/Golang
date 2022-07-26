package book

import (
	"errors"
	"layeredProject/driver"
	"layeredProject/entities"
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
		{"getting all books", "", "", []entities.Book{{1,
			1, "jk rowling", "penguin", "25/04/1990", entities.Author{}},
			{1, 1, "jk rowling", "penguin", "25/04/1990", entities.Author{}}},
		},
		{"getting book with author and particular title", "jk", "true", []entities.Book{
			{2, 1, "jk", "penguin", "25/04/1990", entities.Author{1, "suhani",
				"siddhu", "25/04/2001", "roli"}}},
		},
	}

	for _, tc := range Testcases {

		DB := driver.Connection()
		bookStore := New(DB)

		book := bookStore.GetAllBook(tc.title, tc.includeAuthor)

		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID int

		expected entities.Book
	}{
		{"getting book by id",
			1, entities.Book{1, 1, "jk rowling", "penguin",
				"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}},

		{"invalid id", -1, entities.Book{}},
	}

	for _, tc := range Testcases {

		DB := driver.Connection()
		bookStore := New(DB)

		book := bookStore.GetBookByID(tc.targetID)

		if book != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		err error
	}{
		{"valid case", entities.Book{4, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, nil},

		{"book already present ", entities.Book{4, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, errors.New("already existing")},
	}
	for _, tc := range testcases {

		DB := driver.Connection()
		bookStore := New(DB)

		id, err := bookStore.PostBook(&tc.body)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}

}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc     string
		body     entities.Book
		targetID int

		expectedErr error
	}{
		{"creating  book", entities.Book{4, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, 4, nil},

		{"updating book", entities.Book{4, 1, "deciding decade", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, 4, nil},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		_, err := bookStore.PutBook(&tc.body, tc.targetID)

		if tc.expectedErr != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID int

		err error
	}{
		{"valid id", 2, nil},
		{"invalid id", -1, errors.New("invalid id")},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		id, err := bookStore.DeleteBook(tc.targetID)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}

}
