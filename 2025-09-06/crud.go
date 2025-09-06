package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int
	NAME  string
	EMAIL string
}

func createUser(db *sql.DB, name, email string) error {
	stmt, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, email)
	if err != nil {
		return err
	}

	return nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.NAME, &u.EMAIL); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func getUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.NAME, &user.EMAIL)
	if err != nil {
		return user, err
	}

	return user, nil
}

func updateUser(db *sql.DB, id int, name string) (int64, error) {
	stmt, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func deleteUser(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE from users WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

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

	err = createUser(db, "Bob", "bob3@gmail.com")
	if err != nil {
		log.Fatal("ユーザーの作成に失敗しました：", err)
	}

	// read
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal("ユーザーの取得に失敗しました：", err)
	}
	for _, u := range users {
		log.Printf("ID: %d, NAME: %s, EMAIL: %s\n", u.ID, u.NAME, u.EMAIL)
	}

	// read row
	user, err := getUserByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

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
	rowsAffected, err := deleteUser(db, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d行のレコードが削除されました\n", rowsAffected)
}
