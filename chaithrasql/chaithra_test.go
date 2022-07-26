package chaithrasql

import (
	"errors"
	"reflect"
	"testing"
)

func TestPost(t *testing.T) {
	s := []struct {
		description string
		input       employee
		output      employee
		output2     error
	}{
		//{"insert", employee{1, "rahul", "intern", 0.0, 30000.0}, Employee{1, "rahul", "intern", 0.0, 30000.0}},
		{"success case", employee{1, "SUHANI", "SDE INTERN", 0.0, 30000.0}, employee{1, "SUHANI", "SDE INTERN", 0.0, 30000.0}, nil},
		{"invalid case", employee{0, "suhi", "dev", 1.0, 35000.0}, employee{}, errors.New("invalid id")},
		{"emp already exist", employee{2, "arvind", "sde", 1.0, 80000.0}, employee{}, errors.New("emp already exists")},
	}

	for _, val := range s {
		res, err := post(val.input)
		if !reflect.DeepEqual(err, val.output2) {
			t.Errorf("case failed")
		}
		if !reflect.DeepEqual(res, val.output) {
			t.Errorf("case failed")
		}
	}
}
func TestGet(t *testing.T) {
	st := []struct {
		description string
		input       int64
		output      employee
		output2     error
	}{
		//{"get", 1, Employee{1, "rahul", "intern", 0.0, 30000.0}},
		{"get", 4, employee{4, "rahul", "developer", 1.0, 60000.0}, nil},
		{"get", 5, employee{5, "anuk", "developer", 2.0, 70000.0}, nil},
		{"emp not exist", 8, employee{}, errors.New("emp does not exist")},
	}
	for _, v := range st {
		res, err := get(v.input)
		if !reflect.DeepEqual(err, v.output2) {
			t.Errorf("case failed")
		}
		if !reflect.DeepEqual(res, v.output) {
			t.Errorf("case failed")
		}

	}
}
