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
