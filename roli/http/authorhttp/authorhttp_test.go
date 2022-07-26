package authorhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"layeredProject/entities"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{"valid author", entities.Author{
			1, "suhani", "siddhu", "25/04/2001", "roli"}, http.StatusBadRequest},
		{"exiting author", entities.Author{
			1, "suhani", "siddhu", "25/04/2001", "roli"}, http.StatusBadRequest},
		{"invalid firstname", entities.Author{
			1, "suhan", "siddhu", "25/04/2001", "roli"}, http.StatusBadRequest},
		{"invalid DOB", entities.Author{
			1, "suhani", "siddhu", "25/04/1990", "roli"}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Print(err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PostAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{"valid author", entities.Author{
			2, "abhi", "siddhu", "25/04/2012", "abhi"}, http.StatusCreated},
		{"exiting author", entities.Author{
			1, "suhani", "siddhu", "25/04/2001", "roli"}, http.StatusBadRequest},
		{"invalid firstname", entities.Author{
			1, "", "siddhu", "25/04/2001", "roli"}, http.StatusBadRequest},
		{"invalid DOB", entities.Author{
			1, "suhani", "siddhu", "25/00/2001", "roli"}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Print(err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PostAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc     string
		target   string
		expected int
	}{
		{"valid authorId", "1", http.StatusNoContent},
		{"invalid authorId", "-2", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+tc.target, nil)
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PostAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

type mockService struct{}

func (h mockService) PutAuthor(author2 entities.Author) (entities.Author, error) {
	if author2.AuthorID == 3 {
		return entities.Author{}, nil
	}

	return entities.Author{}, errors.New("invalid constraints")
}

func (h mockService) DeleteAuthor(id int) (int, error) {
	if id == 3 {
		return id, nil
	}

	return -1, errors.New("invalid")
}

func (h mockService) PostAuthor(author2 entities.Author) (entities.Author, error) {
	if author2.AuthorID == 3 {
		return entities.Author{}, nil
	}

	return entities.Author{}, errors.New("invalid constraints")
}
