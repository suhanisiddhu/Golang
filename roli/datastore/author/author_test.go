package author

import (
	"layeredProject/driver"
	"layeredProject/entities"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expectedID int
	}{
		{"valid author", entities.Author{
			4, "suhani", "siddhu", "25/04/2001", "roli"}, 4},
		{"exiting author", entities.Author{
			1, "suhani", "siddhu", "25/04/2001", "roli"}, 1},
	}

	for _, tc := range testcases {
		Db := driver.Connection()
		authorStore := New(Db)

		id, _ := authorStore.PostAuthor(tc.body)

		if id != tc.expectedID {
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
			4, "suhani", "siddhu", "25/04/2001", "roli"}, -1},
		{"exiting author", entities.Author{
			1, "suhani", "siddhu", "25/04/2001", "roli"}, 1},
	}
	for _, tc := range testcases {
		Db := driver.Connection()
		authorStore := New(Db)

		id, _ := authorStore.PostAuthor(tc.body)

		if id != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc     string
		target   int
		expected int
	}{
		{"valid authorId", 4, 4},
		{"invalid authorId", -2, 0},
	}

	for _, tc := range testcases {
		Db := driver.Connection()
		authorStore := New(Db)

		count, _ := authorStore.DeleteAuthor(tc.target)

		if count != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
