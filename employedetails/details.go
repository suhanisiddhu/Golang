package employedetails

type employee struct {
	id   int
	name string
	age  int
}

type output struct {
	bo bool
	e  employee
}

func details(e employee) output {
	if e.age < 22 {
		return output{false, employee{}}
	} else {
		return output{true, e}
	}
}
