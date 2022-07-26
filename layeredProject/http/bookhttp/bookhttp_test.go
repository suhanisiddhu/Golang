package bookhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"layeredProject/entities"
	"layeredProject/service"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expectedBooks  []entities.Book
		expectedErr    error
		expectedStatus int
	}{
		{desc: "success case", title: "", includeAuthor: "", expectedBooks: []entities.Book{{BookID: 1,
			AuthorID: 1, Title: "jk", Publication: "penguin", PublishedDate: "25/04/2001",
			Author: entities.Author{}}, {BookID: 2, AuthorID: 1, Title: "book 2", Publication: "penguin",
			PublishedDate: "25/04/2010", Author: entities.Author{}}}, expectedErr: nil, expectedStatus: http.StatusOK,
		},
		{desc: "invalid case", title: "book+two", includeAuthor: "true", expectedBooks: []entities.Book{},
			expectedErr: errors.New("does not exist"), expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book?"+"title="+tc.title+"&"+"includeAuthor="+tc.includeAuthor, nil)
		w := httptest.NewRecorder()

		if tc.title == "" {
			mockService.EXPECT().GetAllBook(req.Context(), tc.title, tc.includeAuthor).Return(tc.expectedBooks, tc.expectedErr)
		}

		if tc.title == "book+two" {
			mockService.EXPECT().GetAllBook(req.Context(), "book two", tc.includeAuthor).Return(tc.expectedBooks, tc.expectedErr)
		}

		mock.GetAllBook(w, req)

		res := w.Result()
		if !assert.Equal(t, tc.expectedStatus, res.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

// TestGetBookByID : test the GetBookByID
func TestGetBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	Testcases := []struct {
		desc     string
		targetID string

		expected           entities.Book
		expectedStatusCode int
		expectedErr        error
	}{
		{desc: "fetching book by id", targetID: "1", expected: entities.Book{BookID: 1, AuthorID: 1, Title: "jk rowling",
			Publication: "penguin", PublishedDate: "25/04/1990", Author: entities.Author{AuthorID: 1, FirstName: "suhani",
				LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}, expectedStatusCode: http.StatusOK, expectedErr: nil,
		},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book/{id}"+tc.targetID, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})

		id, _ := strconv.Atoi(tc.targetID)
		if tc.targetID != "-1" {
			mockService.EXPECT().GetBookByID(req.Context(), id).Return(tc.expected, tc.expectedErr)
		}

		mock.GetBookByID(w, req)

		result := w.Result()

		if !reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

// TestPost : test the post
func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc string
		body entities.Book

		expected           entities.Book
		expectedErr        error
		expectedStatusCode int
	}{
		{desc: "invalid case", body: entities.Book{BookID: 0, AuthorID: 1, Title: "3 mistakes", Publication: "penguin",
			PublishedDate: "25/04/2010", Author: entities.Author{}},
			expected: entities.Book{}, expectedErr: errors.New("error"),
			expectedStatusCode: http.StatusBadRequest,
		},

		{desc: "valid case", body: entities.Book{BookID: 0, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "25/04/2010", Author: entities.Author{}},
			expected: entities.Book{BookID: 15, AuthorID: 1, Title: "jk", Publication: "penguin",
				PublishedDate: "25/04/2010", Author: entities.Author{}},
			expectedErr: nil, expectedStatusCode: http.StatusOK,
		},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		c := "unmarshalling error"
		if tc.desc == c {
			data = []byte("suhani")
		}

		req := httptest.NewRequest("POST", "localhost:8000/book", bytes.NewBuffer(data))
		w := httptest.NewRecorder()

		if tc.desc != c {
			mockService.EXPECT().PostBook(req.Context(), &tc.body).Return(tc.expected, tc.expectedErr)
		}

		mock.PostBook(w, req)

		result := w.Result()

		if !reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

// TestPut : test the put
func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc    string
		input   entities.Book
		inputID string

		expected           entities.Book
		expectedErr        error
		expectedStatusCode int
	}{
		{desc: "invalid case", input: entities.Book{BookID: 0, AuthorID: 1, Title: "jk", Publication: "penguin",
			PublishedDate: "25/04/2010", Author: entities.Author{}}, inputID: "2",
			expected: entities.Book{}, expectedErr: errors.New("error"),
			expectedStatusCode: http.StatusBadRequest,
		},

		{desc: "valid case", input: entities.Book{BookID: 15, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "25/04/2010", Author: entities.Author{}}, inputID: "4",
			expected: entities.Book{BookID: 15, AuthorID: 1, Title: "jk", Publication: "penguin",
				PublishedDate: "25/04/2010", Author: entities.Author{}},
			expectedErr: nil, expectedStatusCode: http.StatusOK,
		},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.inputID, bytes.NewBuffer(data))
		req = mux.SetURLVars(req, map[string]string{"id": tc.inputID})
		w := httptest.NewRecorder()

		id, _ := strconv.Atoi(tc.inputID)
		mockService.EXPECT().PutBook(req.Context(), &tc.input, id).Return(tc.expected, tc.expectedErr)
		mock.PutBook(w, req)

		result := w.Result()

		if !reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

// TestDelete : test the delete book Handler
func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc    string
		inputID string

		expectedStatus int
		expectedErr    error
	}{
		{"valid id", "1", http.StatusNoContent, nil},
		{"invalid id", "-1", http.StatusBadRequest, errors.New("something wrong")},
	}

	for _, tc := range testcases {
		id, _ := strconv.Atoi(tc.inputID)
		w := httptest.NewRecorder()
		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.inputID, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.inputID})

		ctx := req.Context()

		if tc.desc != "invalid id" {
			mockService.EXPECT().DeleteBook(ctx, id).Return(1, tc.expectedErr)
		}

		mock.DeleteBook(w, req)

		result := w.Result()
		if !reflect.DeepEqual(tc.expectedStatus, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

/*func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expectedBooks      []entities.Book
		expectedStatusCode int
	}{
		{"getting all books", "", "", []entities.Book{{1,
			1, "book one", "scholastic", "20/06/2018", entities.Author{}},
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}},
			http.StatusOK},
		{"getting book with author and  title", "jk", "true", []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{1, "suhani",
				"siddhu", "25/04/2001", "sk"}}}, http.StatusOK},
		{"getting book without author", "jk", "true", []entities.Book{
			{2, 1, "jk", "penguin", "25/04/2000", entities.Author{}}}, http.StatusOK},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book?"+"title="+tc.title+"&"+"includeAuthor="+tc.includeAuthor, nil)
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.GetAllBook(w, req)

		result := w.Result()
		body, err := io.ReadAll(result.Body)
		if err != nil {
			log.Print(err)
		}

		var books []entities.Book

		_ = json.Unmarshal(body, &books)
		if reflect.DeepEqual(books, tc.expectedBooks) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID string

		expectedBook       entities.Book
		expectedStatusCode int
	}{
		{"fetching book by id",
			"1", entities.Book{1, 1, "jk rowling", "penguin", "24/04/1990",
				entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusOK},

		{"invalid id", "-1", entities.Book{}, http.StatusBadRequest},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book/{id}"+tc.targetID, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})
		h := New(mockService{})

		h.GetBookByID(w, req)

		result := w.Result()
		body, err := io.ReadAll(result.Body)
		if err != nil {
			log.Print(err)
		}

		var books entities.Book

		_ = json.Unmarshal(body, &books)
		if reflect.DeepEqual(books, tc.expectedBook) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedStatusCode int
	}{
		{"valid case", entities.Book{2, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusCreated},
		{"already existing book", entities.Book{2, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid bookID", entities.Book{-1, 1, "jk", "penguin",
			"20/03/2010", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid author's DOB", entities.Book{2, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/00/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid title", entities.Book{2, 1, "", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusBadRequest},
		{"invalid publication", entities.Book{2, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid published date", entities.Book{4, 1, "jk", "penguin",
			"00/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusBadRequest},
	}
	for _, tc := range testcases {
		b, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/book", bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PostBook(w, req)

		result := w.Result()
		if reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedStatusCode int
	}{
		{"inserting book", entities.Book{3, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusCreated},
		{"already existing book", entities.Book{4, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
			http.StatusCreated},
		{"invalid bookID", entities.Book{-1, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid DOB", entities.Book{2, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/00/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid title", entities.Book{2, 1, "", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusBadRequest},
		{"invalid publication", entities.Book{2, 1, "jk", "penguin",
			"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
			http.StatusBadRequest},
		{"invalid published date", entities.Book{2, 1, "jk", "penguin",
			"00/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, http.StatusBadRequest},
	}
	for _, tc := range testcases {
		b, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		req := httptest.NewRequest("PUT", "localhost:8000/book", bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PutBook(w, req)

		result := w.Result()
		if reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID string

		expectedStatus int
	}{
		{"valid id", "1", http.StatusNoContent},
		{"invalid id", "-1", http.StatusBadRequest},
	}

	for _, tc := range testcases {

		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.targetID, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.DeleteBook(w, req)

		result := w.Result()
		if reflect.DeepEqual(tc.expectedStatus, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}

}*/
