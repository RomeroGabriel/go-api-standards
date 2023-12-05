package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/RomeroGabriel/go-api-standards/internal/entity"
)

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS User 
		(Id VARCHAR(255) NOT NULL PRIMARY KEY, Name VARCHAR(255), Email VARCHAR(255), Password VARCHAR(255));
	`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *entity.User) error {
	ctx := context.Background()
	conn, err := u.DB.Conn(ctx)
	if err != nil {
		log.Println("Error on func Create from User")
		return err
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "INSERT INTO User (Id, Name, Email, Password) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	ctx := context.Background()
	conn, err := u.DB.Conn(ctx)
	if err != nil {
		log.Println("Error on func Create from User")
		return nil, err
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "SELECT Id, Name, Email, Password FROM User WHERE Email = ?")
	if err != nil {
		return nil, err
	}
	var result entity.User
	err = stmt.QueryRowContext(ctx, email).Scan(&result.ID, &result.Name, &result.Email, &result.Password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
