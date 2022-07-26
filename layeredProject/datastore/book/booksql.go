package book

import (
	"context"
	"errors"
	"log"
	"database/sql"
	"layeredProject/entities"

)

type BookStore struct {
	DB *sql.DB
}

func New(db *sql.DB) BookStore {
	return BookStore{db}
}

func (bs BookStore) GetAllBook(ctx context.Context, title string, includeAuthor string) ([]entities.Book, error) {
	var books []entities.Book

	books = FetchingAllBooks(ctx, title, bs.DB)
	if includeAuthor == "true" {
		books = BooksWithAuthor(books, bs.DB)
	}

	return books, nil
}

func (bs BookStore) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	var b entities.Book

	row := bs.DB.QueryRowContext(ctx, "select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err != nil {
		log.Print(err)
		return entities.Book{}, errors.New("error")
	}

	return b, nil
}

func (bs BookStore) PostBook(ctx context.Context, book *entities.Book) (int, error) {
	result, err := bs.DB.ExecContext(ctx, "insert into book(authorId,title,publication,publishedDate)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return int(id), nil
}

func (bs BookStore) PutBook(ctx context.Context, book *entities.Book, id int) (int, error) {
	var b entities.Book

	row := bs.DB.QueryRowContext(ctx, "select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err == nil {
		_, _ = bs.DB.ExecContext(ctx, "update book set id=?,authorId=?,title=?,publication=?,publishedDate=? where id=?",
			book.BookID, book.AuthorID, book.Title, book.Publication, book.PublishedDate, id)
		return book.BookID, nil
	}

	result, err := bs.DB.ExecContext(ctx, "insert into book(authorId,title,publication,publishedDate)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		return -1, errors.New("error")
	}

	i, err := result.LastInsertId()

	if err != nil {
		return -1, errors.New("error")
	}

	return int(i), nil
}
func (bs BookStore) DeleteBook(ctx context.Context, id int) (int, error) {
	result, err := bs.DB.ExecContext(ctx, "delete from book where id=?", id)
	if err != nil {
		return -1, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return int(i), err
	}

	return int(i), nil
}

func FetchingAllBooks(ctx context.Context, title string, db *sql.DB) []entities.Book {

	var (
		rows *sql.Rows
		err  error
	)

	if title == "" {
		rows, err = db.QueryContext(ctx, "SELECT * FROM book")
		if err != nil {
			log.Print(err)
		}
	} else {
		rows, err = db.QueryContext(ctx, "SELECT * FROM book where title=?", title)
		if err != nil {
			log.Print(err)
		}
	}

	var bk []entities.Book

	for rows.Next() {
		var b entities.Book

		err = rows.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
		if err != nil {
			log.Print(err)
		}

		bk = append(bk, b)
	}

	return bk
}

func FetchingAuthor(id int, db *sql.DB) (int, entities.Author) {

	row := db.QueryRow("SELECT * FROM author where authorId=?", id)

	defer db.Close()

	var author entities.Author

	if err := row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		log.Printf("failed: %v\n", err)
		return 0, entities.Author{}
	}

	return author.AuthorID, author
}

/*func BookWithAuthor(books entities.Book, db *sql.DB) entities.Book {
	_, a := FetchingAuthor(books.AuthorID, db)
	books.Author = a

	return books
}*/

func BooksWithAuthor(books []entities.Book, db *sql.DB) []entities.Book {
	for i := range books {
		_, a := FetchingAuthor(books[i].AuthorID, db)
		books[i].Author = a
	}

	return books
}
