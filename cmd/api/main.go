package main

import (
	"github.com/pinghoyk/budget-api/internal/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден!")
	}

}

// пока все пишу тут, потом надо разделить файлы и кинут в internal