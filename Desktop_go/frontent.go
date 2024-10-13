package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIGHT        float32 = 1080
	HEIGHT       float32 = 700
	DELETE_IN_DB int     = 1
	ADD_IN_DB    int     = 2
)

type User_info struct {
	db_name  string
	db_user  string
	username string
	trigger  int
	email    string
	password string
}

func Frontent(information User_info) {

	myApp := app.New()
	myWindow := myApp.NewWindow("Server Support")

	image := canvas.NewImageFromFile("/home/daniel/Desktop/Desktop_go/back.jpg")
	image.FillMode = canvas.ImageFillContain

	db_name, user_db, password := widget.NewEntry(), widget.NewEntry(), widget.NewEntry()

	table := createTable([][]string{})

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Пользователь", Widget: user_db},
			{Text: "Имя Базы", Widget: db_name},
			{Text: "Пароль", Widget: password},
		},
		OnSubmit: func() { // optional, handle form submission
			information.db_user, information.db_name = user_db.Text, db_name.Text
			sql_connect(information, myWindow)
			data := check_user_in_db(information)
			updateTable(table, data)
			log.Println("Хост:", user_db.Text)
			log.Println("Логин:", db_name.Text)
			log.Println("Пароль:", password.Text)

			information.db_name = db_name.Text
			information.db_user = user_db.Text

			user_db.SetText("")
			db_name.SetText("")
			password.SetText("")

			// myWindow.Close()
		},
		OnCancel: func() {
			user_db.SetText("")
			db_name.SetText("")
			password.SetText("")
			// myWindow.Close()
		},
	}

	user_del := widget.NewEntry()
	form_delete_user := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Имя пользователя", Widget: user_del},
		},
		OnSubmit: func() { // optional, handle form submission
			information.username = user_del.Text
			information.trigger = DELETE_IN_DB
			tr := editing_db(information)
			data := check_user_in_db(information)
			err_sql(user_del.Text, tr, myWindow)
			updateTable(table, data)
			log.Println("Имя:", user_del.Text)

			user_del.SetText("")

		},
		OnCancel: func() {
			user_del.SetText("")
		},
	}
	user_add, email_add, passord_add := widget.NewEntry(), widget.NewEntry(), widget.NewEntry()
	form_add_user := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Имя пользователя", Widget: user_add},
			{Text: "Email пользователя", Widget: email_add},
			{Text: "Пароль пользователя", Widget: passord_add},
		},
		OnSubmit: func() { // optional, handle form submission
			information.trigger = ADD_IN_DB
			information.username, information.email, information.password = user_add.Text, email_add.Text, passord_add.Text
			tr := editing_db(information)
			data := check_user_in_db(information)
			err_sql(user_add.Text, tr, myWindow)
			updateTable(table, data)
			log.Println("Имя:", user_add.Text)
			log.Println("Email:", email_add.Text)
			log.Println("Password:", passord_add.Text)

			user_add.SetText("")
			email_add.SetText("")
			passord_add.SetText("")

		},
		OnCancel: func() {
			user_add.SetText("")
			email_add.SetText("")
			passord_add.SetText("")
		},
	}
	scrollableTable := container.NewScroll(table)
	content := container.NewBorder(nil, nil, nil, nil, scrollableTable)

	left := container.NewAppTabs(
		container.NewTabItem("Подключение к базе", form),
		container.NewTabItem("Добавление Пользователя", form_add_user),
		container.NewTabItem("Удаление Пользователя", form_delete_user),
		container.NewTabItem("Список Пользователей", content),
	)
	left.SetTabLocation(container.TabLocationLeading)

	tabs := container.NewAppTabs(
		container.NewTabItem("Список Пользователей", container.NewStack(image)),
		container.NewTabItem("База данных", left),
	)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(WIGHT, HEIGHT))
	myWindow.ShowAndRun()
}
