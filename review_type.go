package main

import (
	"github.com/graphql-go/graphql"
)

var ReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Review",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if review, ok := p.Source.(*Review); ok {
					return review.ID, nil
				}
				return nil, nil
			},
		},
		"number_of_stars": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if review, ok := p.Source.(*Review); ok {
					return review.NumberOfStarts, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if review, ok := p.Source.(*Review); ok {
					return review.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	ReviewType.AddFieldConfig("user", &graphql.Field{
		Type: UserType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if review, ok := p.Source.(*Review); ok {
				return GetUserByID(review.UserID)
			}
			return nil, nil
		},
	})
	ReviewType.AddFieldConfig("book", &graphql.Field{
		Type: BookType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if review, ok := p.Source.(*Review); ok {
				return GetBookByID(review.BookID)
			}
			return nil, nil
		},
	})
}
