package main

type Book struct {
	ID          int
	AuthorID    int
	Title       string
	Description string
}

func InsertBook(book *Book) error {
	var id int
	err := db.QueryRow(`
	    INSERT INTO books(author_id, title, description)
			VALUES($1, $2, $3)
			RETURNING id
	`, book.AuthorID, book.Title, book.Description).Scan(&id)
	if err != nil {
		return err
	}
	book.ID = id
	return nil
}

func GetBookByID(id int) (*Book, error) {
	var (
		authorID           int
		title, description string
	)
	err := db.QueryRow("SELECT author_id, title, description FROM books WHERE id=$1", id).Scan(&authorID, &title, &description)
	if err != nil {
		return nil, err
	}
	return &Book{
		ID:          id,
		AuthorID:    authorID,
		Title:       title,
		Description: description,
	}, nil
}

func GetBooks() ([]*Book, error) {
	var books []*Book
	rows, err := db.Query("SELECT id, author_id, title, description FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id, authorID       int
			title, description string
		)
		err = rows.Scan(&id, &authorID, &title, &description)
		if err != nil {
			// handle this error
			panic(err)
		}

		var book = &Book{
			ID:          id,
			AuthorID:    authorID,
			Title:       title,
			Description: description,
		}
		books = append(books, book)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return books, nil
}
