package authorservice

import (
	"errors"
	"layeredProject/datastore"
	"layeredProject/entities"
	"strconv"
	"strings"
)

type AuthorService struct {
	datastore datastore.AuthorStorer
}

func New(s datastore.AuthorStorer) AuthorService {
	return AuthorService{s}
}

func (s AuthorService) PostAuthor(a entities.Author) (entities.Author, error) {
	if a.FirstName == "" || !checkDob(a.DOB) {
		return entities.Author{}, nil
	}

	id, err := s.datastore.PostAuthor(a)
	if err != nil || id == -1 {
		return entities.Author{}, err
	}

	a.AuthorID = id
	return a, nil
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
	case year <= 0:
		return false
	}

	return true
}

// PutAuthor - business logic
func (s AuthorService) PutAuthor(a entities.Author) (entities.Author, error) {
	if a.FirstName == "" || !checkDob(a.DOB) {
		return entities.Author{}, errors.New("can't empty")
	}

	id, err := s.datastore.PutAuthor(a)
	if err != nil {
		return entities.Author{}, err
	}
	//if id == -1 {
	//	return entities.Author{}, nil
	//}

	a.AuthorID = id
	return a, nil
}

// DeleteAuthor
func (s AuthorService) DeleteAuthor(id int) (int, error) {
	if id < 0 {
		return -1, nil
	}

	count, err := s.datastore.DeleteAuthor(id)
	if err != nil {
		return -1, err
	}

	return count, nil
}
