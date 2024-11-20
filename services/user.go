package services

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
	query := "INSERT INTO user (username , email , password ) VALUES (?,?,?)"
	result, err := db.Exec(query, user.Username, user.Email, hashedPassword)
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(userID), err

}

func UpdateUser(db *sql.DB, user User) (int, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	query := "UPDATE user  set username = ? , email = ? , password = ? where ID = ?"
	result, err := db.Exec(query, user.Username, user.Email, hashedPassword, user.Id)
	if err != nil {
		return 0, err
	}
	userID, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(userID), err
}

func GetUser(db *sql.DB, id int) (User, error) {
	user := User{}
	var createdAt []byte
	query := "SELECT * FROM user where id = ?"
	err := db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &createdAt)
	if err != nil {
		return user, err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return user, err
	}
	user.CreatedAt = parsedTime
	return user, err

}
