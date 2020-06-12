<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Web Service Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Turn your Grep library into a gRPC BookGrep service

> Create a grep gRPC service that will accept a book and a word and returns a response
> representing the number of occurrences of that word.

1. Clone the [Labs Repo](https://github.com/gopherland/labs2)
2. Cd grpc
3. Edit the proto definition file and define your service definition
   1. Use a unary protocol for both client and server
4. Implement your gRPC service to return the correct response by invoking grep to gather the counts.
   1. Return a response that contains the following fields:
      1. Book
      2. Word
      3. Total
5. Write a test (grepper_test.go) to make sure your gRPC handler is working correctly
   1. Use testdata/fred.txt to test your handler
6. Run your gRPC server
   1. Next run your client and make sure you can connect and it returns the correct payload
   2. Now install gRPCURL CLI
      1. List and Describe your service (commands below)
      2. Next Hit the server grep handler and ensure gRPCURL returns the correct answer
7. Edit your server implementation and define 2 custom interceptors:
   1. A logging interceptor to log the server requests
   2. A Measure interceptor to track the requests duration
8. Launch your service and make sure your endpoint is correctly decorated!

## Expectations

```shell
   go run cmd/client/main.go
```

Produces...

```text
2020/05/19 13:38:27 Client Dialing "localhost:50052"...
2020/05/19 13:38:27 Book: 3lpigs
2020/05/19 13:38:27 Word: pig
2020/05/19 13:38:27 Count: 26
```

## Installation

### OSX Install Protobuf compiler

```shell
# Install xcode dev tools
sudo xcode-select --install
# Install protobuff compiler
brew install protobuf
# Install GO libraries
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

For other platforms please see the [install instructions](https://github.com/protocolbuffers/protobuf/blob/master/README.md#protocol-compiler-installation)

### gRPCURL

```shell
brew install grpcurl
```

## Commands

### Generate stubs

```shell
# Generate stubs
# NOTE!! On OSX, you may need to enable security for this binary.
protoc -I . protos/grep.proto --go_out=plugins=grpc:.
```

### gRPCURL

```shell
# List server capabilities
grpcurl -v -plaintext localhost:50052 list
# Describe grep API
grpcurl -v -plaintext localhost:50052 describe grep.Grepper
# Issue request
grpcurl -plaintext -d '{"book": "3lpigs", "word": "pig"}' localhost:50052 grep.Grepper/Grep
 ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
