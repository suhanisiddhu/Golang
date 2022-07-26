package book

import (
	"database/sql"
	"errors"
	"layeredProject/entities"
	"log"
)

type BookStore struct {
	DB *sql.DB
}

func New(DB *sql.DB) BookStore {
	return BookStore{DB}
}

func (bs BookStore) GetAllBook(title string, includeAuthor string) []entities.Book {
	var books []entities.Book

	if title != "" {
		books = FetchingAllBooks(title, bs.DB)
		if includeAuthor == "true" {

			books = BooksWithAuthor(books, bs.DB)
		}
	} else {
		books = FetchingAllBooks("", bs.DB)
		if includeAuthor == "true" {

			books = BooksWithAuthor(books, bs.DB)
		}
	}
	return books
}

func (bs BookStore) GetBookByID(id int) entities.Book {
	var b entities.Book

	row := bs.DB.QueryRow("select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err != nil {
		log.Print(err)
		return entities.Book{}
	}

	return b
}

func (bs BookStore) PostBook(book *entities.Book) (int, error) {
	result, err := bs.DB.Exec("insert into book(authorId,title,publication,publishedDate)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		log.Print(err)
		return -1, errors.New("already existing")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return int(id), nil
}

func (bs BookStore) PutBook(book *entities.Book, id int) (int, error) {
	var b entities.Book

	row := bs.DB.QueryRow("select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err == nil {
		// updating
		_, _ = bs.DB.Exec("update book set id=?,authorId=?,title=?,publication=?,publishedDate=? where id=?",
			book.BookID, book.AuthorID, book.Title, book.Publication, book.PublishedDate, id)
		return book.BookID, nil
	}

	result, err := bs.DB.Exec("insert into book(authorId,title,publication,publishedDate)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		return -1, err
	}

	i, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(i), nil
}
func (bs BookStore) DeleteBook(id int) (int, error) {
	result, err := bs.DB.Exec("delete from book where id=?", id)
	if err != nil {
		return -1, nil
	}
	i, err := result.RowsAffected()
	if err != nil {
		return -1, nil
	}
	return int(i), nil
}

func FetchingAllBooks(title string, Db *sql.DB) []entities.Book {

	var (
		Rows *sql.Rows
		err  error
	)

	if title == "" {
		Rows, err = Db.Query("SELECT * FROM book")
		if err != nil {
			log.Print(err)
		}
	} else {
		Rows, err = Db.Query("SELECT * FROM book where title=?", title)
		if err != nil {
			log.Print(err)
		}
	}

	var bk []entities.Book

	for Rows.Next() {
		var b entities.Book

		err = Rows.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
		if err != nil {
			log.Print(err)
		}

		bk = append(bk, b)
	}

	return bk
}

func FetchingAuthor(id int, DB *sql.DB) (int, entities.Author) {

	row := DB.QueryRow("SELECT * FROM author where authorId=?", id)

	defer DB.Close()

	var author entities.Author

	if err := row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		log.Printf("failed: %v\n", err)
		return 0, entities.Author{}
	}

	return author.AuthorID, author
}

func BooksWithAuthor(books []entities.Book, db *sql.DB) []entities.Book {
	for i := range books {
		_, a := FetchingAuthor(books[i].AuthorID, db)
		books[i].Author = a
	}

	return books
}
