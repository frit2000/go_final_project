# Файлы для итогового задания

Финальный проект Я.Практикум Go разработчик с нуля.
Реализация бэкенда для планировщика задач.

По ТЗ сделаны задания со звездочкой:
В шаге 1: возможность определять извне порт при запуске сервера через переменную окружения TODO_PORT
В шаге 2: возможность определять путь к файлу базы данных через переменную окружения TODO_DBFILE
В шаге 3: назначение задач на указанные дни недели и на указанные дни месяца
В шаге 5: возможность выбрать задачи через строку поиска
В шаге 8: аутентификация и создание докер контейнера

Для запуска тестов в tests/settings.go необходимы следующие параметры
var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = true
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.gW0hFNtJo0_Q9z1n13dR1SCMfvFQltEh3z9weFL6YQ4`

Ключи для запуска программы:
	password (Пароль для приложения, в данном случае пароль должен быть 12345678, т.к. для него прописан токен в настройках тестов)
	port (Порт для запуска веб сервера, по умолчанию 7540)
	dbpath (Путь к базе данных, по умолчанию база создается в текущей директории)


Директория `db` содержит функцию проверки существоваяи БД и ее создания, а также саму БД
Директория `server` содержит функции по созданию сервера и функции для API
Директория `nextdate` содержит функции для вычисления дат, на которые переносятся задачи с параметром repeat
Директория `env` содержит функции и параметры для работы с переменными окружения
Директория `structs` содержит экспортируемые структуры и константы
В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.
Директория `web` содержит файлы фронтенда.
