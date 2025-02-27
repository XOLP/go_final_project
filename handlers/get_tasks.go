package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

const tasksLimit = 50

func GetTasks(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {

	log.Println("Задачи получены GetTasks")
	// Лимит на количество возвращаемых задач

	// Получаем текущую дату
	now := time.Now().Format("20060102")

	// Запрос к базе данных для получения задач
	tasks := []Task{}
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE date >= ? ORDER BY date ASC LIMIT ?`
	err := db.Select(&tasks, query, now, tasksLimit)
	if err != nil {
		erresponse := ErrorResponse{Error: "Ошибка запроса к базе данных"}
		createErrorResponse(w, http.StatusInternalServerError, erresponse)
		log.Println("Ошибка запроса:", err)
		return
	}

	// Собираем ответ
	response := TasksResponse{
		Tasks: tasks,
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Отправляем Json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		erresponse := ErrorResponse{Error: "Ошибка кодирования ответа"}
		createErrorResponse(w, http.StatusInternalServerError, erresponse)
		log.Println("Error ", err)
	}
}
