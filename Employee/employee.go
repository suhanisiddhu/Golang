package Employee

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Employee struct {
	name       string
	dob        string
	experience int
	out        bool
}

func (e Employee) exp() int {
	//1 year of experience=20000,2 to 5 = 50000, 5 to 10 = 100000, more than 10 200000
	if e.experience == 1 && e.experience > 0 && e.experience < 2 {
		return 2000
	} else if e.experience >= 2 && e.experience <= 5 {
		return 5000
	} else if e.experience > 5 && e.experience <= 10 {
		return 10000
	}
	return 20000
}
func (e Employee) calculateAge() int {
	res := e.dob[len(e.dob)-4:]

	i, _ := strconv.Atoi(res)
	return 2022 - i
}
func (e Employee) printAllDetail(salaryBonus, age int) {
	fmt.Printf("emp name %s emp dob %s emp experience %v emp age %v  emp salaryBonus %d\n", e.name, e.dob, e.experience, age, salaryBonus)
}

func (e Employee) printdetail() bool {
	if strings.TrimSpace(e.name) == "" || strings.TrimSpace(e.dob) == "" || e.experience <= 0 {
		log.Println("invalid details")
		return false
	}

	fmt.Printf("Hello %s, Congratulations for completing %v years\n", e.name, e.experience)
	age := e.calculateAge()
	salaryBonus := e.exp()
	e.printAllDetail(salaryBonus, age)
	return true
}
