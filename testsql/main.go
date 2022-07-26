package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	emp_id         int
	emp_name       string
	emp_position   string
	emp_experience float64
	emp_salary     float64
}

func get(id int64) (employee, error) {
	db, err := sql.Open("mysql", "root:Suhani@123@tcp(127.0.0.1:3306)/Test")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var em employee
	//if em.Emp_id<0{

	//}
	row := db.QueryRow("SELECT * FROM employee WHERE emp_id = ?", id)
	if err := row.Scan(&em.emp_id, &em.emp_name, &em.emp_position, &em.emp_experience, &em.emp_salary); err != nil {
		if err == sql.ErrNoRows {
			return em, fmt.Errorf("get %d: no such album", id)
		}
		return em, fmt.Errorf("get %d: %v", id, err)
	}
	return em, nil
}

/*func post(emp employee) (employee, error) {
	db, err := sql.Open(dbName, dbSource)
	if err != nil {
		log.Print(err)
	}
	//defer db.Close()
	if emp.empId <= 0 {
		return employee{}, errors.New("invalid id")
	}
	query := fmt.Sprintf("INSERT INTO employee (emp_id, emp_name, emp_position, emp_experience, emp_salary)\nVALUES\n(%v,'%v','%v',%v,'%v');", emp.empId, emp.empName, emp.empPos, emp.empExp, emp.empSalary)
	data, err := db.Exec(query)
	if err != nil {
		return employee{}, errors.New("emp Already Exists")
	}
	fmt.Println(data)
	return emp, nil
}*/
