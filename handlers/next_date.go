package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const maxDays = 400

func NextDate(w http.ResponseWriter, r *http.Request) {
	nowStr := r.URL.Query().Get("now")
	dateStr := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	if nowStr == "" || dateStr == "" || repeat == "" {
		http.Error(w, "empty params", http.StatusBadRequest)
		return
	}

	now, err := time.Parse("20060102", nowStr)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	nextDate, err := nextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, nextDate)
}

func nextDate(now time.Time, date string, repeat string) (string, error) {
	if date == "" {
		return "", errors.New("date empty")
	}
	// Узнаем исходную дату
	startDate, err := time.Parse("20060102", date)
	if err != nil {
		return "", errors.New("error date format")
	}

	// Проверяем повторения
	if repeat == "" {
		return "", errors.New("repeat rule empty")
	}

	repeatParts := strings.Fields(repeat)
	if len(repeatParts) < 1 {
		return "", errors.New("error repeat rule")
	}

	repeatType := repeatParts[0]

	var days int
	var nextDate time.Time
	switch repeatType {
	case "y":
		if len(repeatParts) != 1 {
			return "", errors.New("Error year repeat rule")
		}
		nextDate = startDate.AddDate(1, 0, 0)
	case "d":
		if len(repeatParts) != 2 {
			return "", errors.New("Error day repeat rule")
		}

		days, err = strconv.Atoi(repeatParts[1])

		if err != nil || days <= 0 || days > maxDays {
			return "", errors.New("Error day interval")
		}
		nextDate = startDate.AddDate(0, 0, days)

	default:
		return "", errors.New("Error repeat rule")
	}

	if time.Now().Format("20060102") == startDate.Format("20060102") {
		return time.Now().Format("20060102"), nil
	}

	for !nextDate.After(now) {
		switch repeatType {
		case "y":
			nextDate = nextDate.AddDate(1, 0, 0)
		case "d":
			nextDate = nextDate.AddDate(0, 0, days)

		}
	}

	return nextDate.Format("20060102"), nil
}
