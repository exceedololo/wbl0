package database

import (
	"context"

	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

//tolko soedinenie

type DataBase struct {
	Conn *pgx.Conn
}

//some info for database

type DBconfig struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSchemeName string
}

func NewDataBase(config DBconfig) (*DataBase, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		config.DBHost, config.DBPort,
		config.DBUser, config.DBPassword, config.DBName, config.DBSchemeName)
	//connecting to my database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	//checking the connection
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The connection has been accomplished")
	return &DataBase{Conn: conn}, nil
}

func (db *DataBase) Close() error {
	return db.Conn.Close(context.Background())
}

//vinesti v repositories
