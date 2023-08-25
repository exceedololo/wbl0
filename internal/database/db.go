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
	DBUser       string
	DBPassword   string
	DBName       string
	DBSchemeName string
	DBHost       string
}

func NewDataBase(config DBconfig) (*DataBase, error) {
	//connecting to .env to get credentials
	/*err := godotenv.Load("/home/alex/GolandProjects/WB/bwTechLvl0/database/config.env")
	if err != nil {
		return nil, err
	}*/
	/*dbConfig := DBconfig{
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSchemeName: os.Getenv("DB_SCHEME_NAME"),
		DBHost:       os.Getenv("DB_HOST"),
	}*/
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable search_path=%s host=%s",
		config.DBUser, config.DBPassword, config.DBName, config.DBSchemeName, config.DBHost)
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
