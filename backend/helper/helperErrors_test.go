package helper

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIsHTTPStatusError(t *testing.T) {
	mockHTTPRes := httptest.NewRecorder()

	err := errors.New("💣")

	// Test when error is not nil
	if result := IsHTTPStatusError(mockHTTPRes, err, http.StatusInternalServerError); !result {
		t.Errorf("Expected true but got false")
	}

	// Check if the response status code is as expected
	if status := mockHTTPRes.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, status)
	}

	err = nil

	if result := IsHTTPStatusError(mockHTTPRes, err, http.StatusInternalServerError); result {
		t.Errorf("Expected false but got true")
	}
}

func TestLogHTTPStatusError(t *testing.T) {
	mockHTTPRes := httptest.NewRecorder()

	err := errors.New("💣")

	LogHTTPStatusError(mockHTTPRes, err, http.StatusInternalServerError)

	if status := mockHTTPRes.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, status)
	}

	// Check if the response body contains the status text
	expectedBody := http.StatusText(http.StatusInternalServerError)
	if body := mockHTTPRes.Body.String(); !strings.Contains(body, expectedBody) {
		t.Errorf("Expected response body to contain '%s' but got '%s'", expectedBody, body)
	}
}

func TestHandleHTTPStatusError(t *testing.T) {
	mockHTTPRes := httptest.NewRecorder()

	HandleHTTPStatusError(mockHTTPRes, http.StatusInternalServerError)

	if status := mockHTTPRes.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, status)
	}

	expectedBody := http.StatusText(http.StatusInternalServerError)

	if body := mockHTTPRes.Body.String(); !strings.Contains(body, expectedBody) {
		t.Errorf("Expected response body to contain '%s' but got '%s'", expectedBody, body)
	}
}
