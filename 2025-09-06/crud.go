package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABle IF NOT EXISTS users(
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	// create
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("George", "george7@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	// read
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		log.Println(id, name, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// read row
	var email string
	err = db.QueryRow("SELECT email FROM users WHERE id = ?", 2).Scan(&email)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(email)

	// update
	stmt2, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt2.Exec("TOM", 2)
	if err != nil {
		log.Fatal(err)
	}

	// delete
	stmt3, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt3.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
}
