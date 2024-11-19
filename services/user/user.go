package user

import (
	"database/sql"
	"e-book_webapp/utils"
	"time"
)

type User struct {
	Id        string    `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func CreateUser(db *sql.DB, user User) (int, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	query := "INSERT INTO users (name , email , password , created_at) VALUES (?,?,?,?)"
	result, err := db.Exec(query, user.Username, user.Email, hashedPassword, time.Now())
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(userID), err

}
