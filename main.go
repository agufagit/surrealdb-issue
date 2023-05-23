package main

import (
	"log"

	"github.com/surrealdb/surrealdb.go"
)

func main() {
	url := "ws://localhost:8000/rpc"
	conn, err := surrealdb.New(url)
	if err != nil {
		log.Fatalf("Failed to connect to database server: %+v", err)
	}

	_, err = conn.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	})
	if err != nil {
		log.Fatalf("Failed to signin to database server: %+v", err)
	}

	_, err = conn.Query("DEFINE NAMESPACE test;USE NS test;DEFINE DATABASE test;USE DB test;", nil)
	if err != nil {
		log.Fatalf("Failed to create namespace and database specified: %+v", err)
	}

	_, err = conn.Use("test", "test")
	if err != nil {
		log.Fatalf("Failed to use namespace and database specified: %+v", err)
	}

	param := map[string]any{}

	returnData, err := conn.Query("select * from ->contact->profile;", param)
	if err != nil {
		log.Fatalf("Failed to query database: %+v", err)
	}

	log.Printf("Query result: %+v", returnData)
}
