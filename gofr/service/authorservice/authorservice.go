package authorservice

import (
	"context"
	"errors"
	"gofr/datastore"
	"gofr/entities"
	"strconv"
	"strings"
)

type AuthorService struct {
	datastore datastore.AuthorStorer
}

func New(s datastore.AuthorStorer) AuthorService {
	return AuthorService{s}
}

func (s AuthorService) PostAuthor(ctx context.Context, a entities.Author) (entities.Author, error) {
	if a.FirstName == "" || !checkDob(a.DOB) || a.AuthorID < 0 || a.LastName == "" || a.PenName == "" {
		return entities.Author{}, errors.New("invalid params")
	}

	id, err := s.datastore.PostAuthor(ctx, a)
	if err != nil {
		return entities.Author{}, errors.New("error from db")
	}

	a.AuthorID = id

	return a, nil
}

// PutAuthor - business logic
func (s AuthorService) PutAuthor(ctx context.Context, a entities.Author) (entities.Author, error) {
	if a.FirstName == "" || !checkDob(a.DOB) || a.AuthorID < 0 || a.LastName == "" || a.PenName == "" {
		return entities.Author{}, errors.New("can't empty")
	}

	id, err := s.datastore.PutAuthor(ctx, a)
	if err != nil {
		return entities.Author{}, err
	}

	a.AuthorID = id

	return a, nil
}

// DeleteAuthor
func (s AuthorService) DeleteAuthor(ctx context.Context, id int) (int, error) {
	if id < 0 {
		return -1, errors.New("invalid")
	}

	count, err := s.datastore.DeleteAuthor(ctx, id)
	if err != nil {
		return -1, errors.New("invalid")
	}

	return count, nil
}
func checkDob(dob string) bool {
	Dob := strings.Split(dob, "/")
	day, _ := strconv.Atoi(Dob[0])
	month, _ := strconv.Atoi(Dob[1])
	year, _ := strconv.Atoi(Dob[2])

	switch {
	case day <= 0 || day > 31:
		return false
	case month <= 0 || month > 12:
		return false
	case year <= 0 || year > 2022:
		return false
	}

	return true
}
