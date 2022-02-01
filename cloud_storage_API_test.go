package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	//This creates a new request to the /healthcheck endpoint
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}
	//Creates a new recorder to record the response received by the entries endpoint.
	rr := httptest.NewRecorder()
	//Hits the endpoint with the recorder and request.
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(rr, req)
	//Checks if the response is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		//Sends an error tagging as a test failure
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Is an expected output from the endpoint.
	expected := `[{
		"message": "health check is healthy"
	}]`
	//Check if the response is equal to expected.
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
