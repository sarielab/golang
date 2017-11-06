package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	//sslmode=verify-full
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		fmt.Println("error2", err.Error())
		return
	}
	defer db.Close()

	var stu = student{}
	name := "Jason Bourne"
	err = db.
		QueryRow("SELECT * FROM tb_student WHERE name = $1", name).
		Scan(&stu.id, &stu.name, &stu.age, &stu.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(stu)
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println("error1", err.Error())
		return
	}
	defer db.Close()

	name := "Jerome"
	rows, err := db.Query("SELECT * FROM tb_student WHERE name = $1", name)
	var students []student

	if err != nil {
		fmt.Println("error2", err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var stu = student{}
		var err = rows.Scan(&stu.id, &stu.name, &stu.age, &stu.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		students = append(students, stu)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, stu := range students {
		fmt.Println(stu.name)
	}
}

func sqlPrepare() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT name, grade FROM tb_student WHERE id = $1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("B001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name:	%s\ngrade:	%d\n", result1.name, result1.grade)
	var result2 = student{}
	stmt.QueryRow("B003").Scan(&result2.name, &result2.grade)
	fmt.Printf("name:	%s\ngrade:	%d\n", result2.name, result2.grade)
	var result3 = student{}
	stmt.QueryRow("B002").Scan(&result3.name, &result3.grade)
	fmt.Printf("name:	%s\ngrade:	%d\n", result3.name, result3.grade)
}

func sqlExec() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values($1,$2,$3,$4)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = $1", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}

func main() {
	sqlQuery()
	sqlQueryRow()
	sqlPrepare()
	sqlExec()
}
