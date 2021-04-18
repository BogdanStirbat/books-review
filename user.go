package main

type User struct {
	ID       int
	Username string
	Email    string
}

func InsertUser(user *User) error {
	var id int
	err := db.QueryRow(`
	        INSERT INTO users(username, email)
					VALUES ($1, $2)
					RETURNING id
	`, user.Username, user.Email).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func GetUserByID(id int) (*User, error) {
	var username, email string
	err := db.QueryRow("SELECT username, email FROM users WHERE id=$1", id).Scan(&username, &email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
	}, nil
}
