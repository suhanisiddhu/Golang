package author

import (
	"database/sql"
	"errors"
	"layeredProject/entities"
	"log"
)

type Store struct {
	DB *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db}
}

// PostAuthor
func (s Store) PostAuthor(author entities.Author) (int, error) {
	var a entities.Author

	row := s.DB.QueryRow("select * from author where authorId=?", author.AuthorID)

	err := row.Scan(&a.AuthorID)
	if err != nil {
		log.Print(err)
	}

	if a.AuthorID == author.AuthorID {
		return -1, errors.New("already exists")
	}

	res, err := s.DB.Exec("insert into author(authorId,firstName,lastName,dob,penName)values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		return -1, err
	}

	id, _ := res.LastInsertId()
	author.AuthorID = int(id)

	return int(id), nil
}

// PutAuthor
func (s Store) PutAuthor(author entities.Author) (int, error) {
	var a entities.Author

	row := s.DB.QueryRow("select * from author where authorId=?", author.AuthorID)

	err := row.Scan(&a.AuthorID, &a.FirstName, &a.LastName, &a.DOB, &a.PenName)
	if err != nil {
		res, _ := s.DB.Exec("insert into author(authorId,firstName,lastName,dob,penName)values(?,?,?,?,?)",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
		id, err := res.LastInsertId()
		if err != nil {
			log.Print(err)
		}
		return int(id), nil
	}

	res, _ := s.DB.Exec("update author set authorId=?,firstName=?,lastName=?,dob=?,penName=? where authorId=?",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, a.AuthorID)
	id, _ := res.LastInsertId()
	author.AuthorID = int(id)

	return int(id), nil
}

// DeleteAuthor
func (s Store) DeleteAuthor(id int) (int, error) {
	res, _ := s.DB.Exec("delete from author where authorId=?", id)

	count, err := res.RowsAffected()
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
