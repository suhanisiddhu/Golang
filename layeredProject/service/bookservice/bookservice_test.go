package bookservice

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"github.com/golang/mock/gomock"
	"layeredProject/datastore"
	"layeredProject/entities"

)

func TestBookPost(t *testing.T) {
	testcases := []struct {
		desc     string
		req      entities.Book
		response entities.Book
		err      error
	}{
		{desc: "valid details", req: entities.Book{BookID: 1, AuthorID: 1,
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"},
			response: entities.Book{BookID: 1,
				AuthorID: 1, Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
				Title: "jk", Publication: "penguin", PublishedDate: "25/04/2010"}},
		{desc: "valid details", req: entities.Book{BookID: 2, AuthorID: 1,
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"},
			response: entities.Book{BookID: 2, AuthorID: 1,
				Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
				Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"}},
	}

	for i, v := range testcases {
		ctr := gomock.NewController(t)
		mockdatastore := datastore.NewMockBookStorer(ctr)
		mock := New(mockdatastore)

		if v.desc == "valid details" {
			mockdatastore.EXPECT().PostBook(context.TODO(), &v.req).Return(v.response, v.err)
		}

		resp, err := mock.PostBook(context.TODO(), &v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.response)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

func TestBookGetAll(t *testing.T) {
	testcases := []struct {
		desc string
		resp []entities.Book
		err  error
	}{
		{desc: "valid details ", resp: []entities.Book{{BookID: 1, AuthorID: 1,
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"}}},
	}

	for i, v := range testcases {
		ctr := gomock.NewController(t)
		mockdatastore := datastore.NewMockBookStorer(ctr)
		mock := New(mockdatastore)

		mockdatastore.EXPECT().GetAllBook(context.TODO(), "", "true").Return(v.resp, v.err)

		resp, err := mock.GetAllBook(context.TODO(), "", "true")

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

func TestBookGetbyid(t *testing.T) {
	testcases := []struct {
		desc string
		id   int
		resp entities.Book
		err  error
	}{
		{desc: "valid detail", id: 1, resp: entities.Book{BookID: 1, AuthorID: 1,
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"}},
	}

	ctr := gomock.NewController(t)
	mockdataStore := datastore.NewMockBookStorer(ctr)
	mock := New(mockdataStore)

	for i, v := range testcases {
		if v.desc == "valid detail" {
			mockdataStore.EXPECT().GetBookByID(context.TODO(), v.id).Return(v.resp, v.err)
		}

		resp, err := mock.GetBookByID(context.TODO(), v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

func TestBookPut(t *testing.T) {
	testcases := []struct {
		desc string
		id   int
		req  entities.Book
		resp entities.Book
		err  error
	}{
		{desc: "valid ", id: 1, req: entities.Book{BookID: 1, AuthorID: 1,
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"},
			resp: entities.Book{BookID: 1, AuthorID: 1,
				Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
				Title:  "jk", Publication: "penguin", PublishedDate: "25/04/2010"}},
		{desc: "invalid id", req: entities.Book{BookID: 1, AuthorID: 1,
			Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"},
			Title:  "jk", Publication: "arihant", PublishedDate: "25/04/2010"},
			id: -13, err: errors.New("invalid id")},
	}

	for i, v := range testcases {
		ctr := gomock.NewController(t)
		mockdataStore := datastore.NewMockBookStorer(ctr)
		mock := New(mockdataStore)

		if v.desc == "valid" {
			mockdataStore.EXPECT().PutBook(context.TODO(), v.id, &v.req).Return(v.resp, v.err)
		}

		resp, err := mock.PutBook(context.TODO(), &v.req, v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

func TestBookDelete(t *testing.T) {
	testcases := []struct {
		desc        string
		id          int
		rowaffected int
		err         error
	}{
		{desc: "valid", id: 1, rowaffected: 1},
		//{desc: "invalid id", id: 10, err: errors.New("invalid")},
	}

	ctr := gomock.NewController(t)
	mockdataStore := datastore.NewMockBookStorer(ctr)
	mock := New(mockdataStore)

	for i, v := range testcases {
		if v.desc == "valid" {
			mockdataStore.EXPECT().DeleteBook(context.TODO(), v.id).Return(v.rowaffected, v.err)
		}

		resp, err := mock.DeleteBook(context.TODO(), v.id)

		if !reflect.DeepEqual(resp, v.rowaffected) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.rowaffected)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

/*func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expected []entities.Book
	}{
		{"all books details", "", "", []entities.Book{{1,
			1, "jk rowling", "penguin", "25/04/1990", entities.Author{}},
			{1, 1, "jk rowling", "penguin", "25/04/1990", entities.Author{}}},
		},
		{"getting book with author and given title", "jk", "true", []entities.Book{
			{2, 1, "jk", "penguin", "25/04/1990", entities.Author{1, "suhani",
				"siddhu", "25/04/2001", "roli"}}},
		},
	}

	for _, tc := range Testcases {
		b := New(mockStore{})
		book := b.GetAllBook(tc.title, tc.includeAuthor)

		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID int

		expectedBody entities.Book
		expectedErr  error
	}{
		{"getting book by id",
			1, entities.Book{1, 1, "jk rowling", "penguin",
				"25/04/2000", entities.Author{1, "suhani", "siddhu", "25/04/2001", "roli"}}, nil},

		{"invalid id", -1, entities.Book{}, errors.New("invalid id")},
	}

	for _, tc := range Testcases {
		b := New(mockStore{})
		book, err := b.GetBookByID(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBody) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{desc: "already exists", body: entities.Book{BookID: 1, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "24/04/1990", Author: entities.Author{AuthorID: 1, FirstName: "suhani",
				LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}},
		{desc: "invalid bookID", body: entities.Book{BookID: -1, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "24/04/1990", Author: entities.Author{AuthorID: 1, FirstName: "suhani",
				LastName: "siddhu", DOB: "25/04/2001", PenName: "roli"}}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})

		book, err := b.PostBook(&tc.body)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBook) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{desc: "inserting  book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "jk", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu",
				DOB: "25/04/2001", PenName: "roli"}}, expectedBook: entities.Book{BookID: 4, AuthorID: 1, Title: "jk",
			Publication: "penguin", PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1,
				FirstName: "suhani", LastName: "suhani", DOB: "25/04/2001", PenName: "roli"}}},

		{desc: "updating book", body: entities.Book{BookID: 3, AuthorID: 1, Title: "jk rowling", Publication: "penguin",
			PublishedDate: "25/04/2000", Author: entities.Author{AuthorID: 1, FirstName: "suhani", LastName: "siddhu",
				DOB: "25/04/2001", PenName: "roli"}}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})

		book, err := b.PostBook(&tc.body)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBook) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID int

		expected int
	}{
		{"valid id", 1, 1},
		{"invalid id", -1, -1},
	}

	for _, tc := range testcases {
		b := New(mockStore{})

		id, err := b.DeleteBook(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(id, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}

	}

}*/
