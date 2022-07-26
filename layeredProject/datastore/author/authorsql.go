package author

import (
	"context"
	"errors"
	"log"
	"database/sql"
	"layeredProject/entities"

)

type Store struct {
	DB *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db}
}

// PostAuthor
func (s Store) PostAuthor(ctx context.Context, author entities.Author) (int, error) {
	var a entities.Author

	row := s.DB.QueryRowContext(ctx, "select * from author where authorId=?", author.AuthorID)

	err := row.Scan(&a.AuthorID)
	if err != nil {
		log.Print(err)
	}

	if author.AuthorID != 0 && a.AuthorID == author.AuthorID {
		return -1, errors.New("already exists")
	}

	res, err := s.DB.ExecContext(ctx, "insert into author(authorId,firstName,lastName,dob,penName)values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		return -1, err
	}

	id, _ := res.LastInsertId()
	author.AuthorID = int(id)

	return int(id), nil
}

// PutAuthor

func (s Store) PutAuthor(ctx context.Context, author entities.Author) (int, error) {
	var a entities.Author

	row := s.DB.QueryRowContext(ctx, "select * from author where authorId=?", author.AuthorID)

	err := row.Scan(&a.AuthorID, &a.FirstName, &a.LastName, &a.DOB, &a.PenName)
	if err != nil {
		res, _ := s.DB.ExecContext(ctx, "insert into author(authorId,firstName,lastName,dob,penName)VALUES (?,?,?,?,?)",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)

		if err != nil {
			log.Fatal("error")
		}

		id, err := res.LastInsertId()

		if err != nil {
			return 0, errors.New("error")
		}

		return int(id), nil
	}

	res, _ := s.DB.ExecContext(ctx, "update author set authorId=?,firstName=?,lastName=?,dob=?,penName=? where authorId=?",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, a.AuthorID)
	id, _ := res.LastInsertId()
	author.AuthorID = int(id)

	return int(id), nil
}

// DeleteAuthor
func (s Store) DeleteAuthor(ctx context.Context, id int) (int, error) {
	res, err := s.DB.ExecContext(ctx, "delete from author where authorId=?", id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
