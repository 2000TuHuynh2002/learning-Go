```bash
go test -v
```

```plaintext
=== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      example/testing_and_benchmarking        1.004s
```

---

```bash
go test -bench=.
```

```plaintext
goos: windows
goarch: amd64
pkg: example/testing_and_benchmarking
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
BenchmarkIntMin-16      774210031                1.538 ns/op
PASS
ok      example/testing_and_benchmarking        1.814s
```
