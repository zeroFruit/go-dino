package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // as soon as package load init run
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	// connect to the db
	db, err := sql.Open("mysql", "root:111111@/dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// general query with args
	rows, err := db.Query("SELECT * FROM dino.animals WHERE age>?", 10)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{} // init slice, we don't have initial vaue and count
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)

	// query a single row
	row := db.QueryRow("SELECT * FROM dino.animals where id=?", 2)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
	/*
		// insert data
		result, err := db.Exec(
			"INSERT INTO dino.animals (animal_type, nickname, zone, age) VALUES (?, ?, ?, ?)",
			"Carnotaurus", "Carno", 3, 22)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	*/
	// update a row
	age := 14
	id := 2
	result, err := db.Exec("UPDATE dino.animals SET age=? WHERE id=?", age, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

}
