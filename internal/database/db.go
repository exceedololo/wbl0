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
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewDataBase(config DBconfig) (*DataBase, error) {

	/*connStr := fmt.Sprintf("postgres://wbadmin:19azty56@localhost:5432/wborderdb")*/ /*,
	config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName*/
	//connecting to my database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	fmt.Println("DBHost:", config.DBHost)
	fmt.Println("DBPort:", config.DBPort)
	fmt.Println("DBUser:", config.DBUser)
	fmt.Println("DBPassword:", config.DBPassword)
	fmt.Println("DBName:", config.DBName)
	fmt.Println(connStr)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	//checking the connection
	err = conn.Ping(context.Background())
	if err != nil {
		conn.Close(context.Background())
		log.Fatal(err)
	}
	fmt.Println("The connection has been accomplished")
	return &DataBase{Conn: conn}, nil
}

/*
// Adding data into database

	func (db *DataBase) InsertData(orderUID string, dateCreated time.Time, jsonData []byte) error {
		_, err := db.Conn.Exec(context.Background(), `
	        INSERT INTO wborderscheme.orders(order_uid, date_created, data)
	        VALUES($1, $2, $3)
	    `, orderUID, dateCreated, jsonData)
		return err
	}

// getting info by ID

	func (db *DataBase) GetDataByID(orderUID string) ([]byte, error) {
		var jsonData []byte
		err := db.Conn.QueryRow(context.Background(), `
	        SELECT data FROM wborderscheme.orders WHERE order_uid = $1
	    `, orderUID).Scan(&jsonData)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	}
*/
func (db *DataBase) Close() error {
	return db.Conn.Close(context.Background())
}

//vinesti v repositories
