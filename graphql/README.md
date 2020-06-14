<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# GraphQL Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Barnes And Noble wants to be part of the cool kids crew and wants to
> surface there large library of books as a graphql API so they can integrate
> with other vendors. You've been tasked to develop a graphQL server for them.

> The service needs should surface the following functionality:
>  * list all books
>  * list all authors
>  * list all books by author
>  * create a new author
>  * delete an author by ISAN
>  * delete a book by ISBN

1. Implement your service using [gqlgen](https://github.com/99designs/gqlgen)
2. Define your GraphQL schema to support the required queries and mutations
3. Using GraphIQL ensure your service is providing the correct functionality by insuing queries and mutations

## Commands

### Generate the code

```shell
go generate
```

```shell
docker run -it --name dgraph -p 8080:8080 dgraph/standalone:v20.03.1
https://github.com/prisma-labs/graphql-playground/releases/tag/v1.8.10
http localhost:8080/admin/schema < schema.gql
brew cask install graphql-playground
```


```gql
mutation {
  addActor(input: [
    {first: "Fernand", last: "Galiana"},
    {first: "Fred", last: "Blee"},
  ]) {
    actor{
      first
      last
    }
  }
}

query{
  queryActor{
    first
    last
  }
}
```

```gql
mutation {
  addMovie(input: [
     {
      title: "Rango",
      year: 2010,
      actors: [
        { id: "0x2" }
      ]
    }
  ]) {
      movie{
        title
        year
      }
    }
}

query{
  queryMovie{
    title
    year
    actors{
      first
      last
    }
  }
}
```
---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)