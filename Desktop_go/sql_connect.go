package main

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lib/pq"
)

func sql_connect(info User_info, myWindow fyne.Window) {
	connStr := fmt.Sprintf("user=%s  dbname=%s sslmode=disable", info.db_user, info.db_name)

	db, err_open := sql.Open("postgres", connStr)
	if err_open != nil {
		dialog.ShowInformation("Ошибка", "Подключение к базе данных не удалось!", myWindow)
		return
	}
	defer db.Close()

	// Проверяем статус подключения к базе данных
	err := db.Ping()
	if err != nil {
		dialog.ShowInformation("Ошибка", "Пинг к базе данных не удался!", myWindow)
		return
	}
	// Если соединение успешно
	dialog.ShowInformation("Успех", "Подключение к базе данных успешно!", myWindow)
	fmt.Println("Подключение к базе данных успешно!")

	check_user_in_db(info)
}

func check_user_in_db(info User_info) [][]string {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", info.db_user, info.db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, email FROM registration_user ORDER BY username DESC")
	if err != nil {
		log.Println("Ошибка выполнения запроса:", err)
		return nil
	}
	defer rows.Close()

	var data [][]string
	for rows.Next() {
		var username, email string
		err := rows.Scan(&username, &email)
		if err != nil {
			log.Println("Ошибка чтения строки:", err)
			continue
		}
		data = append(data, []string{username, email})
	}

	if err = rows.Err(); err != nil {
		log.Println("Ошибка при итерации по строкам:", err)
	}
	return data
}

func createTable(data [][]string) *widget.Table {
	// Возвращаем новую таблицу с заглушкой
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 3 // количество строк и колонок
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Заполняется...")
		},

		func(id widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if id.Col == 0 {
				label.SetText(fmt.Sprintf("%d", id.Row+1)) // номер строки
			} else if id.Col == 1 {
				label.SetText(data[id.Row][0]) // имя пользователя
			} else if id.Col == 2 {
				label.SetText(data[id.Row][1]) // email пользователя
			}
		},
	)
	return table
}

func updateTable(table *widget.Table, data [][]string) {
	table.Length = func() (int, int) {
		return len(data), 3 // обновляем количество строк и колонок
	}
	table.UpdateCell = func(id widget.TableCellID, o fyne.CanvasObject) {
		label := o.(*widget.Label)
		if id.Col == 0 {
			label.SetText(fmt.Sprintf("%d", id.Row+1)) // номер строки
		} else if id.Col == 1 {
			label.SetText(data[id.Row][0]) // имя пользователя
		} else if id.Col == 2 {
			label.SetText(data[id.Row][1]) // email пользователя
		}
	}
	table.Refresh() // Обновляем таблицу
}

func editing_db(info User_info) int {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", info.db_user, info.db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		return -1
	}
	defer db.Close()

	switch info.trigger {
	case 1:
		req := fmt.Sprintf("DELETE FROM registration_user WHERE username='%s'", info.username)
		request, err := db.Query(req)
		if err != nil {
			log.Println("Ошибка выполнения запроса:", err)
			return -1
		}
		for request.Next() {
			var email string

			err_s := request.Scan(&email)
			if err_s != nil {
				log.Printf("ERROR %v", err_s)
			}
			fmt.Printf("EMAIL: %s\n", email)
		}
		defer request.Close()
	case 2:
		req := fmt.Sprintf("INSERT INTO registration_user (username, email, password)  VALUES ('%s', '%s', '%s')", info.username, info.email, info.password)
		request, err := db.Query(req)
		if err != nil {
			log.Println("Ошибка выполнения запроса:", err)
			return -1
		}
		defer request.Close()
		return 787
	case 3:
		request, err := db.Query("SELECT email FROM registration_user")
		if err != nil {
			log.Println("Ошибка выполнения запроса:", err)
			return -1
		}
		defer request.Close()
	}
	return 777
}

func add_req_sql(name, email, passord string) []string {
	var list []string
	list = append(list, name, email, passord)
	return list
}

func err_sql(name string, trigger int, myWindow fyne.Window) {
	switch trigger {
	case 777:
		str := fmt.Sprintf("Пользователь  %s удалён!", name)
		dialog.ShowInformation("Успех", str, myWindow)
	case -1:
		dialog.ShowInformation("Ошибка", "Удаление не удалось!", myWindow)
	case 787:
		str := fmt.Sprintf("Пользователь  %s добавлен!", name)
		dialog.ShowInformation("Успех", str, myWindow)
	}
}
