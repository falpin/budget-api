// описание структур базы данных
package storage

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID           int64     `db:"id"            json:"id"`
	Email        string    `db:"email"         json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	FirstName    string    `db:"first_name"    json:"first_name"`
	CreatedAt    time.Time `db:"created_at"    json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"    json:"updated_at"`
}
}
