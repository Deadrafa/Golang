package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_grades, Happines float64
	Hobbi                []string
}
type Сourse struct {
	Name_curse string
}

type FormData struct {
	FullName        string
	Email           string
	Password        string
	ConfirmPassword string
	Error           string
	//Nullstring      string
}

func home_pAge(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, "")
}
func registr_pAge(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Handling /registration/ request") // Отладочный вывод
	tmpl, err := template.ParseFiles("templates/registr_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := FormData{}

	if r.Method == http.MethodPost {
		// Обработка POST-запроса
		data.FullName = r.FormValue("fullName")
		data.Email = r.FormValue("email")
		data.Password = r.FormValue("password")
		data.ConfirmPassword = r.FormValue("confirmPassword")

		if check_user_in_db(data.FullName, data.Email) == 111 {
			data.Error = "Пользователь уже существует"
			err = tmpl.Execute(w, data)
			return
		} else if check_user_in_db(data.FullName, data.Email) == 222 {
			data.Error = "Email уже существует"
			err = tmpl.Execute(w, data)
			return
		}
		if data.Password != data.ConfirmPassword {
			data.Error = "Пароли не совпадают"
			err = tmpl.Execute(w, data)
			return
		}
		add_user_in_db(data.FullName, data.Email, data.Password)

		// Если все проверки пройдены, можно сохранить данные и т.д.
		http.Redirect(w, r, "/", http.StatusSeeOther)

		fmt.Printf("Name: %s\n", data.FullName)
		fmt.Printf("Email: %s\n", data.Email)
		fmt.Printf("Password: %s\n", data.Password)
		fmt.Printf("confirmPassword: %s\n", data.ConfirmPassword)
		//fmt.Fprintf(w, "Email: %s\nPassword: %s\nRemember me: %s\n", email, password, remember)
	} else {
		// Обработка GET-запроса и отображение HTML-страницы
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func contact_pAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("Py") == "Python" {
			http.Redirect(w, r, "/Python/", http.StatusSeeOther)
		}
		if r.FormValue("CSS") == "CSS" {
			http.Redirect(w, r, "/CSS/", http.StatusSeeOther)
		}
		if r.FormValue("Dizaner") == "Dizaner" {
			http.Redirect(w, r, "/Dizaner/", http.StatusSeeOther)
		}
		if r.FormValue("Game") == "Game" {
			http.Redirect(w, r, "/Game/", http.StatusSeeOther)
		}
		if r.FormValue("Data") == "Data" {
			http.Redirect(w, r, "/Data/", http.StatusSeeOther)
		}
		if r.FormValue("QA") == "QA" {
			http.Redirect(w, r, "/QA/", http.StatusSeeOther)
		}
		if r.FormValue("FullStack") == "FullStack" {
			http.Redirect(w, r, "/FullStack/", http.StatusSeeOther)
		}
		if r.FormValue("ML") == "ML" {
			http.Redirect(w, r, "/ML/", http.StatusSeeOther)
		}
		if r.FormValue("UNIX") == "UNIX" {
			http.Redirect(w, r, "/UNIX/", http.StatusSeeOther)
		}
	}

	tmpl, err := template.ParseFiles("templates/contact_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func python_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("Py") == "Py" {
			http.Redirect(w, r, "https://stepik.org/course/58852/promo?search=4610597764", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по Python"}
	tmpl, err := template.ParseFiles("templates/python_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func css_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("CSS") == "CSS" {
			http.Redirect(w, r, "https://stepik.org/course/123508/promo?search=4610633355", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по CSS и HTML"}
	tmpl, err := template.ParseFiles("templates/css_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func dizaner_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("Dizaner") == "Dizaner" {
			http.Redirect(w, r, "https://stepik.org/course/38218/promo?search=4610669788", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по Дизайнер"}
	tmpl, err := template.ParseFiles("templates/dizaner_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func game_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("Game") == "Game" {
			http.Redirect(w, r, "https://stepik.org/course/126291/promo?search=4610714222", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по GameDev"}
	tmpl, err := template.ParseFiles("templates/game_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func data_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("Data") == "Data" {
			http.Redirect(w, r, "https://stepik.org/course/73952/promo?search=4610723638", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по DataScience"}
	tmpl, err := template.ParseFiles("templates/data_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func testing_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("QA") == "QA" {
			http.Redirect(w, r, "https://stepik.org/course/118842/promo?search=4610733248", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по Тестировшик"}
	tmpl, err := template.ParseFiles("templates/testing_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func fullstack_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("FullStack") == "FullStack" {
			http.Redirect(w, r, "https://stepik.org/course/135466/promo?search=4610759328", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по FullStack разработчик"}
	tmpl, err := template.ParseFiles("templates/fullstack_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func ml_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("ML") == "ML" {
			http.Redirect(w, r, "https://stepik.org/course/4852/promo?search=4610769468", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по Машинное обучение"}
	tmpl, err := template.ParseFiles("templates/ml_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func unix_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if r.FormValue("UNIX") == "UNIX" {
			http.Redirect(w, r, "https://stepik.org/course/762/promo?search=4610775617", http.StatusSeeOther)
		}
	}
	curse := Сourse{"Бесплатный курс по Unix-системы"}
	tmpl, err := template.ParseFiles("templates/unix_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, curse)
}

func marathon_page(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Укажите путь к вашему файлу
		filePath := "static/css/App/RoboDendi.exe"

		// Устанавливаем заголовки для скачивания файла
		w.Header().Set("Content-Disposition", "attachment; filename=RoboDendi_Marathon.exe")
		w.Header().Set("Content-Type", "application/octet-stream")

		// Сервируем файл пользователю
		http.ServeFile(w, r, filePath)
	}
	tmpl, err := template.ParseFiles("templates/marathon_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, "")
}

func dendi_chat_page(w http.ResponseWriter, r *http.Request, st *Stack_message) {
	tmpl, err := template.ParseFiles("templates/dendi_chat_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Обработка отправки формы
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		message := r.FormValue("message")

		if username == "" || message == "" {
			st.Error = "Имя пользователя и сообщение не могут быть пустыми"
		} else if r.FormValue("Enter_massage") == "Enter_massage" {
			if st.size >= STACK_SIZE {
				st.Error = "Закончилось место"
				err = tmpl.Execute(w, st)
				return
			}
			fullMessage := fmt.Sprintf("<%s>: %s", username, message)
			push(st, fullMessage)
			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
			Print_stack(st)
		}
	}

	tmpl.Execute(w, st)
}

func HandleRequest(st *Stack_message) {
	http.HandleFunc("/", home_pAge)
	http.HandleFunc("/contact/", contact_pAge)
	http.HandleFunc("/registration/", registr_pAge)
	http.HandleFunc("/Python/", python_page)
	http.HandleFunc("/CSS/", css_page)
	http.HandleFunc("/Dizaner/", dizaner_page)
	http.HandleFunc("/Game/", game_page)
	http.HandleFunc("/Data/", data_page)
	http.HandleFunc("/QA/", testing_page)
	http.HandleFunc("/FullStack/", fullstack_page)
	http.HandleFunc("/ML/", ml_page)
	http.HandleFunc("/UNIX/", unix_page)
	http.HandleFunc("/Dendi_Marathon/", marathon_page)
	http.HandleFunc("/Dendi_Chat/", func(w http.ResponseWriter, r *http.Request) {
		dendi_chat_page(w, r, st)
	})
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
