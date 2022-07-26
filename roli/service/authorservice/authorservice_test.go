package authorservice

import (
	"github.com/stretchr/testify/assert"
	"layeredProject/entities"
	"log"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected entities.Author
	}{
		{"valid author", entities.Author{
			3, "suhani", "siddhu", "25/04/2001", "roli"},
			entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}},
		{"existing author", entities.Author{
			2, "suhani", "siddhu", "25/04/2001", "roli"}, entities.Author{}},
		{"invalid firstname", entities.Author{
			2, "", "siddhu", "25/04/2001", "roli"}, entities.Author{}},
		{"invalid dob", entities.Author{
			2, "suhani", "siddhu", "25/00/2001", "roli"}, entities.Author{}},
	}

	for _, tc := range testcases {

		m := New(mockStore{})
		id, err := m.PostAuthor(tc.body)
		if err != nil {
			log.Print(err)
		}
		assert.Equal(t, id, tc.expected)
		//if id != tc.expected {
		//	t.Errorf("failed for %v\n", tc.desc)
		//}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected entities.Author
	}{
		{"valid author", entities.Author{
			3, "suhani", "siddhu", "25/04/2001", "roli"}, entities.Author{3, "suhani", "siddhu", "25/04/2001", "roli"}},
		{"existing author", entities.Author{
			2, "suhani", "siddhu", "25/04/2001", "siddhu"}, entities.Author{}},
		{"invalid firstname", entities.Author{
			2, "suhani", "siddhu", "25/04/2001", "roli"}, entities.Author{}},
		{"invalid dob", entities.Author{
			2, "suhani", "siddhu", "25/04/2001", "roli"}, entities.Author{}},
	}

	for _, tc := range testcases {

		m := New(mockStore{})
		a, err := m.PutAuthor(tc.body)
		if err != nil {
			log.Print(err)
		}

		/*	if a != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}*/
		assert.Equal(t, a, tc.expected)
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		//input
		desc   string
		target int
		//output
		expectedID int
	}{
		{"valid authorId", 2, 2},
		{"invalid authorId", -1, -1},
	}

	for _, tc := range testcases {

		m := New(mockStore{})

		id, err := m.DeleteAuthor(tc.target)

		if err != nil {
			log.Print(err)
		}
		assert.Equal(t, id, tc.expectedID)
		//if id != tc.expectedID {
		//	t.Errorf("failed for %v\n", tc.desc)
		//}
	}
}

type mockStore struct{}

func (m mockStore) PostAuthor(author2 entities.Author) (int, error) {
	if author2.AuthorID == 3 {
		return author2.AuthorID, nil
	} else {
		return -1, nil
	}
}

func (m mockStore) PutAuthor(author2 entities.Author) (int, error) {

	if author2.AuthorID == 3 {
		return author2.AuthorID, nil
	} else {
		return -1, nil
	}
}

func (m mockStore) DeleteAuthor(id int) (int, error) {
	if id <= 0 {
		return -1, nil
	} else {
		return id, nil
	}
}
