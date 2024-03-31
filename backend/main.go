package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mattn/go-sqlite3"
)

type Users struct {
	Name string
	Age  int
}

func main() {
	log.Println("...Starting server...")
	var _ *sqlite3.SQLiteDriver
	loadENV()
	log.Println("...Environment variables loaded...")
	db := initDB()
	defer db.Close()
	log.Println("...Initializing DB...")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", getUsersHandler(db))
	http.HandleFunc("/user", createUserHandler(db))
	http.HandleFunc("/update", updateUserHandler(db))
	http.HandleFunc("/delete", deleteUserHandler(db))
	http.HandleFunc("/404", NotFoundHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pong", pongHandler)
	http.HandleFunc("/ci", handleci)

	log.Println("Server is running on port:", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./users.db")

	if err != nil {
		log.Fatal(err)
	}
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating users table: ", err)
	}

	// Insert sample data
	insertDataSQL := `
	INSERT INTO users (name, age) VALUES
	('John Doe', 30),
	('Alice Smith', 25),
	('Bob Johnson', 40);
	`

	_, err = db.Exec(insertDataSQL)
	if err != nil {
		log.Fatal("Error inserting sample data: ", err)
	}

	return db
}

func handleci(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any domain
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests for CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	randomData := make(map[string]string)
	randomData["CI"] = "Github Actions"
	randomData["Deployment"] = "DockerHUB and azure"
	randomData["PipelineTest"] = "Dev"

	jsonResponse, err := json.Marshal(randomData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

// Handlers for CRUD operations
func getUsersHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any domain
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests for CORS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		users := []User{}
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		jsonResponse, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func createUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userID, err := result.LastInsertId()
		user.ID = int(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}
func updateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func deleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM users WHERE id = ?", user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// User struct
type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]string{"hello": "world"}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
func pingHandler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]string{
		"Message": "Pong",
	}
	jsonResponse, err := json.MarshalIndent(response, "", "2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func pongHandler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]string{
		"Message": "Ping",
	}
	jsonResponse, err := json.MarshalIndent(response, "", "2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"error": "404- Not found"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}
