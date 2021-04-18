package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "New Username",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Description: "New User email",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				username := p.Args["username"].(string)
				email := p.Args["email"].(string)
				user := &User{
					Username: username,
					Email:    email,
				}
				err := InsertUser(user)
				return user, err
			},
		},
		"createAuthor": &graphql.Field{
			Type: AuthorType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "Author name",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				author := &Author{
					Name: name,
				}
				err := InsertAuthor(author)
				return author, err
			},
		},
		"createBook": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"author": &graphql.ArgumentConfig{
					Description: "Id of author who wrote the book",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"title": &graphql.ArgumentConfig{
					Description: "Title of the book",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Description: "Descrition of the book",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["author"].(string)
				authorID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				title := p.Args["title"].(string)
				description := p.Args["description"].(string)
				book := &Book{
					AuthorID:    authorID,
					Title:       title,
					Description: description,
				}
				err = InsertBook(book)
				return book, err
			},
		},
		"createReview": &graphql.Field{
			Type: ReviewType,
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Description: "Id of user who created the review",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"book": &graphql.ArgumentConfig{
					Description: "Id of the reviewed book",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"number_of_stars": &graphql.ArgumentConfig{
					Description: "Number of stars",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "Body of review",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["user"].(string)
				userID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				i = p.Args["book"].(string)
				bookID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				i = p.Args["number_of_stars"].(string)
				numberOfStars, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				body := p.Args["body"].(string)
				review := &Review{
					UserID:         userID,
					BookID:         bookID,
					NumberOfStarts: numberOfStars,
					Body:           body,
				}
				err = InsertReview(review)
				return review, err
			},
		},
	},
})
