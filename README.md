# BookStoreApp
Full-Stack Book Store App using MVC design pattern with Golang, MongoDB with a basic UI.

## Description
The Basic Book Store App is a web application that displays book recommendations. Users can browse through existing book lists and create their own recommendations. The app simplifies the process of sharing top book picks with friends and family, eliminating the need to respond to individual requests.

## Installation
Pre-req: MongoDB Set Up.
To run the Book Store App App locally, follow these steps:

1. Clone the repository:
```
git clone <url>
```

2. Navigate to the project directory:
```
cd BookStoreApp
```

3. Configure Env:
- Add a database connection string to MongoDB to the `.env` file under `DB_URL`.
- Specify the desired port in the `.env` file, e.g :80

4. Run the application:
```
go run .
```


5. Access the application in your web browser at [http://localhost{:port}](http://localhost{:port).

## MongoDB
In MongoDB add a DB "readinglist" and collection "books" with data to fulfil below struct:
```
type bookData struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt time.Time          `json:"createdAt"`
	Title     string             `json:"title"`
	Published int                `json:"published,omitempty"`
	Pages     int                `json:"pages,omitempty"`
	Genres    []string           `json:"genres,omitempty"`
	Rating    float64            `json:"rating,omitempty"`
	Version   int32              `json:"version,omitempty"`
}
```

## Usage
- Browse through existing book lists.
- Add your own book recommendations to the platform.
- Edit or delete existing book entries.

## Tech Stack
- **Frontend**: HTML, CSS
- **Backend**: Golang, MongoDB

## Dependencies
- MongoDB: You can use a different database, but you'll need to update `db.go` accordingly.
- Required Go modules are listed in `go.mod`.

## Future Improvements
- **Extend UI with ReactJS**: Enhance user experience and interactivity by migrating the frontend to ReactJS.
- **Permissions Management**: Implement user privilege checks to restrict CRUD operations to admin users only.

