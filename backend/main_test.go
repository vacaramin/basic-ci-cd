package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersHandler(t *testing.T) {
	// Initialize the database connection.
	db := initDB()
	defer db.Close()

	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsersHandler(db))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Here you might want to add more checks, like ensuring
	// the response body is what you expect.
}
