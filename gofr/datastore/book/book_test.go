package book

import (
	"context"
	"log"
	"testing"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	//"layeredProject/driver"
	"layeredProject/entities"

)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("error while establishing dummy connection for booksDB: %v", err)
	}

	return db, mock
}

func TestGetAllBook(t *testing.T) {
	db, mock := NewMock(t)
	b := New(db)
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string
		expected      []entities.Book
		err           error
	}{
		{"getting all books", "", "false", []entities.Book{{BookID: 4, AuthorID: 1, Title: "suhani", Publication: "penguin", PublishedDate: "24/04/1990", Author: entities.Author{}},
			{BookID: 5, AuthorID: 2, Title: "hi", Publication: "arihant", PublishedDate: "24/04/1990", Author: entities.Author{}}}, nil},
	}

	for _, tc := range Testcases {
		rows := sqlmock.NewRows([]string{"id", "authorID", "title", "publication", "publishdate"}).
			AddRow(4, 1, "suhani", "penguin", "24/04/1990").AddRow(5, 1, "hi", "arihant", "24/04/2001")
		mock.ExpectQuery("SELECT * FROM book").WithArgs().WillReturnRows(rows).WillReturnError(tc.err)

		_, err := b.GetAllBook(context.TODO(), tc.title, tc.includeAuthor)

		assert.Equal(t, tc.err, err, "test failed: %v", tc.desc)
	}
}

/*for _, tc := range Testcases {

	DB := driver.Connection()
	bookStore := New(DB)

	book := bookStore.GetAllBook(tc.title, tc.includeAuthor)

	if !reflect.DeepEqual(book, tc.expected) {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/

func TestGetBookByID(t *testing.T) {
	db, mock := NewMock(t)
	b := New(db)
	Testcases := []struct {
		desc     string
		targetID int
		include  string
		output   entities.Book
		err      error
	}{
		{"book fetching", 1, "false", entities.Book{BookID: 1, AuthorID: 1, Title: "jk rowling", Publication: "penguin", PublishedDate: "24/04/1990", Author: entities.Author{}}, nil},
		{"invalid id", -1, "false", entities.Book{}, errors.New("error")},
	}

	for _, tc := range Testcases {
		mock.ExpectQuery("select * from book where id=?").WithArgs(tc.targetID).
			WillReturnRows(sqlmock.NewRows([]string{"bookId", "authorId", "title", "publication", "publishedDate"}).
				AddRow(1, 1, "suhani", "arihant", "24-04-1990")).WillReturnError(tc.err)

		_, err := b.GetBookByID(context.TODO(), tc.targetID)
		assert.Equal(t, tc.err, err, "test failed: %v", tc.desc)
	}
}

/*for _, tc := range Testcases {

	DB := driver.Connection()
	bookStore := New(DB)

	book := bookStore.GetBookByID(context.TODO(),tc.targetID)

	if book != tc.expected {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/

func TestPostBook(t *testing.T) {
	db, mock := NewMock(t)
	b := New(db)
	testcases := []struct {
		desc         string
		body         *entities.Book
		lastInserted int64
		affectedRow  int64
		err          error
	}{

		{"valid book", &entities.Book{BookID: 8, AuthorID: 3, Title: "3 mistakes", Publication: "penguin", PublishedDate: "25/04/2001",
			Author: entities.Author{AuthorID: 3, FirstName: "", LastName: "", DOB: "", PenName: ""}}, 8, 1, nil},
		{"book already exists", &entities.Book{BookID: 8, AuthorID: 3, Title: "3 mistakes", Publication: "penguin", PublishedDate: "25/04/2001",
			Author: entities.Author{AuthorID: 3, FirstName: "", LastName: "", DOB: "", PenName: ""}}, 8, 1, errors.New("already exists")},
	}

	for _, tc := range testcases {
		mock.ExpectExec("insert into book(authorId,title,publication,publishedDate)values(?,?,?,?)").
			WithArgs(tc.body.Author.AuthorID, tc.body.Title, tc.body.Publication, tc.body.PublishedDate).
			WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).WillReturnError(tc.err)

		_, err := b.PostBook(context.TODO(), tc.body)
		assert.Equal(t, tc.err, err, "Test failed: %v", tc.desc)
	}
}

/*for _, tc := range testcases {

	DB := driver.Connection()
	bookStore := New(DB)
	id, err := bookStore.PostBook(&tc.body)

	if id == 0 && tc.err != err {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc         string
		body         *entities.Book
		lastInserted int64
		affectedRow  int64
		err          error
	}{

		{"creating a book", &entities.Book{BookID: 4, AuthorID: 1, Title: "JK", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{}}, 4, 1, errors.New("error")},
		{"updating book", &entities.Book{BookID: 4, AuthorID: 1, Title: "JK", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}, 0, 1, errors.New("error")},
	}

	for _, tc := range testcases {
		db, mock := NewMock(t)
		b := New(db)

		mock.ExpectQuery("SELECT * FROM book WHERE bookId=?").WithArgs(tc.body.BookID).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1)).WillReturnError(tc.err)
		//mock.ExpectQuery("select * from book where id=?").WithArgs(tc.body.AuthorID).WillReturnRows(sqlmock.NewRows([]string{"count"}).WillReturnError(tc.err)
		mock.ExpectExec("UPDATE book SET bookId=?, authorId=?, title=?, publication=?, publishedDate=? WHERE bookId=?").
			WithArgs(tc.body.BookID, tc.body.Author.AuthorID, tc.body.Title, tc.body.Publication, tc.body.PublishedDate, tc.body.BookID).
			WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).WillReturnError(tc.err)

		mock.ExpectExec("INSERT INTO book(bookId, authorId, title, publication, publishedDate)VALUES (?,?,?,?,?)").
			WithArgs(tc.body.BookID, tc.body.Author.AuthorID, tc.body.Title, tc.body.Publication, tc.body.Publication).
			WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).WillReturnError(tc.err)

		_, err := b.PutBook(context.TODO(), tc.body, tc.body.BookID)

		assert.Equal(t, tc.err, err, "Test failed : %v", tc.desc)
	}
}

/*DB := driver.Connection()
	bookStore := New(DB)

	_, err := bookStore.PutBook(&tc.body, tc.targetID)

	if tc.expectedErr != err {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/

func TestDeleteBook(t *testing.T) {
	db, mock := NewMock(t)
	b := New(db)
	testcases := []struct {
		desc         string
		targetID     int
		lastInsertID int64
		rowsAffected int64
		err          error
	}{
		{"valid bookId", 2, 0, 1, nil},
		{"invalid bookId", -1, 0, 0, errors.New("invalid ID")},
	}

	for _, tc := range testcases {
		mock.ExpectExec("DELETE FROM book WHERE bookId=?").WithArgs(tc.targetID).
			WillReturnResult(sqlmock.NewResult(tc.lastInsertID, tc.rowsAffected)).WillReturnError(tc.err)

		count, _ := b.DeleteBook(context.TODO(), tc.targetID)
		assert.Equal(t, int(tc.rowsAffected), count)
	}
}

/*for _, tc := range testcases {
	DB := driver.Connection()
	bookStore := New(DB)
	if id == 0 && tc.err != err {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/
