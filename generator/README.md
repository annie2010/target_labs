<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Generator Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Implement Generics in GO++!

> Using GO generate, write a program to create stack data structure templates a couple data types.

1. Clone the [Labs Repo](https://github.com/gopherland/labs2)
1. cd generator
1. Annotate generator/main.go to call `stacker` cli to generate templated stacks implementation
   1. Stacker should takes 2 arguments:
      1. -t type1,type2 specifying which types of stacks to generate
      1. -p the name of the package to generate the implementation destination
   2. Generate 2 stacks implementation for: float64 and i1t32
1. cd stacker
1. Implement stacker (stacker/main.go) to generate a templated stack and its associated test
   1. The GO templates for both source and test are given (see tpl.go)
   2. Implement stacker to generate the stacks code and test in the package location
   3. Test your implementation by generating a couple templatized stacks
   4. Run your test samples!
   5. Install stacker on your system so it's available in your path ie GOBIN is set and in your PATH
1. Back in the generator directory edit generator/main.go
   1. Generate your code so that generator/main.go finds your generated types: stacks.Float64 and stacks.Int32
   2. Make sure your stacks import path is set
   3. Run your tests!
   4. Run your main application to ensure both stacks implementations are operating as expected!

### Expectations

The output from generator/main.go should look like this...

```go
2020/05/07 15:28:15 🥞 Float64    Pop:42.25 -- Top:20.2 -- Peek:20.2,10.5
2020/05/07 15:28:15 📚 Int32      Pop:300   -- Top:100  -- Peek:100,200
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)