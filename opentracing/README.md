<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# OpenTracing G.O.T Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Game Of Thrones (G.O.T) Reloaded!

> In this lab, we are going to decorate a web server using **OpenTracing**.
> There are 2 services involved: **Castle** and **Knight**. The
> Knights want to melt Castles, but if you're a G.O.T fan, you already
> know that only the NightKing can melt a Castle using his undead dragon 🙀🐉...
> The Knight service accepts post requests on */api/v1/melt* and issues a
> post */api/v1/melt* on the Castle service with a given Knight name.
> The Castle service returns either a 200 with a castle melted message if the
> knight is the `NightKing` 😵 or a 417 error with *only NightKing can melt* otherwise.

1. Clone the [labs repo](https://github.com/gopherland/labs2)
2. cd opentracing
3. Instrument the Castle service (cmd/castle/root.go) by tracing incoming *melt* requests
   1. Create a top level span for all new incoming requests
   2. Decorate your span to indicate who is the knight trying to melt the castle
   3. Edit your newSpanFromReq function child span and add the following info:
      1. http.method=the incoming request method
      2. http.url=the incoming request url
      3. component=the component name that received the request
   4. Trace the readQuest function
      1. Create a new span from the given context
      2. Add a tag action=castle.readQuest
      3. Add log to trace the name of the knight issuing a melt request
         1. The message should be of the form `knight xxx requested a melt`
   5. If the given Knight is *NightKing* add a log to the castle span to indicate `the castle is melted`.
4. All other knights should produce a span error (internal/http.go:SpanError).
5. Span errors are indicated as follows:
   1. Setting a span tag error=true
   2. Adding a structured log on the span using
      1. event=error
      2. message=only the NightKing can melt the castle
6. Using the provided docker command start the Jaeger service
7. In a separate terminals start your Castle and Knight services.
8. Using the Jaeger Dashboard (see command below) validate that your traces are correctly tracking the workload by using different knights.

## Commands

1. Download and Install Docker on your machine [SKIP IF ALREADY INSTALLED!]
   1. See [Docker install](https://www.docker.com/products/docker-desktop) instructions
2. Start a Jaeger server

   ```shell
   docker run --name jaeger -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
   ```

3. Jaeger Dashboard

   ```shell
   open http://localhost:16686
   ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
