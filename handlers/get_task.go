package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func GetTask(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	// Получаем значение параметра id из запроса
	id := r.URL.Query().Get("id")

	// Проверяем, что id не пустой
	if id == "" {
		erresponse := ErrorResponse{Error: "Не указан идентификатор"}
		createErrorResponse(w, http.StatusBadRequest, erresponse)
		return
	}

	// Запрос к базе данных для получения задачи по идентификатору
	var task Task
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	err := db.Get(&task, query, id)
	if err != nil {
		log.Println("Error:", err)
		erresponse := ErrorResponse{Error: "Задача не найдена"}
		createErrorResponse(w, http.StatusNotFound, erresponse)
		return
	}

	// Формируем ответ
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println("Error:", err)
		erresponse := ErrorResponse{Error: "Ошибка кодирования ответа"}
		createErrorResponse(w, http.StatusInternalServerError, erresponse)
		return
	}
}
