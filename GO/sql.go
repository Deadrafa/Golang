package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func sql_reques() {
	connStr := "user=postgres  dbname=registration sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM registration_user;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, email string
		var password sql.NullString
		err := rows.Scan(&id, &username, &email, &password)
		if err != nil {
			panic(err)
		}
		var passwordValue string
		if password.Valid {
			passwordValue = password.String
		} else {
			passwordValue = "<NULL>"
		}
		fmt.Printf("ID: %d, Username %s, Email: %s Password: %s\n", id, username, email, passwordValue)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func check_user_in_db(Username string, Email string) int {
	connStr := "user=postgres  dbname=registration sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, username, email FROM registration_user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var username, email string
		err := rows.Scan(&id, &username, &email)
		if err != nil {
			panic(err)
		}
		if Username == username {
			return 111
		}
		if Email == email {
			return 222
		}
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return 0
}

func add_user_in_db(Username string, Email string, Password string) {
	connStr := "user=postgres dbname=registration sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Используем параметризованный запрос для избежания SQL-инъекций
	stmt, err := db.Prepare("INSERT INTO registration_user (username, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatalf("Ошибка при подготовке SQL-запроса: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(Username, Email, Password)
	if err != nil {
		log.Fatalf("Ошибка при выполнении SQL-запроса: %v", err)
	}

	// Отправка на email после успешной вставки
	// send_to_email(Email)
	fmt.Println("Пользователь успешно добавлен!")
}
