package main

import (
	"context"
	"encoding/json"
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
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

type URL struct {
	Url string `json:"url"`
}

func getURL(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	if r.Method == "GET" {
		var url URL
		var path string
		path = r.URL.Path[1:]

		err := conn.QueryRow(context.Background(), "SELECT url FROM uniform_resource_locator WHERE path = $1", path).Scan(&url.Url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Error querying database: %v\n", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json, err := json.Marshal(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Error marshalling JSON: %v\n", err)
			return
		}

		w.Write(json)
		fmt.Printf("Successfully got url: %v\n", url)
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
