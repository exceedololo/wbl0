package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

//tolko soedinenie

type DataBase struct {
	Conn *sql.DB
}

//some info for database

type DBconfig struct {
	DBUser       string //:= os.Getenv("wbadmin")
	DBPassword   string //:= os.Getenv("19az%&ty56")
	DBName       string //:= os.Getenv("wborderbase")
	DBSchemeName string //:= "wborderscheme"
	//string for connection - easy to modify
	//connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable search_path=%s", dbUser, dbPassword, dbName, schemeName)
}

func NewDataBase(config DBconfig) (*DataBase, error) {
	//connecting to .env to get credentials
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dbConfig := DBconfig{
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSchemeName: os.Getenv("DB_SHEME_NAME"),
	}
	connStr := fmt.Sprintf("user=%s passwprd=%s dbname=%s sslmode=disable search_path=%s",
		dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBSchemeName)

	//connecting to my database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Creating a context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//checking the connection
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The connection has been accomplished")
	return &DataBase{Conn: conn}, nil
}

func (db *DataBase) Close() error {
	return db.Conn.Close()
}

//vinesti v repositories
