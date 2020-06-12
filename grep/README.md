<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Benchmarking Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission

> Implement Grep!

> Your CLI application should take a word and a file name args and
> report back the number of occurrences of the word in the file.

* Clone the [labs repo](https://github.com/gopherland/labs2)
* Cd grep
* An initial implementation of `CountWord` function is provided
* Implement a test for the `CountWord` function
* Benchmark your CountWord function.
* Is it cpu or memory bound?
* Using `benchstat` checkout to variation and make sure you have a solid measurement.
* Can you speed up your initial implementation?
* Implement and test a second implementation.
* Using `benchstat` compare your first and second implementations.

### Setup

```shell
# Install benchstat
# IMPORTANT! Make sure GOBIN is set to $HOME/gopherland/bin and is in your PATH!!
go env GOBIN
# Set if NOT SET!
go env -w GOBIN=$HOME/gopherland/bin
go get -u golang.org/x/perf/cmd/benchstat
```

### Commands

```shell
# Compare implementations
go test --run xxx --bench V1 --count 10 --benchmem | tee v1.out
# Ensure the variance is cool for a single run
benchstat v1.out
# Run the second implementation.
go test --run xxx --bench V2 --count 10 --benchmem | tee v2.out
# Normalize output so they have the same benchmark name
sed -i '' 's/V1//g' v1.out && sed -i '' 's/V2//g' v2.out
# Compare the two implementations
benchstat v1.out v2.out
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)