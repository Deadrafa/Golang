#---Не-Весь-Проэкт--Добавить--Если-Надо---#
all:
	go build -o shavel ./backend.go ./sql.go ./main.go  ./stack_message.go 
	./shavel &
	sleep 2
	xdg-open http://localhost:8080

#--Запуск--Базы--Данных---#
db_start:
	sudo systemctl start postgresql

#--Остановка--Базы--Данных---#
db_stop:
	sudo systemctl stop postgresql

#--Чекнуть--Статус--Базы--Данных---#
db_status:
	sudo systemctl status postgresql

#--Это--Чтобы--С--Файла--Запросы--Кидать---#
db_request:
	psql -U postgres registration < sql_request_files/request_in_back.sql

#--Чистит--Базу--Данных--от--всех--записей--При--этом--Сама--Не--удаляется---#
db_clear:
	psql -U postgres registration < sql_request_files/clear.sql

#--Билдит--Проект--в--Docker--Контейнер---#
docker_build:
	docker build -t shavel .

#--Запускает--Собранный--Docker--Образ---#
docker_run: 
#	docker run -p 8080:8080 shavel
#	sleep 2
	docker compose up --build

#--Если--Занят--Процессом---Удалит--Его---#
server_stop:
	lsof -i :8080

clean:
	rm -rf shavel
#---ОСТОРОЖНО--УДАЛЯЕТ--ВСЕ--ОБРАЗЫ--DOCKER--#
	docker system prune -a
#---ОСТОРОЖНО--УДАЛЯЕТ--ВСЕ--ОБРАЗЫ--DOCKER--#
	clear