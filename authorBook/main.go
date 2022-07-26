package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	BookID          int    `json:"BookID"`
	Title           string `json:"title"`
	Author          Author `json:"author"`
	Publication     string `json:"publication"`
	PublicationDate string `json:"published_date"`
}

type Author struct {
	AuthorId  int    `json:"BookID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	PenName   string `json:"pen_name"`
}

func Connection() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:Suhani@123@tcp(127.0.0.1:3306)/public")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getBook(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	defer db.Close()
	x := request.URL.Query()
	title := x.Get("title")
	includeAuthor := x.Get("includeAuthor")
	var rows *sql.Rows
	var err error
	if title == "" {
		rows, err = db.Query("select * from book;")
	} else {
		rows, err = db.Query("select * from book where title=?;", title)
	}
	if err != nil {
		log.Print(err)
	}
	books := []Book{}
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Title, &book.BookID, &book.Author.AuthorId, &book.Publication, &book.PublicationDate)
		if err != nil {
			log.Print(err)
		}
		if includeAuthor == "true" {
			row := db.QueryRow("select * from author where authorId=?", book.Author.AuthorId)
			row.Scan(&book.Author.AuthorId, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
		}
		books = append(books, book)
	}
	json.NewEncoder(response).Encode(books)

}

func getBookById(response http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(mux.Vars(request)["BookID"])

	if err != nil {
		log.Print(err)
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	db := Connection()
	defer db.Close()
	bookinfo := db.QueryRow("select * from book where bookId=?;", id)
	var book Book
	err = bookinfo.Scan(&book.Title, &book.BookID, &book.Author.AuthorId, &book.Publication, &book.PublicationDate)
	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			response.WriteHeader(404)
			json.NewEncoder(response).Encode(book)
			return
		}
	}
	authorrow := db.QueryRow("select * from author where authorId=?;", book.Author.AuthorId)
	err = authorrow.Scan(&book.Author.AuthorId, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
	if err != nil {
		log.Print(err)
	}
	json.NewEncoder(response).Encode(book)

}

func postBook(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	defer db.Close()
	decoder := json.NewDecoder(request.Body)
	b := Book{}
	err := decoder.Decode(&b)
	if b.Title == "" {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	var BookId int
	err = db.QueryRow("select bookId from book where title=? and authorId=?;", b.Title, b.Author.AuthorId).Scan(&BookId)
	if err == nil {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}

	authorRow := db.QueryRow("select authorId from author where authorId=?;", b.Author.AuthorId)
	var authorId int
	err = authorRow.Scan(&authorId)
	if err != nil {
		log.Print(err)
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	if !(b.Publication == "Scholastic" || b.Publication == "pengiun" || b.Publication == "arihant") {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Book{})
		return
	}
	pubYear, err := strconv.Atoi(strings.Split(b.PublicationDate, "/")[2])
	if err != nil {
		log.Print("invalid date")
		json.NewEncoder(response).Encode(Book{})
		return
	}
	if !(pubYear >= 1880 && pubYear <= time.Now().Year()) {
		log.Print("invalid date")
		json.NewEncoder(response).Encode(Book{})
		return
	}
	res, err := db.Exec("INSERT INTO book (title,authorId, publication, publishdate)\nVALUES (?,?,?,?);", b.Title, b.Author.AuthorId, b.Publication, b.PublicationDate)
	id, _ := res.LastInsertId()
	if err != nil {
		log.Print(err)
		json.NewEncoder(response).Encode(Book{})
	} else {
		b.BookID = int(id)
		json.NewEncoder(response).Encode(b)
	}
}

func postAuthor(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	defer db.Close()
	decoder := json.NewDecoder(request.Body)
	a := Author{}
	err := decoder.Decode(&a)
	fmt.Println(a)
	if a.FirstName == "" || a.Dob == "" {
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Author{})
		return
	}
	authoridExits := 0
	err = db.QueryRow("SELECT authorId from author where firstName=? and lastName=? and dob=? and penName=?", a.FirstName, a.LastName, a.Dob, a.PenName).Scan(&authoridExits)
	if err == nil {
		log.Print("author already exists")
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Author{})
		return
	}
	res, err := db.Exec("INSERT INTO author (authorId,firstName, lastName, dob, penName)\nVALUES (?,?,?,?,?);", a.AuthorId, a.FirstName, a.LastName, a.Dob, a.PenName)
	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		response.WriteHeader(400)
		json.NewEncoder(response).Encode(Author{})
	} else {
		a.AuthorId = int(id)
		json.NewEncoder(response).Encode(a)
	}
}

func putAuthor(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	var author Author
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Print(err)
		return
	}
	if author.FirstName == "" || author.LastName == "" || author.PenName == "" || author.Dob == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(request)
	ID, err := strconv.Atoi(params["BookID"])
	if ID <= 0 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := db.Query("SELECT authorId FROM author WHERE authorId = ?", ID)
	if err != nil {
		log.Print(err)
	}
	if !res.Next() {
		log.Print("BookID not present")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var id int
	err = res.Scan(&id)
	if err != nil {
		log.Print(err)
		return
	}

	_, err = db.Exec("UPDATE author SET firstName = ? ,lastName = ? ,dob = ? ,penName = ?  WHERE authorId =?", author.FirstName, author.LastName, author.Dob, author.PenName, ID)
	if err != nil {
		log.Print(err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func putBook(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	var book Book
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		return
	}
	if book.Title == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	if !(book.Publication == "penguin" || book.Publication == "Scholastic" || book.Publication == "arihant") {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	publicationDate := strings.Split(book.PublicationDate, "/")
	if len(publicationDate) < 3 {
		return
	}
	yr, _ := strconv.Atoi(publicationDate[2])
	if yr > time.Now().Year() || yr < 1880 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(request)
	ID, err := strconv.Atoi(params["BookID"])
	if ID <= 0 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Query("SELECT authorId FROM author WHERE authorId = ?", book.Author.AuthorId)
	if err != nil {
		log.Print(err)
	}

	if !result.Next() {
		log.Print("author not present", book.Author.AuthorId)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err = db.Query("SELECT * FROM book WHERE bookId = ?", book.BookID)
	if err != nil {
		log.Print(err)
	}
	if !result.Next() {
		log.Print("Book not present")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err = db.Query("UPDATE book SET title = ? ,publication = ? ,published_date = ?,authorId=?  WHERE bookId =?", book.Title, book.Publication, book.PublicationDate, book.Author.AuthorId, ID)
	if err != nil {
		log.Print(err)
	}
}

func deleteAuthor(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	defer db.Close()
	id, err := strconv.Atoi(mux.Vars(request)["BookID"])

	fmt.Println(id)
	_, err = db.Exec("delete from book where authorId=?;", id)
	if err != nil {
		log.Print(err)
		response.WriteHeader(400)
		return
	}
	_, err = db.Exec("delete from author where authorId=?;", id)
	if err != nil {
		response.WriteHeader(400)
		return
	}
	response.WriteHeader(200)
}

func deleteBook(response http.ResponseWriter, request *http.Request) {
	db := Connection()
	defer db.Close()
	id, err := strconv.Atoi(mux.Vars(request)["BookID"])

	fmt.Println(id)
	bookId := 0
	err = db.QueryRow("select bookId from book where bookId=?;", id).Scan(&bookId)
	if err == nil {
		_, err = db.Exec("delete from book where bookId=?;", id)
		if err != nil {
			response.WriteHeader(400)
			return
		}
	} else {
		response.WriteHeader(400)
		return
	}

	response.WriteHeader(200)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/book", getBook).Methods(http.MethodGet)

	r.HandleFunc("/book/{BookID}", getBookById).Methods(http.MethodGet)

	r.HandleFunc("/book", postBook).Methods(http.MethodPost)

	r.HandleFunc("/author", postAuthor).Methods(http.MethodPost)

	r.HandleFunc("/book/{BookID}", putBook).Methods(http.MethodPut)

	r.HandleFunc("/author/{BookID}", putAuthor).Methods(http.MethodPut)

	r.HandleFunc("/book/{BookID}", deleteBook).Methods(http.MethodDelete)

	r.HandleFunc("/author/{BookID}", deleteAuthor).Methods(http.MethodDelete)

	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	fmt.Println("Server started at 8000")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
