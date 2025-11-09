package main

import (
	"github.com/pinghoyk/budget-api/internal/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден!")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "budget-data/budget.db"
		log.Printf("DB_PATH не задан - используем значение по умолчанию: %s", dbPath)
	}
}

// пока все пишу тут, потом надо разделить файлы и кинут в internal