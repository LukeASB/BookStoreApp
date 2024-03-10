package view

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"readinglistapp/internal/data"
	"strconv"
	"strings"
)

const (
	BASEHTML   = "./ui/html/base.html"
	NAVHTML    = "./ui/html/partials/nav.html"
	HOMEHTML   = "./ui/html/pages/home.html"
	VIEWHTML   = "./ui/html/pages/view.html"
	CREATEHTML = "./ui/html/pages/create.html"
)

type Envelope map[string]any

/*
RenderJSON marshals the given Envelope data into JSON format with indentation and returns it as a byte slice.
Returns an error if the marshaling fails.

Parameters:

	param1: data Envelope - JSON envelope

Returns:

	return1: slice of bytes
	return2: error
*/
func RenderJSON(data Envelope) ([]byte, error) {
	jsonResponse, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return nil, err
	}

	jsonResponse = append(jsonResponse, '\n')

	return jsonResponse, nil
}

/*
ReadJSON decodes the JSON data from the request body r.Body into the provided destination interface dst,
which must be a pointer to the appropriate data structure.
It sets a maximum size limit for the request body to prevent denial of service attacks.
The function returns an error if there are any issues during the decoding process,
such as invalid JSON format or exceeding the maximum size limit.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request

Returns:

	return1: error
*/
func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(data); err != nil {
		return err
	}

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return err
	}

	return nil
}

/*
This function BookCreateForm renders the HTML form for creating a book.
It parses the HTML template files, including the base template, navigation template, and specific create book template.
If there are any errors during the parsing or execution of the templates, it returns the error.
Otherwise, it renders the form successfully.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request

Returns:

	return1: error
*/
func BookCreateForm(w http.ResponseWriter, r *http.Request) error {
	files := []string{BASEHTML, NAVHTML, CREATEHTML}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		return err
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		return err
	}

	return nil
}

/*
BookCreateProcess handles form data parsing, constructs a book struct, marshals it into JSON.
It returns the JSON data or error.

Parameters:

	param1: w http.ResponseWriter
	param2: r *http.Request

Returns:

	return1: slice of bytes
	return1: error
*/
func BookCreateProcess(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}

	title := r.PostForm.Get("title")

	published, err := strconv.Atoi(r.PostForm.Get("published"))

	if err != nil {
		return nil, err
	}

	pages, err := strconv.Atoi(r.PostForm.Get("pages"))

	if err != nil {
		return nil, err
	}

	genres := strings.Split(r.PostForm.Get("genres"), ",")

	rating, err := strconv.ParseFloat(r.PostForm.Get("rating"), 64)

	if err != nil {
		return nil, err
	}

	book := struct {
		Title     string   `json:"title"`
		Pages     int      `json:"pages,omitempty"`
		Published int      `json:"published,omitempty"`
		Genres    []string `json:"genres,omitempty"`
		Rating    float64  `json:"rating,omitempty"`
	}{
		Title:     title,
		Pages:     pages,
		Published: published,
		Genres:    genres,
		Rating:    rating,
	}

	data, err := json.Marshal(book)

	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Renders the book view page, parsing template files, executing the template with book data.
Returning any error encountered.

Parameters:

	param1: w http.ResponseWriter
	param2: id of the book
	param3: pointer of book data

Returns:

	return1: error
*/
func BookView(w http.ResponseWriter, id string, book *data.Book) error {
	files := []string{BASEHTML, NAVHTML, VIEWHTML}

	// Used to convert comma-separated genres to a slice within the template.
	funcs := template.FuncMap{"join": strings.Join}

	ts, err := template.New("showBook").Funcs(funcs).ParseFiles(files...)

	if err != nil {
		return err
	}

	err = ts.ExecuteTemplate(w, "base", book)

	if err != nil {
		return err
	}

	return nil
}

/*
Renders the home page with a list of books, parsing template files and executing the template with book data.
Returning any error encountered.

Parameters:

	param1: w http.ResponseWriter
	param2: pointer of a slice of book data

Returns:

	return1: error
*/
func BookHome(w http.ResponseWriter, books []*data.Book) error {
	files := []string{BASEHTML, NAVHTML, HOMEHTML}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	err = ts.ExecuteTemplate(w, "base", books)

	if err != nil {
		return err
	}

	return nil
}
