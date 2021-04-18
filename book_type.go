package main

import (
	"github.com/graphql-go/graphql"
)

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if author, ok := p.Source.(*Book); ok {
					return author.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(*Book); ok {
					return book.Title, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(*Book); ok {
					return book.Description, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	BookType.AddFieldConfig("author", &graphql.Field{
		Type: AuthorType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if book, ok := p.Source.(*Book); ok {
				return GetAuthorByID(book.AuthorID)
			}
			return nil, nil
		},
	})
}
