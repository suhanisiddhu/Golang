package authorservice

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"gofr/datastore"
	"gofr/entities"
	"reflect"
	"testing"
)

func TestAuthorPost(t *testing.T) {
	testcases := []struct {
		desc     string
		req      entities.Author
		response entities.Author
		err      error
	}{
		{desc: "valid details", req: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			response: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "suhani", DOB: "25/04/2001", PenName: "roli"}, err: nil},
		{desc: "invalid id", req: entities.Author{AuthorID: -15, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			err: errors.New("invalid")},
	}

	ctr := gomock.NewController(t)
	mockdataStore := datastore.NewMockAuthorStorer(ctr)
	mock := New(mockdataStore)

	for i, v := range testcases {
		if v.desc == "valid details" {
			mockdataStore.EXPECT().PostAuthor(context.TODO(), v.req).Return(v.response, v.err)
		}

		resp, err := mock.PostAuthor(context.TODO(), v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.response)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

func TestAuthorPut(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		req  entities.Author
		err  error
	}{
		{desc: "valid", id: "1", req: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu",
			DOB: "25/04/2001", PenName: "roli"}},
		{desc: "invalid id", id: "-15", err: errors.New("invalid id")},
	}

	ctr := gomock.NewController(t)
	mockdataStore := datastore.NewMockAuthorStorer(ctr)
	mock := New(mockdataStore)

	for i, v := range testcases {
		if v.desc == "valid" {
			mockdataStore.EXPECT().PutAuthor(context.TODO(), v.req).Return(v.req, v.err)
		}

		resp, err := mock.PutAuthor(context.TODO(), v.req)

		if !reflect.DeepEqual(resp, v.req) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.req)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

func TestAuthorDelete(t *testing.T) {
	testcases := []struct {
		desc        string
		id          int
		rowaffected int
		err         error
	}{
		{desc: "valid", id: 1, rowaffected: 1, err: nil},

		//{desc: "invalid id", id: 20, err: errors.New("invalid")},
	}
	ctr := gomock.NewController(t)
	mockdataStore := datastore.NewMockAuthorStorer(ctr)
	mock := New(mockdataStore)

	for i, v := range testcases {
		if v.desc == "valid" {
			mockdataStore.EXPECT().DeleteAuthor(context.TODO(), v.id).Return(v.rowaffected, v.err)
		}

		resp, err := mock.DeleteAuthor(context.TODO(), v.id)

		if !reflect.DeepEqual(resp, v.rowaffected) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.rowaffected)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

/*func TestPostAuthor(t *testing.T) {
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
		}
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

		id, err := m.DeleteAuthor(tc.target)*/

//if err != nil {
//	log.Print(err)
//	}
//assert.Equal(t, id, tc.expectedID)
//if id != tc.expectedID {
//	t.Errorf("failed for %v\n", tc.desc)
//}
/*}
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
}*/
