package main

type Author struct {
	ID   int
	Name string
}

func InsertAuthor(author *Author) error {
	var id int
	err := db.QueryRow(`
	    INSERT INTO authors(name)
			VALUES($1)
			RETURNING id
	`, author.Name).Scan(&id)
	if err != nil {
		return err
	}
	author.ID = id
	return nil
}

func GetAuthorByID(id int) (*Author, error) {
	var name string
	err := db.QueryRow("SELECT name FROM authors WHERE id=$1", id).Scan(&name)
	if err != nil {
		return nil, err
	}
	return &Author{
		ID:   id,
		Name: name,
	}, nil
}
