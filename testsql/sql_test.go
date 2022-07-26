package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	testcases := []struct {
		des    string
		input  int64
		output employee
		expErr error
	}{
		{"inserted", 1, employee{1, "SUHANI", "SDE INTERN", 0, 30000}, nil},
		{"inserted", 2, employee{}, errors.New("emp not exists")},
	}
	for i, tc := range testcases {
		res, err := get(tc.input)
		if res != tc.output {
			t.Errorf("%v test failed %v", i, tc.des)
		}

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%v test failed%v", i, tc.des)
		}
	}
}

/*func TestPost(t *testing.T) {
s := []struct {
	des    string
	input  employee
	output employee
}{
	{"insert value", employee{1, "SUHANI", "SDE INTERN", 0.0, 30000.00}, employee{1, "SUHANI", "SDE INTERN", 0.0, 30000.00}},
	//{"insert values", employee{2, "arvind", "sde", 1.0, 80000.0}, {2, "arvind", "sde", 1.0, 80000.0}},
	{"insert value", employee{3, "anukriti", "sde intern", 0.0, 30000.0}, employee{3, "anukriti", "sde intern", 0.0, 30000.0}},
}
for _, v := range s {
	res, _ := post(v.input)
	if res != v.output {
		t.Errorf("failed")
	}
}*/
