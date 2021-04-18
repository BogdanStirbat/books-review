package main

type Review struct {
	ID             int
	UserID         int
	BookID         int
	NumberOfStarts int
	Body           string
}

func InsertReview(review *Review) error {
	var id int
	err := db.QueryRow(`
	    INSERT INTO reviews(user_id, book_id, number_of_stars, body)
			VALUES($1, $2, $3, $4)
			RETURNING id
	`, review.UserID, review.BookID, review.NumberOfStarts, review.Body).Scan(&id)
	if err != nil {
		return err
	}
	review.ID = id
	return nil
}

func GetReviewByID(id int) (*Review, error) {
	var (
		userID, bookID, numberOfStars int
		body                          string
	)
	err := db.QueryRow("SELECT user_id, book_id, number_of_stars, body FROM reviews WHERE id=$1", id).Scan(&userID, &bookID, &numberOfStars, &body)
	if err != nil {
		return nil, err
	}
	return &Review{
		ID:             id,
		UserID:         userID,
		BookID:         bookID,
		NumberOfStarts: numberOfStars,
		Body:           body,
	}, nil
}
