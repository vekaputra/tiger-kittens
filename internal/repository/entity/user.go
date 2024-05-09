package entity

import "time"

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
