package controller

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	helper "readinglistapp/helper"
	"readinglistapp/initialisers"
	"readinglistapp/model"
	"readinglistapp/view"

	"github.com/gorilla/mux"
)

/*
HealthCheck handles the health check endpoint.
It responds with the current environment, and version.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
*/
func HealthCheck(w http.ResponseWriter, r *http.Request, v view.IViewFuncs) {
	if r.Method != http.MethodGet {
		helper.HandleHTTPStatusError(w, http.StatusMethodNotAllowed)
		return
	}

	res := model.ResponseHealthCheck{
		Endpoint:    "Health Check Endpoint",
		Environment: os.Getenv("ENV"),
		Version:     os.Getenv("VERSION"),
	}

	jsonResponse, err := v.RenderJSON(view.Envelope{"healthcheck": res})

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	writeJSONResponse(w, http.StatusOK, jsonResponse, nil)
}

/*
Home displays the home page of the application.
It retrieves all books from the model and renders them using the view.BookHome function.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
*/
func Home(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	books, err := m.GetAll(bookCollection)
	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	err = v.BookHome(w, books)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}
}

/*
BookView displays details of a book.
It retrieves the book details using the provided book ID from the model,
then renders the book details using the view.BookView function.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
*/
func BookView(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		http.NotFound(w, r)
		return
	}

	book, err := m.Get(bookCollection, id)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	err = v.BookView(w, id, book)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}
}

/*
BookCreate handles book creation requests.
It routes GET requests to view.BookCreateForm and POST requests to view.BookCreateProcess and bookCreateProcess.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
*/
func BookCreate(w http.ResponseWriter, r *http.Request, v view.IViewFuncs) {
	switch r.Method {
	case http.MethodGet:
		err := v.BookCreateForm(w, r)

		if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
			return
		}

	case http.MethodPost:
		data, err := v.BookCreateProcess(w, r)

		if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
			return
		}

		bookCreateProcess(w, r, data)
	default:
		helper.HandleHTTPStatusError(w, http.StatusMethodNotAllowed)
	}
}

/*
bookCreateProcess sends a POST request to create a book.
It constructs an HTTP request with the provided data, sends it to the specified endpoint, and handles the response.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
	param3: data []byte The JSON data to be sent in the request body.
*/
func bookCreateProcess(w http.ResponseWriter, r *http.Request, data []byte) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s:%s/v1/books", os.Getenv("SITE_URL"), os.Getenv("PORT")), bytes.NewBuffer(data))

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		log.Printf("unexpected status: %s", response.Status)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/*
GetBooksHandler retrieves all books.
It fetches books from the model, renders them as JSON, and sends an HTTP response.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
*/
func GetBooksHandler(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	fmt.Println("GetBooksHandler")
	books, err := m.GetAll(bookCollection)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	jsonResponse, err := v.RenderJSON(view.Envelope{"books": books})

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	writeJSONResponse(w, http.StatusOK, jsonResponse, nil)
}

/*
The CreateBooksHandler function handles the creation of books.
It reads JSON input from the request, inserts the book into the model,
and returns a JSON response with appropriate status codes and headers.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request
*/
func CreateBooksHandler(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	var input model.Input

	err := v.ReadJSON(w, r, &input)

	if helper.IsHTTPStatusError(w, err, http.StatusBadRequest) {
		return
	}

	id, book, err := m.Insert(bookCollection, input)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/books/%s", id))

	jsonResponse, err := v.RenderJSON(view.Envelope{"book": book})

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	writeJSONResponse(w, http.StatusCreated, jsonResponse, headers)
}

/*
GetBook retrieves the details of a book identified by its ID from the model layer.
If the book is found, it returns a JSON response containing the book details.
Otherwise, it logs the error and returns an appropriate HTTP status code.

Parameters:

	param1: http.ResponseWriter
	param2: *http.Request
*/
func GetBook(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	var err error
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		helper.HandleHTTPStatusError(w, http.StatusBadRequest)
		return
	}

	book, err := m.Get(bookCollection, id)

	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			helper.LogHTTPStatusError(w, err, http.StatusNotFound)
		default:
			helper.LogHTTPStatusError(w, err, http.StatusInternalServerError)
		}
		return
	}

	jsonResponse, err := v.RenderJSON(view.Envelope{"book": book})

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	writeJSONResponse(w, http.StatusOK, jsonResponse, nil)
}

/*
UpdateBook handles the updating of a book identified by its ID.
It retrieves the ID from the request parameters, fetches the existing book from the model layer,
updates its fields based on the provided JSON input, and returns a JSON response with the updated
book details or an appropriate error status.

Parameters:

	param1: http.ResponseWriter
	param2: *http.Request
*/
func UpdateBook(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	var err error
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		helper.HandleHTTPStatusError(w, http.StatusBadRequest)
		return
	}

	book, err := m.Get(bookCollection, id)

	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			helper.HandleHTTPStatusError(w, http.StatusNotFound)
		default:
			helper.HandleHTTPStatusError(w, http.StatusInternalServerError)
		}
		return
	}

	var input struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float64 `json:"rating"`
	}

	err = v.ReadJSON(w, r, &input)

	if helper.IsHTTPStatusError(w, err, http.StatusBadRequest) {
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.Published != nil {
		book.Published = *input.Published
	}

	if input.Pages != nil {
		book.Pages = *input.Pages
	}

	if len(input.Genres) > 0 {
		book.Genres = input.Genres
	}

	if input.Rating != nil {
		book.Rating = *input.Rating
	}

	err = m.Update(bookCollection, id, book)

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	jsonResponse, err := v.RenderJSON(view.Envelope{"book": book})

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	writeJSONResponse(w, http.StatusOK, jsonResponse, nil)
}

/*
DeleteBook handles the deletion of a book identified by its ID.
It retrieves the ID from the request parameters,
attempts to delete the corresponding record in the model layer, and returns an appropriate JSON response
with the status code indicating success or failure.

Parameters:

	param1: http.ResponseWriter
	param2: *http.Request
*/
func DeleteBook(w http.ResponseWriter, r *http.Request, v view.IViewFuncs, m model.IModelFuncs, bookCollection initialisers.IBookCollection) {
	var err error
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		helper.HandleHTTPStatusError(w, http.StatusBadRequest)
		return
	}

	err = m.Delete(bookCollection, id)

	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			if helper.IsHTTPStatusError(w, err, http.StatusNotFound) {
				return
			}
		default:
			if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
				return
			}
		}
		return
	}

	jsonResponse, err := v.RenderJSON(view.Envelope{"message": "book successfully deleted"})

	if helper.IsHTTPStatusError(w, err, http.StatusInternalServerError) {
		return
	}

	writeJSONResponse(w, http.StatusOK, jsonResponse, nil)
}

/*
writeJSONResponse writes a JSON response to the provided http.ResponseWriter with the specified status code,
JSON content, and headers.
It sets the provided headers, specifies the content type as JSON, writes the status code,
and writes the JSON response body.

Parameters:

	param1: w http.ResponseWriter - the http.ResponseWriter to write the response to
	param2: status int - the HTTP status code to set in the response
	param3: jsonResponse []byte - the JSON content to include in the response body
	param4: headers http.Header - the headers to include in the response
*/
func writeJSONResponse(w http.ResponseWriter, status int, jsonResponse []byte, headers http.Header) {
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
