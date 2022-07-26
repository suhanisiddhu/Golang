package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"gofr/datastore/author"
	"gofr/datastore/book"
	"gofr/driver"
	"gofr/http/authorhttp"
	"gofr/http/bookhttp"
	"gofr/service/authorservice"
	"gofr/service/bookservice"
)

func main() {
	//r := mux.NewRouter()

	DB := driver.Connection()
	defer DB.Close()

	authorStore := author.New(DB)
	authorService := authorservice.New(authorStore)
	authorHandler := authorhttp.New(authorService)
	k := gofr.New()
	k.POST("/author", authorHandler.PostAuthor)
	k.PUT("/author/{id}", authorHandler.PutAuthor)
	k.DELETE("/author/{id}", authorHandler.DeleteAuthor)

	//r.HandleFunc("/author", authorHandler.PostAuthor).Methods("POST")
	//r.HandleFunc("/author/{id}", authorHandler.PutAuthor).Methods("PUT")
	//r.HandleFunc("/author/{id}", authorHandler.DeleteAuthor).Methods("DELETE")

	bookStore := book.New(DB)
	bookService := bookservice.New(bookStore)
	bookHandler := bookhttp.New(bookService)
	k.GET("/book", bookHandler.GetAllBook)
	k.GET("/book/{id}", bookHandler.GetBookByID)
	k.POST("/book", bookHandler.PostBook)
	k.PUT("/book/{id}", bookHandler.PutBook)
	k.DELETE("/book/{id}", bookHandler.DeleteBook)
	k.Start()
}

//r.HandleFunc("/book", bookHandler.GetAllBook).Methods("GET")
//r.HandleFunc("/book/{id}", bookHandler.GetBookByID).Methods("GET")
//r.HandleFunc("/book", bookHandler.PostBook).Methods("POST")
//r.HandleFunc("/book/{id}", bookHandler.PutBook).Methods("PUT")
//r.HandleFunc("/book/{id}", bookHandler.DeleteBook).Methods("DELETE")

/*server := http.Server{
	Addr:    ":8000",
	Handler: r,
}

fmt.Println("server started")

err := server.ListenAndServe()

if err != nil {
	fmt.Println(err)
}*/

//log.Fatal(http.ListenAndServe(":8000", r))
