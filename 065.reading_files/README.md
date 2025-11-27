```bash
echo "hello" > /tmp/dat
echo "go" >>   /tmp/dat
```

---

```bash
go run reading-files.go
```

```plaintext
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello
```
