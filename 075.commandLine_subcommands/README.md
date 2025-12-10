```bash
go run ./075.commandLine_subcommands/main.go foo -enable -name=joe a1 a2
```

```plaintext
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]
```

---

```bash
go run ./075.commandLine_subcommands/main.go bar -level 8 a1
```

```plaintext
subcommand 'bar'
  level: 8
  tail: [a1]
```

---

```bash
go run ./075.commandLine_subcommands/main.go bar -enable a1
```

```plaintext
flag provided but not defined: -enable
Usage of bar:
  -level int
        level
```
