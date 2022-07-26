package author

import (
	"context"
	"log"
	"testing"
	"errors"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"layeredProject/entities"

)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestPostAuthor(t *testing.T) {
	db, mock := NewMock(t)
	a := New(db)
	testcases := []struct {
		desc         string
		body         entities.Author
		lastInserted int64
		affectedRow  int64
		err          error
	}{
		{"valid author", entities.Author{AuthorID: 4, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}, 4, 1, nil},
		{"existing author", entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}, 0, 0, errors.New("author already exists")},
	}

	for _, tc := range testcases {
		query := "insert into author(authorId,firstName,lastName,dob,penName)values(?,?,?,?,?)"
		mock.ExpectExec(query).WithArgs(tc.body.AuthorID, tc.body.FirstName, tc.body.LastName, tc.body.DOB, tc.body.PenName).
			WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).WillReturnError(tc.err)

		_, err := a.PostAuthor(context.TODO(), tc.body)
		assert.Equal(t, tc.err, err, "test failed: %v", tc.desc)
	}
}

/*Db := driver.Connection()
authorStore := New(Db)

id, _ := authorStore.PostAuthor(tc.body)

if id != tc.expectedID {
	t.Errorf("failed for %v\n", tc.desc)
}*/

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc         string
		body         entities.Author
		lastInserted int64
		affectedRow  int64
		err          error
	}{
		{"valid author", entities.Author{AuthorID: 4, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}, 2, 0, nil},
		{"existing author", entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}, 1, 0, errors.New("error")},
	}
	for _, tc := range testcases {
		db, mock := NewMock(t)
		a := New(db)

		mock.ExpectQuery("SELECT * FROM author WHERE authorId=?").
			WithArgs(tc.body.AuthorID).WillReturnRows(sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"})).
			WillReturnError(tc.err)
		mock.ExpectExec("INSERT INTO author(authorId, firstName, lastName, DOB, penName) VALUES(?,?,?,?,?)").
			WithArgs(tc.body.AuthorID, tc.body.FirstName, tc.body.LastName, tc.body.DOB, tc.body.PenName).
			WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).WillReturnError(tc.err)
		mock.ExpectExec("UPDATE author SET authorId=?, firstName=?, lastName=?, dob=?, penName=? WHERE authorId=?").
			WithArgs(tc.body.AuthorID, tc.body.FirstName, tc.body.LastName, tc.body.DOB, tc.body.PenName, tc.body.AuthorID).
			WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).WillReturnError(tc.err)

		_, err := a.PutAuthor(context.TODO(), tc.body)
		assert.Equal(t, tc.err, err, "Test failed: %v", tc.desc)
	}
}

/*for _, tc := range testcases {
	Db := driver.Connection()
	authorStore := New(Db)

	id, _ := authorStore.PostAuthor(tc.body)

	if id != tc.expected {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/

func TestDeleteAuthor(t *testing.T) {
	db, mock := NewMock(t)
	a := New(db)
	testcases := []struct {
		desc         string
		target       int
		lastInserted int64
		affectedRow  int64
		err          error
	}{
		{"Author deleted.", 3, 3, 1, nil},
		{"Author not present.", 33, 3, 0, errors.New("invalid id")},
	}

	for _, tc := range testcases {
		mock.ExpectExec("delete from author where authorId=?").
			WithArgs(tc.target).WillReturnResult(sqlmock.NewResult(tc.lastInserted, tc.affectedRow)).
			WillReturnError(tc.err)

		_, err := a.DeleteAuthor(context.TODO(), tc.target)
		assert.Equal(t, tc.err, err, "Test failed: %v", tc.desc)
	}
}

/*for _, tc := range testcases {
	Db := driver.Connection()
	authorStore := New(Db)

	count, _ := authorStore.DeleteAuthor(tc.target)

	if count != tc.expected {
		t.Errorf("failed for %v\n", tc.desc)
	}
}*/
