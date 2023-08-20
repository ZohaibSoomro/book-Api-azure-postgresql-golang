# Book API with Azure Postgresql in Go

This is a Book API created using Go. It provides various endpoints for managing books.

## Endpoints

- **Get All Books**: Retrieve a list of all books.

  - **Route**: `/api/books`
  - **HTTP Methods**: GET

- **Get Book by ID**: Retrieve a book by its unique identifier.

  - **Route**: `/api/books/{id}`
  - **HTTP Methods**: GET

- **Create a New Book**: Add a new book to the database.

  - **Route**: `/api/book`
  - **HTTP Methods**: POST

- **Update Book by ID**: Modify an existing book by its unique identifier.

  - **Route**: `/api/book/{id}`
  - **HTTP Methods**: PUT

- **Delete a Book by ID**: Remove a book from the database by its unique identifier.

  - **Route**: `/api/book/{id}`
  - **HTTP Method**: DELETE
