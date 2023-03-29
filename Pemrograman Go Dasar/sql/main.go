package main

import (
	"database/sql"
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type student struct {
	id string
	name string
	age int
	grade int
}

func connect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	// contoh koneksi ke postgres
	db, err := sql.Open("postgres", "user=postgres password=secret dbname=db_belajar_golang sslmode=disable")

	// contoh koneksi ke postgres
	// db, err := sql.Open("postgres", "user=postgres password=secret dbname=test sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 23
	// MySQL
	// rows, err := db.Query("select id, name, age, grade from tb_student where age = ?", age)

	// PostgreSQL
	rows, err := db.Query("select id, name, age, grade from tb_student where age = $1", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

// membaca 1 record data menggunakan QueryRow()
func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "E001"
	// MySQL
	// err = db.QueryRow("select id, name, age, grade from tb_student where id = ?", id).Scan(&result.id, &result.name, &result.age, &result.grade)

	// PostgreSQL
	err = db.QueryRow("select id, name, age, grade from tb_student where id = $1", id).Scan(&result.id, &result.name, &result.age, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

// eksekusi query menggunakan Prepare()
func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// MySQL
	// stmt, err := db.Prepare("select name, grade from tb_student where id = ?")

	// PostgreSQL
	stmt, err := db.Prepare("select name, grade from tb_student where id = $1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	err = stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	err = stmt.QueryRow("B001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	err = stmt.QueryRow("B002").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)
}

// insert, update, delete using Exec()
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// 1. insert data menggunakan cara biasa
	// MySQL
	// _, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)

	// PostgreSQL
	_, err = db.Exec("insert into tb_student values ($1, $2, $3, $4)", "G001", "Galahad", 29, 2)

	// 2. menggunakan metode prepared statement
	// MySQL
	// stmt, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
	// stmt.Exec("G001", "Galahad", 29, 2)

	// PostgreSQL
	// stmt, err := db.Prepare("insert into tb_student values ($1, $2, $3, $4)")
	// stmt.Exec("G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	// update data
	// MySQL
	// _, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")

	// PostgreSQL
	_, err = db.Exec("update tb_student set age = $1 where id = $2", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	// delete data
	// MySQL
	// _, err = db.Exec("delete from tb_student where id = ?", "G001")

	// PostgreSQL
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