#---Не-Весь-Проэкт--Добавить--Если-Надо---#
all:
	go build -o shavel ./backend.go ./sql.go ./main.go  ./stack_message.go
	./shavel

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


clean:
	rm -rf shavel
	clear