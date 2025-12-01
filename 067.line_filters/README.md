```bash
echo 'hello'   > /tmp/lines
echo 'filter' >> /tmp/lines
```

---

```bash
cat /tmp/lines | go run line-filters.go
```

```plaintext
HELLO
FILTER
```
