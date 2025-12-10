```bash
go run ./074.commandLine_flags/main.go -word=opt -numb=7 -fork -svar=flag
```

```plaintext
word: opt
numb: 7
fork: true
svar: flag
tail: []
```

---

```bash
go run ./074.commandLine_flags/main.go -word=opt
```

```plaintext
word: opt
numb: 42
fork: false
svar: bar
tail: []
```

---

```bash
go run ./074.commandLine_flags/main.go -word=opt a1 a2 a3
```

```plaintext
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3]
```

---

```bash
go run ./074.commandLine_flags/main.go -word=opt a1 a2 a3 -numb=7
```

```plaintext
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]
```

---

```bash
go run ./074.commandLine_flags/main.go -h
```

```plaintext
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
```

---

```bash
go run ./074.commandLine_flags/main.go -wat
```

```plaintext
flag provided but not defined: -wat
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
```
