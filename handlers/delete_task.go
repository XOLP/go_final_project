package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func DeleteTask(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	// Получаем значение параметра id из запроса

	idStr := r.URL.Query().Get("id")

	// Проверяем, что id не пустой
	if idStr == "" || idStr == "0" {

		erresponse := ErrorResponse{Error: "Не указан идентификатор"}
		createErrorResponse(w, http.StatusBadRequest, erresponse)
		return
	}

	// Преобразуем id в число
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		erresponse := ErrorResponse{Error: "Некорректный идентификатор"}
		createErrorResponse(w, http.StatusBadRequest, erresponse)
		return
	}

	// Удаляем задачу из базы данных
	deleteSQL := `DELETE FROM scheduler WHERE id = ?`
	result, err := db.Exec(deleteSQL, id)
	if err != nil {

		log.Printf("Error: %v\n", err)
		response := ErrorResponse{Error: "Ошибка при удалении задачи"}
		createErrorResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Проверяем количество удаленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error: %v\n", err)
		response := ErrorResponse{Error: "Ошибка при удалении задачи"}
		createErrorResponse(w, http.StatusInternalServerError, response)
		return
	}

	if rowsAffected == 0 {
		response := ErrorResponse{Error: "Задача не найдена"}
		createErrorResponse(w, http.StatusNotFound, response)
		return
	}

	// Успешный ответ
	createResponse(w, http.StatusOK, map[string]interface{}{})
}
