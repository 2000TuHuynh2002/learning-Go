```bash
mkdir -p folder
echo "hello go" > folder/single_file.txt
echo "123" > folder/file1.hash
echo "456" > folder/file2.hash
```

---

```bash
go run embed-directive.go
```

```plaintext
hello go
hello go
123
456
```
