package main

import (
	"final/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var webDir = "./web"
var install bool

func main() {

	serverStart()
}

func serverStart() {

	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	if err != nil {
		install = true
	}
	db, err := sqlx.Connect("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if install {
		createTableSQL := `CREATE TABLE IF NOT EXISTS scheduler (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        date TEXT NOT NULL,
        title TEXT NOT NULL,
        comment TEXT,
        repeat TEXT CHECK (length(repeat) <= 128)
        );`

		_, err = db.Exec(createTableSQL)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("CREATE INDEX ID_Date ON scheduler (date);")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("База данных успешно создана")
	} else {
		log.Println("База данных уже существует")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "7540"
	}

	fileServ := http.FileServer(http.Dir(webDir))

	http.HandleFunc("/api/nextdate", handlers.NextDate)
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasks(w, r, db)
	})
	http.HandleFunc("/api/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			handlers.GetTask(w, r, db)
		case http.MethodPost:
			handlers.AddTask(w, r, db)
		case http.MethodPut:
			handlers.UpdateTask(w, r, db)
		case http.MethodDelete:
			handlers.DeleteTask(w, r, db)

		default:
			http.Error(w, "Error method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/task/done", func(w http.ResponseWriter, r *http.Request) {
		handlers.TaskAsDone(w, r, db)
	})

	http.Handle("/", fileServ)

	log.Printf("Сервер запущен на порту " + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
