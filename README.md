# Book reviews API

This is an GraphQL API for reviewing books, written in Go, using `github.com/graphql-go/graphql` and Postgres DB.

## Getting started

It is assumed that Go is already installed. To run this project, you first need to setup a Postgres DB database. For this, you can either install it on your machine, either use a Docker image.

Instructions for using Postgres DB inside a Docker container:
 - create and start a container:
   `docker run --name <container-name> -p 5432:5432 -e POSTGRES_PASSWORD=somePassword -d postgres # <container-name> is a container name, example postgresql-container-books`

 - connect to the container, to create the database
   `docker exec -it <container-name> psql -U postgres # <container-name> is the container name added above, example postgresql-container-books` 
   `create user book_db_user;`
   `create database book_db;`
   `alter user book_db_user with encrypted password 'abc123def';`
   `grant all privileges on database book_db to book_db_user;`

If you need later to start this container, just run `docker start <container-name>`

## Using the API

Now, the DB is configured, and you can just thart the Go backend: `go run .`

On `http://localhost:8080/graphql`, you can execute now GraphQL queries, either using Postman or the command line:
```bash
curl -XPOST http://localhost:8080/graphql -d 'mutation {createUser(username:"test1", email:"test1@test.com"){id, username, email}}'
curl -XPOST http://localhost:8080/graphql -d '{user(id:1){id, username, email}}'
curl -XPOST http://localhost:8080/graphql -d 'mutation {createAuthor(name:"Author 1"){id, name}}'
curl -XPOST http://localhost:8080/graphql -d '{author(id:1){id, name}}'
curl -XPOST http://localhost:8080/graphql -d 'mutation {createBook(author: 1, title: "Title 1", description: "Description 1"){id, title, description, author{id, name}}}'
curl -XPOST http://localhost:8080/graphql -d '{book(id:1){id, title, description, author{id, name}}}'
curl -XPOST http://localhost:8080/graphql -d 'mutation {createBook(author: 1, title: "Title 1", description: "Description 1"){id, title, description, author{id, name}}}'
curl -XPOST http://localhost:8080/graphql -d 'mutation {createReview(user: 1, book: 1, number_of_stars: "5", body: "A very good book."){id, number_of_stars, body, user{id, username, email}, book{id, title, description}}}'
curl -XPOST http://localhost:8080/graphql -d '{review(id:1){id, number_of_stars, body, user{id, username, email}, book{id, title, description}}}'
```

## References
https://graphql.org/code/#go
https://www.howtographql.com/basics/0-introduction/
https://www.inovex.de/blog/graphql-application-golang-tutorial/
https://www.freecodecamp.org/news/react-apollo-client-2020-tutorial/
http://alexandrutopliceanu.ro/post/graphql-with-go-and-postgresql/
