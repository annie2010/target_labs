<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# GraphQL Lab

## Queries

1. Find all books

      ```graphql
      query{
        allBooks{
          id
          title
          category
          authors{
            id
            first
            last
          }
        }
      }
      ```

1. Find all authors

      ```graphql
      query {
        allAuthors {
          id
          first
          last
        }
      }
      ```

1. Books by Author

      ```graphql
      query {
        booksByAuthor(id: "isan-0") {
          title
          category
          authors {
            first
            last
          }
        }
      }
      ```

## Mutations

1. Create author

   ```graphql
   mutation CreateAuthor($input: authorInput!) {
     createAuthor(input: $input) {
       id
       first
       last
     }
   }
   ```

1. Delete author

      ```graphql
      mutation deleteAuthor($id: ID!){
        deleteAuthor(id: $id) {
          id
          first
          last
        }
      }
      ```


---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)