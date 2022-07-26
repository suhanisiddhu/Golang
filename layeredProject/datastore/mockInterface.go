// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package datastore is a generated GoMock package.
package datastore

import (
	context "context"
	entities "layeredProject/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorStorer is a mock of AuthorStorer interface.
type MockAuthorStorer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorStorerMockRecorder
}

// MockAuthorStorerMockRecorder is the mock recorder for MockAuthorStorer.
type MockAuthorStorerMockRecorder struct {
	mock *MockAuthorStorer
}

// NewMockAuthorStorer creates a new mock instance.
func NewMockAuthorStorer(ctrl *gomock.Controller) *MockAuthorStorer {
	mock := &MockAuthorStorer{ctrl: ctrl}
	mock.recorder = &MockAuthorStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorStorer) EXPECT() *MockAuthorStorerMockRecorder {
	return m.recorder
}

// DeleteAuthor mocks base method.
func (m *MockAuthorStorer) DeleteAuthor(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthor", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAuthor indicates an expected call of DeleteAuthor.
func (mr *MockAuthorStorerMockRecorder) DeleteAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthor", reflect.TypeOf((*MockAuthorStorer)(nil).DeleteAuthor), arg0, arg1)
}

// PostAuthor mocks base method.
func (m *MockAuthorStorer) PostAuthor(arg0 context.Context, arg1 entities.Author) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostAuthor", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostAuthor indicates an expected call of PostAuthor.
func (mr *MockAuthorStorerMockRecorder) PostAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostAuthor", reflect.TypeOf((*MockAuthorStorer)(nil).PostAuthor), arg0, arg1)
}

// PutAuthor mocks base method.
func (m *MockAuthorStorer) PutAuthor(arg0 context.Context, arg1 entities.Author) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutAuthor", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutAuthor indicates an expected call of PutAuthor.
func (mr *MockAuthorStorerMockRecorder) PutAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutAuthor", reflect.TypeOf((*MockAuthorStorer)(nil).PutAuthor), arg0, arg1)
}

// MockBookStorer is a mock of BookStorer interface.
type MockBookStorer struct {
	ctrl     *gomock.Controller
	recorder *MockBookStorerMockRecorder
}

// MockBookStorerMockRecorder is the mock recorder for MockBookStorer.
type MockBookStorerMockRecorder struct {
	mock *MockBookStorer
}

// NewMockBookStorer creates a new mock instance.
func NewMockBookStorer(ctrl *gomock.Controller) *MockBookStorer {
	mock := &MockBookStorer{ctrl: ctrl}
	mock.recorder = &MockBookStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookStorer) EXPECT() *MockBookStorerMockRecorder {
	return m.recorder
}

// DeleteBook mocks base method.
func (m *MockBookStorer) DeleteBook(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookStorerMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookStorer)(nil).DeleteBook), arg0, arg1)
}

// GetAllBook mocks base method.
func (m *MockBookStorer) GetAllBook(ctx context.Context, string2, string3 string) ([]entities.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBook", ctx, string2, string3)
	ret0, _ := ret[0].([]entities.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBook indicates an expected call of GetAllBook.
func (mr *MockBookStorerMockRecorder) GetAllBook(ctx, string2, string3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBook", reflect.TypeOf((*MockBookStorer)(nil).GetAllBook), ctx, string2, string3)
}

// GetBookByID mocks base method.
func (m *MockBookStorer) GetBookByID(arg0 context.Context, arg1 int) (entities.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", arg0, arg1)
	ret0, _ := ret[0].(entities.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockBookStorerMockRecorder) GetBookByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockBookStorer)(nil).GetBookByID), arg0, arg1)
}

// PostBook mocks base method.
func (m *MockBookStorer) PostBook(ctx context.Context, book *entities.Book) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBook", ctx, book)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBook indicates an expected call of PostBook.
func (mr *MockBookStorerMockRecorder) PostBook(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBook", reflect.TypeOf((*MockBookStorer)(nil).PostBook), ctx, book)
}

// PutBook mocks base method.
func (m *MockBookStorer) PutBook(ctx context.Context, book *entities.Book, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBook", ctx, book, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBook indicates an expected call of PutBook.
func (mr *MockBookStorerMockRecorder) PutBook(ctx, book, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBook", reflect.TypeOf((*MockBookStorer)(nil).PutBook), ctx, book, id)
}
