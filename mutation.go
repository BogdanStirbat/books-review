package main

import (
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
	},
})
