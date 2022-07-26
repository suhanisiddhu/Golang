package authorhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"gofr/entities"
	"gofr/service"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc  string
		input entities.Author

		expected       entities.Author
		expectedStatus int
		expectedErr    error
	}{
		{desc: "valid case:", input: entities.Author{
			AuthorID: 3, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			expected: entities.Author{
				AuthorID: 3, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			expectedStatus: http.StatusCreated, expectedErr: nil,
		},
		{desc: "returning error from service", input: entities.Author{AuthorID: 4, FirstName: "suhani", LastName: "siddhu",
			DOB: "25/04/2001", PenName: "roli"}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: errors.New("not valid constraints"),
		},
	}

	for _, tc := range testcases {
		data, _ := json.Marshal(tc.input)

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()

		if tc.input.AuthorID == 4 {
			mockService.EXPECT().PostAuthor(req.Context(), tc.input).Return(tc.expected, tc.expectedErr)
		} else {
			mockService.EXPECT().PostAuthor(req.Context(), tc.input).Return(tc.expected, tc.expectedErr).AnyTimes()
		}

		mock.PostAuthor(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

// TestPut
func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc           string
		input          entities.Author
		TargetID       string
		expected       entities.Author
		expectedStatus int
		expectedErr    error
	}{
		{desc: "valid case:", input: entities.Author{
			AuthorID: 3, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			TargetID: "4", expected: entities.Author{AuthorID: 3, FirstName: "suhani",
				LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}, expectedStatus: http.StatusCreated,
			expectedErr: nil,
		},

		{desc: "unmarshalling error ", input: entities.Author{}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: nil,
		},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			log.Print(err)
		}

		if tc.desc == "unmarshalling error " {
			data = []byte("hello")
		}

		req := httptest.NewRequest("PUT", "localhost:8000/author/{id}"+tc.TargetID, bytes.NewReader(data))
		req = mux.SetURLVars(req, map[string]string{"id": tc.TargetID})
		w := httptest.NewRecorder()

		mockService.EXPECT().PutAuthor(req.Context(), tc.input).Return(tc.expected, tc.expectedErr).AnyTimes()

		mock.PutAuthor(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

// TestDelete
func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc   string
		target string

		expectedStatus int
		expectedErr    error
	}{
		{"valid authorId", "4", http.StatusNoContent, nil},
		{"invalid authorId", "-5", http.StatusBadRequest, errors.New("invalid id")},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+tc.target, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()

		id, err := strconv.Atoi(tc.target)
		if err != nil {
			log.Print(err)
		}

		mockService.EXPECT().DeleteAuthor(req.Context(), id).Return(tc.expectedStatus, tc.expectedErr).AnyTimes()

		mock.DeleteAuthor(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v Expected :%v Got:%v\n", tc.desc, tc.expectedStatus, res.StatusCode)
		}
	}
}

/*func TestPostAuthor(t *testing.T) {
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
}*/
