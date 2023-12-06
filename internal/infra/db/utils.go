package db

import (
	"context"
	"database/sql"
	"log"
	"os"
)

func createDBFileName() string {
	fileName := "test_db.db"
	finalFileName := "./" + fileName
	path, err := os.Getwd()
	if err == nil {
		finalFileName = path + "/" + fileName
	} else {
		log.Print(err.Error())
	}
	return finalFileName
}

func createDb() *sql.DB {
	finalFileName := createDBFileName()
	db, _ := sql.Open("sqlite3", finalFileName)
	return db
}

func deleteDbTest() {
	finalFileName := createDBFileName()
	os.Remove(finalFileName)
}

func acquireConnection(ctx context.Context, db *sql.DB) (*sql.Conn, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Println("Error creating connection!")
		log.Println(err.Error())
		return nil, err
	}
	return conn, nil
}
