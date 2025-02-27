

# Файловая структура проекта
.
├── tests/ # Интеграционные и модульные тесты
├── web/ # Статические файлы фронтенда
└── handlers/ #


## Описание проекта

Веб-сервер для управления задачами с функциями:
- Создание/редактирование/улаление задач
- Поддержка периодических задач
- Автоматический пересчет даты выполнения
- REST API для интеграции с фронтендом


---

## Инструкция по запуску

### Требования
- Go 1.16+
- SQLite3

### Сборка и запуск

1. Компиляция для Windows:
```bash
go build -o scheduler.exe *.go
Запуск сервера:



./scheduler.exe
#Тестирование
Запуск всех тестов


go test ./tests 
#Индивидуальные тесты
Тест	Команда
Проверка работы API	go test -run ^TestApp$ ./tests 
Операции с БД	go test -run ^TestDB$ ./tests 
Расчет следующих дат	go test -run ^TestNextDate$ ./tests 
Добавление задач	go test -run ^TestAddTask$ ./tests 
Фильтрация задач	go test -run ^TestTasks$ ./tests 
Целостность данных	go test -run ^TestTask$ ./tests 
Редактирование задач	go test -run ^TestEditTask$ ./tests 
Завершение задач	go test -run ^TestDone$ ./tests
Удаление задач	go test -run ^TestDelTask$ ./tests 
