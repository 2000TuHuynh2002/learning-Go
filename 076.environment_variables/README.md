```bash
go run ./076.environment_variables/main.go
```

```plaintext
FOO: 1
BAR:

TERM_PROGRAM
PATH
SHELL
...
FOO
```

---

```bash
BAR=2 go run ./076.environment_variables/main.go
```

```plaintext
FOO: 1
BAR: 2

TERM_PROGRAM
PATH
SHELL
...
FOO
```
