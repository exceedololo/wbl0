package main

import (
	"bwTechLvl0/internal/database"
	"bwTechLvl0/internal/models"
	"bwTechLvl0/internal/repositories"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"time"
)

// global variable used for caching data
/*var dataCache map[string]models.Order
var cacheMutex sync.RWMutex*/

func main() {
	//dbUser := os.Getenv("DB_USER")

	err := godotenv.Load("C:/Projects/Go/wbl0/internal/database/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("Current working directory:", currentDir)

	ctx := context.Background()

	//
	//host.docker.internal
	config := database.DBconfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	host := "host.docker.internal"
	port := "5432"

	timeout := time.Second * 2
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connection successful!")

	//creating an example of database
	time.Sleep(5 * time.Second)
	db, err := database.NewDataBase(config)
	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	//creating repositories
	repo, err := repositories.NewOrderRepo( /*ctx,*/ db)
	if err != nil {
		fmt.Println("Error creating repository:", err)
		return
	}

	//example of using methods of "repositories"
	order := models.Order{
		OrderUID:    "order123",
		DateCreated: time.Now(),
		Data:        []byte(`{"key": "value"}`),
	}
	err = repo.Upsert(ctx, models.Order{})
	if err != nil {
		fmt.Println("Error upserting order:", err)
		return
	}

	foundOrder, err := repo.GetById(ctx, order.OrderUID)
	if err != nil {
		fmt.Println("Error getting order by ID:", err)
		return
	}
	fmt.Println("Found order:", foundOrder)
	/*
		// Пример добавления записи в базу данных
		orderToAdd := models.Order{
			OrderUID:    "your_order_uid",
			DateCreated: time.Now(),
			Data:        []byte(`{"key": "value"}`),
		}
		err = repo.Upsert(context.Background(), orderToAdd)
		if err != nil {
			log.Fatal(err)
		}

		// Пример получения записи из базы данных по идентификатору
		orderToRetrieve := models.Order{
			OrderUID: "your_order_uid",
		}
		retrievedOrder, err := repo.GetById(context.Background(), orderToRetrieve)
		if err != nil {
			log.Fatal(err)
		}
		if retrievedOrder == nil {
			fmt.Println("Order not found")
		} else {
			fmt.Printf("Retrieved Order: %+v\n", retrievedOrder)
		}*/

	//starting database and NATS-stream
	//db, sc := initialize()

	//restoring cache from database
	//restoreCacheFromDB(db)

	//function that subscribes to NATS and is processing messages
	//subscribeToNATS(sc, db)

	//starting HTTP-server
	//startHTTPServer()

}

//

// name explains

// name explains
/*func restoreCacheFromDB(db *sql.DB) {
	rows, err := db.Query("SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard FROM orders")
	if err != nil {
		log.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale,
			&order.InternalSignature, &order.CustomerID, &order.DeliveryService,
			&order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		dataCache[order.OrderUID] = order
	}
}*/

//func initialize() (*sql.DB, stan.Conn) {
//opening JSON-file
//before next // useless
/*file, err := os.Open("model.json")
if err != nil {
	log.Fatal(err)
}
defer file.Close()*/

//getting data from JSON-file
/*var OrderData Order
decoder := json.NewDecoder(file)
err = decoder.Decode(&OrderData)
if err != nil {
	log.Fatal(err)
}*/
//some info for database
/*dbUser := os.Getenv("wbadmin")
	dbPassword := os.Getenv("19az%&ty56")
	dbName := os.Getenv("wborderbase")
	schemeName := "wborderscheme"
	//string for connection - easy to modify
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable search_path=%s", dbUser, dbPassword, dbName, schemeName)

	//connecting to my database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//checking the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The connection has been accomplished")

	//connecting to NATS-streaming server
	clusterID := "test-cluster"
	clientID := "client-1"
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	return db, sc
}

func subscribeToNATS(sc stan.Conn, db *sql.DB) {
	//subscribing to the NATS-channel
	channel := "order-updates"
	//making a handler func
	_, err := sc.Subscribe(channel, func(msg *stan.Msg) {
		var orderData Order
		err := json.Unmarshal(msg.Data, &orderData)
		if err != nil {
			log.Println("Error decoding message:", err)
			return
		}
		//adding data into cash and writing data into database
		cacheMutex.Lock()
		dataCache[orderData.OrderUID] = orderData
		cacheMutex.Unlock()

		err = insertOrderIntoDB(db, orderData)
		if err != nil {
			log.Println("Error inserting into database:", err)
			return
		}
		log.Printf("Receiver order: %s", orderData.OrderUID)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func getDataFromCache(orderUID string) (Order, bool) {
	cacheMutex.RLock()
	data, exists := dataCache[orderUID]
	cacheMutex.RUnlock()
	return data, exists

}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderUID := r.URL.Query().Get("orderUID")
	data, exists := getDataFromCache(orderUID)
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func startHTTPServer() {
	//initialisation of HTTP-server to handle data from cash
	http.HandleFunc("/getOrder", getOrderHandler)
	port := "8080"
	log.Printf("starting HTTP server on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}*/
