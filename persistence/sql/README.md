<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Persistence SQL Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Your Mission

> Barnes And Noble brick and mortar stores are suffering and you've been tasked in developing
> a web service to showcase their vast inventory of books and authors.
> No pressure! But you're their last Hoorah to save their Biz!!

> The web service should surface the following endpoints:
>  * /api/v1/books -- lists out all books
>  * /api/v1/authors -- lists out all authors
>  * /api/v1/books/author_name -- list out all books from the given author last name

1. Clone the [labs repo](https://github.com/gopherland/target_labs)
2. Cd persistent/sql
3. The service migration and seeding are already implemented for you
4. Using the provided command below deploy a postgres container on your Docker instance. Be sure to note your connection details!
5. In internal/pg/dial.go implement the necessary logic to connect to your postgres containerized instance
6. In internal/model/author.go implement the method to list all authors
7. In internal/model/book.go:
   1. Implement the method to retrieve all books
   2. Implement the method ByAuthor to retrieve a list of books given an author last name.
8. Run your web service and ensure all endpoints are working nominally
9. Terminate your service and your postgres container
10. BONUS!! Leverage a prepared statement to handle /api/v1/books/author_name queries!
11. BONUS!! Note the response times given by the request logger.
    What can you do to improve your service performance?

## Commands

### Run PostgresSQL via Docker

```shell
docker run --rm --name pg -p 5432:5432 -e POSTGRES_PASSWORD=YOUR_PASSWORD -e POSTGRES_DB=YOUR_DB_NAME -d postgres@12.3
```

### Run psql CLI in container

```shell
docker exec -it --rm pg psql -U postgres -W -d YOUR_DB_NAME
```

### Postgres CLI Survival Cheats

```shell
# Exit
\q
# List all databases
\l
# Use a given db fred
\c fred
# List all tables
\d
# Show table fred info
\d fred
# Check query planner
explain select * from fred;
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)