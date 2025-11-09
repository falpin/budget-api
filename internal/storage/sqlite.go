// как подключиться к бд, соединение, инициализация схемы БД
package storage

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)
