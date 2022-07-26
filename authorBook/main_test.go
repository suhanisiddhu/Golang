package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetBookAllDetails(t *testing.T) {
	testcases := []struct {
		desc      string
		req       string
		expRes    []Book
		expStatus int
	}{
		{"get all books", "/book", []Book{
			{1, "jk", Author{AuthorId: 1}, "arihant", "2009"},
			{2, "jk ROWLING", Author{AuthorId: 2}, "penguin", "2000"},
		}, http.StatusOK},
		{"get all books with query param", "/book?title=jk", []Book{
			{1, "jk", Author{AuthorId: 1}, "arihant", "2009"}}, http.StatusOK},
	}
	for j, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "localhost:8000"+tc.req, nil)

		getBook(w, req)

		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", j, tc.desc)
		}

		res, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("expected error got %v", err)
		}

		resBooks := []Book{}

		err = json.Unmarshal(res, &resBooks)
		if err != nil {
			t.Errorf("expected error  got %v", err)
		}

		if !reflect.DeepEqual(resBooks, tc.expRes) {
			t.Errorf("%v test failed %v", j, tc.desc)
		}
	}
}

func TestGetBooksById(t *testing.T) {
	testcases := []struct {
		desc      string
		req       string
		expRes    Book
		expStatus int
	}{
		{"show book", "1", Book{1, "jk", Author{1, "abc", "singh", "25/04/1997", "yee"}, "arihant", "2009"}, http.StatusOK},
		{"BookID doesn't exist", "1000", Book{}, http.StatusNotFound},
		{"Invalid BookID", "asd", Book{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/book/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"BookID": tc.req})
		getBookById(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}

		res, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("expected error  got %v", err)
		}

		resBook := Book{}

		err = json.Unmarshal(res, &resBook)
		if err != nil {
			t.Errorf("expected error  got %v", err)
		}

		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPostBookInformation(t *testing.T) {
	testcases := []struct {
		desc      string
		reqBody   Book
		expRes    Book
		expStatus int
	}{
		{"valid details", Book{2, "abcdef", Author{5, "", "", "", ""}, "Scholastic", "2/4/2018"}, Book{2, "abcdef", Author{5, "", "", "", ""}, "Scholastic", "2/4/2018"}, http.StatusOK},
		//{"Publication should be Scholastic/penguin/arihant", Book{Title: "hey", Author: Author{AuthorId: 2}, Publication: "epical", PublicationDate: "2000"}, Book{}, http.StatusBadRequest},
		//{"Published date should be between 1880 and 2022", Book{Title: "", Author: Author{AuthorId: 2}, Publication: "", PublicationDate: "1873"}, Book{}, http.StatusBadRequest},
		//{"Author should exist", Book{Title: "hey", Author: Author{AuthorId: 263}, Publication: "Penguin", PublicationDate: "2004"}, Book{}, http.StatusBadRequest},
		//{"Title can't be empty", Book{Title: "", Author: Author{AuthorId: 2}, Publication: "", PublicationDate: ""}, Book{}, http.StatusBadRequest},
		//{"Book already exists", Book{Title: "jk", Author: Author{AuthorId: 1}, Publication: "arihant", PublicationDate: "2009"}, Book{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000/book/", bytes.NewReader(body))
		postBook(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		res, _ := io.ReadAll(w.Result().Body)
		resBook := Book{}
		json.Unmarshal(res, &resBook)
		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPostAuthorInformation(t *testing.T) {
	testcases := []struct {
		desc      string
		reqBody   Author
		expRes    Author
		expStatus int
	}{
		//{"Valid details", Author{FirstName: "roli", LastName: "siddhu", Dob: "23/03/1989", PenName: "abhi"}, Author{5, "abhi", "siddhhu", "23/03/1989", "abhi"}, http.StatusOK},
		{"InValid details", Author{FirstName: "", LastName: "Sirohi", Dob: "20/12/1990", PenName: "Sh"}, Author{}, http.StatusBadRequest},
		{"Author already exists", Author{FirstName: "roli", LastName: "siddhu", Dob: "23/03/1989", PenName: "abhi"}, Author{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000/author", bytes.NewReader(body))
		postAuthor(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		res, _ := io.ReadAll(w.Result().Body)
		resAuthor := Author{}
		json.Unmarshal(res, &resAuthor)
		if resAuthor != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPutBookInformation(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		reqBody   Book
		expRes    Book
		expStatus int
	}{
		{"Valid Details", "1", Book{Title: "jk", Author: Author{AuthorId: 1}, Publication: "arihant", PublicationDate: "2009"}, Book{}, 200},
		{"Publication not from Scholastic/pengiun/arihant", "1", Book{Title: "Arvind", Author: Author{AuthorId: 1}, Publication: "Arvind", PublicationDate: "11/03/2002"}, Book{}, http.StatusBadRequest},
		{"Publication date should be between 1880 and 2022", "1", Book{Title: "", Author: Author{AuthorId: 1}, Publication: "", PublicationDate: "1/1/1870"}, Book{}, http.StatusBadRequest},
		{"Publication date should be between 1880 and 2022", "1", Book{Title: "", Author: Author{AuthorId: 1}, Publication: "", PublicationDate: "1/1/2222"}, Book{}, http.StatusBadRequest},
		{"Author not exists", "1", Book{}, Book{}, http.StatusBadRequest},
		{"Title empty", "1", Book{Title: "", Author: Author{AuthorId: 1}, Publication: "", PublicationDate: ""}, Book{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000/book/"+tc.reqId, bytes.NewReader(body))
		putBook(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		res, _ := io.ReadAll(w.Result().Body)
		resBook := Book{}
		json.Unmarshal(res, &resBook)
		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		reqData   Author
		expStatus int
	}{
		{"Valid case ", "1", Author{3, "roli", "siddhu", "23/03/1989", "abhi"}, http.StatusOK},

		{"Valid case BookID not present.", "1000", Author{1000, "Mohan", "chandra", "01/07/2001", "GCC"}, http.StatusBadRequest},
	}

	for i, tc := range testcases {

		body, err := json.Marshal(tc.reqData)
		if err != nil {
			t.Errorf("can't convert data into []byte")
		}
		req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/author/{id}", bytes.NewReader(body))
		res := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"BookID": tc.reqId})
		putAuthor(res, req)
		if res.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test cases fail at %v", i, tc.desc)
		}

	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		expStatus int
	}{
		{"Valid Details", "4", http.StatusOK},
		{"Book not exists", "90", http.StatusNotFound},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/book/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"BookID": tc.reqId})
		deleteBook(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		expStatus int
	}{
		{"Valid Details", "4", http.StatusOK},
		//{"Author not exists", "90", http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/author/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"BookID": tc.reqId})
		deleteAuthor(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}
