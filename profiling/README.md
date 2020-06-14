<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Profiling Lab

## Commands

1. Run fib benchmarks

```shell
cd internal/fib
go test --run xxx --bench Rec --cpuprofile cpur.out
go test --run xxx --bench Iter --cpuprofile cpui.out
go tool pprof -base cpui.out cpur.out
```

1. Run Handler benchmarks

```shell
cd internal
go test --run xxx --bench .
go test --run xxx --bench Rec --memprofile memr.out
go test --run xxx --bench Iter --memprofile memi.out
go tool pprof -base memi.out memr.out
```

1. Benchstat

```shell
go test --run xxx --bench Rec --count=10 | tee fib1.out
go test --run xxx --bench Iter --count=10 | tee fib2.out
sed -i '' 's/Rec//g' fib1.out && sed -i '' 's/Iter//g' fib2.out
benchstat fib1.out fib2.out
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)