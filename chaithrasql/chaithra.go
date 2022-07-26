package chaithrasql

import (
	"database/sql"
	"errors"
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

var db *sql.DB

func get(id int64) (employee, error) {
	db, err := sql.Open("mysql", "root:Suhani@123@tcp(127.0.0.1:3306)/Test")

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	row := db.QueryRow("SELECT * FROM employee WHERE emp_id = ?", id)
	em := employee{}
	if err := row.Scan(&em.emp_id, &em.emp_name, &em.emp_position, &em.emp_experience, &em.emp_salary); err != nil {
		if err == sql.ErrNoRows {
			return employee{}, errors.New("emp does not exist")
		}
	}
	return em, nil
}

func post(em employee) (employee, error) {
	db, err := sql.Open("mysql", "root:Suhani@123@tcp(127.0.0.1:3306)/Test")

	if err != nil {
		log.Print(err)
	}
	emp := employee{}
	//defer db.Close()
	//var em Employee
	if em.emp_id <= 0 {
		return emp, errors.New("invalid id")
	}

	result, err := db.Exec("INSERT INTO employee (emp_id,emp_name, emp_position,emp_experience,emp_salary)VALUES(?,?,?,?,?);", em.emp_id, em.emp_name, em.emp_position, em.emp_experience, em.emp_salary)
	if err != nil {
		return emp, errors.New("emp already exists")
	}
	fmt.Println(result)
	/*err = result.Scan(&em.Emp_id, &em.Emp_name, &em.Emp_position, &em.Emp_experience, &em.Emp_salary)
	  if err != nil {
	     if err == sql.ErrNoRows {
	        return em, fmt.Errorf("post %v: no such emp", em)
	     }
	     return em, fmt.Errorf("post %v: %v", em, err)
	  }*/
	return em, nil

}
