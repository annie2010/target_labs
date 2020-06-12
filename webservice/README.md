<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Web Service Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Turn the book grep into a web service.

> Create a grep service that will accept a book and a word and returns a JSON response
> representing the number of occurrences of that word.

1. Clone the [Labs Repo](https://github.com/gopherland/labs2)
2. Cd webservice
3. Edit internal/handler/count.go and implement a CountHandler
   1. The handler gets a book and a word from the request url
   2. Use the countWords function to compute the number of occurrences
   3. Return a json response that contains the following fields:
      1. Book
      2. Word
      3. Occurrences
   4. Write a test (count_test.go) to make sure your handler is working correctly
4. Edit main.go and define a gorilla mux route to your CountHandler
   1. Use: /api/v1/grep/book/word as your endpoint
   2. Next define a logging middleware to log incoming requests
      1. Use the [Gorilla Handlers Repo](https://github.com/gorilla/handlers)
5. Launch your service and make sure your endpoint is working correctly!

## OSX Install HTTPie (Totally Optional!!)

```shell
# Install httpie
brew install httpie
```

## Expectations

```shell
http :5000/api/v1/grep/3lpigs/pig
# Or...
curl -XGET http://localhost:5000/api/v1/grep/3lpigs/pig
```

Produces...

```text
HTTP/1.1 200 OK
Content-Length: 46
Content-Type: application/json; charset=utf-8
Date: Fri, 08 May 2020 03:36:19 GMT

{
    "Book": "3lpigs",
    "Occurrences": 26,
    "Word": "pig"
}
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
