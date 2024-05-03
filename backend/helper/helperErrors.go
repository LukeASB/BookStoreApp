package helper

import (
	"log"
	"net/http"
)

/*
IsHTTPStatusError checks if there is an error. If an error is present, it logs the error and sends
an HTTP error response with the specified status code to the client.
It returns true if there is an error, otherwise false.

Parameters:

	param1: w http.ResponseWriter
	param2: error
	param3: HTTP Status Code

Returns:

	return1: boolean
*/
func IsHTTPStatusError(w http.ResponseWriter, err error, statusCode int) bool {
	if err != nil {
		LogHTTPStatusError(w, err, statusCode)
		return true
	}

	return false
}

/*
LogHTTPStatusError logs the provided error and sends an HTTP error response with the specified
status code to the client.

Parameters:

	param1: w http.ResponseWriter
	param2: error
	param3: HTTP Status Code
*/
func LogHTTPStatusError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}

/*
HandleHTTPStatusError sends an HTTP error response with the specified status code to the client.

Parameters:

	param1: w http.ResponseWriter
	param3: HTTP Status Code
*/
func HandleHTTPStatusError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}
