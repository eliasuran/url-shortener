package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	loadENV()
	conn := connect()
	defer conn.Close(context.Background())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getURL(w, r, conn)
	})
	http.HandleFunc("/newURL", func(w http.ResponseWriter, r *http.Request) {
		setURL(w, r, conn)
	})

	corsHandler := cors.Default().Handler(http.DefaultServeMux)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connect() *pgx.Conn {
	fmt.Println(os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Unable to connect to database")
		os.Exit(1)
	}
	return conn
}

func getURL(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	if r.Method == "GET" {
		var path string
		err := conn.QueryRow(context.Background(), "SELECT path FROM uniform_resource_locator").Scan(&path)
		if err != nil {
			fmt.Fprintf(w, "Error querying database: %v", err)
			return
		}
		fmt.Fprintf(w, "Found: %s", path)
		return
	}
	fmt.Println("Only GET is allowed")
	return
}

func setURL(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	if r.Method == "POST" {
		fmt.Println(r.URL)
		fmt.Fprintf(w, "Hello World!")
		return
	}
	fmt.Println("Only POST is allowed")
	return
}
