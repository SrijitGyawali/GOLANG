# Go Basics — Learning Playground

This folder is a hands-on tour of Go fundamentals, one topic per folder,
in the order you should learn them. It maps to **Month 1, Week 1** of your
roadmap ("Go idioms & concurrency core").

---

## 0. First: install Go (you don't have it yet)

1. Download the Windows installer from **https://go.dev/dl/** (pick the
   `.msi` for `windows-amd64`).
2. Run it — it installs Go to `C:\Program Files\Go` and adds `go` to your PATH.
3. **Close and reopen** your terminal (VS Code too), then verify:

   ```
   go version
   ```

   You should see something like `go version go1.23.x windows/amd64`.

---

## 1. How this folder is structured

```
BASICS/
├── go.mod                 <- defines the "module" (project root for Go)
├── README.md              <- this file
├── 01-hello/main.go       <- printing
├── 02-variables/main.go   <- variables & constants
├── 03-datatypes/main.go   <- int, float, string, bool, conversion
├── 04-conditionals/main.go<- if / else / switch
├── 05-loops/main.go       <- for (Go's only loop) + range
├── 06-functions/main.go   <- params, multiple returns, variadic
├── 07-arrays-slices/main.go
├── 08-maps/main.go
├── 09-structs/main.go     <- Go's version of objects
├── 10-pointers/main.go
├── 11-errors/main.go      <- Go's error handling (no exceptions)
├── 12-interfaces/main.go
└── 13-goroutines/main.go  <- concurrency (Go's superpower)
```

### Why one folder per topic?
Every runnable Go program needs a `package main` with a `func main()`.
You can only have **one** `main()` per folder. So each topic lives in its
own folder with its own `main.go`. The single `go.mod` at the top makes
them all part of one module named `basics`.

### What is `go.mod`?
It marks the root of a Go module and records the Go version. You create it
once with `go mod init <name>` — it's already done here (`module basics`).

---

## 2. How to run a program

Open a terminal **in the `BASICS` folder**, then run any topic by its folder:

```
go run ./01-hello
go run ./05-loops
go run ./13-goroutines
```

`go run` compiles and runs in one step — nothing to clean up afterward.

Other useful commands:

```
go build ./01-hello     # produces an .exe you can run directly
go fmt ./...            # auto-formats ALL your code (run this often!)
go vet ./...           # catches common mistakes
```

> Tip: `go fmt` is not optional in Go culture — the whole community uses the
> same formatting. Run it before every commit.

---

## 3. Suggested learning order

Go through the folders **in number order** (01 → 13). For each one:

1. **Read** `main.go` top to bottom — the comments explain every line.
2. **Run** it with `go run ./NN-topic` and match the output to the code.
3. **Break it** — change a value, delete a line, see the error. This is the
   fastest way to actually learn (errors in Go are very readable).
4. **Rebuild from memory** — close the file and rewrite it yourself. The
   roadmap calls this out: "rewrite every example yourself from memory — no
   copy-paste."

---

## 4. What each topic teaches

| Folder | You'll learn |
|--------|--------------|
| 01-hello | `package main`, `import`, `Println` vs `Printf` |
| 02-variables | `var`, `:=`, zero values, constants |
| 03-datatypes | number/string/bool types, `%T`, type conversion |
| 04-conditionals | `if`/`else`, `switch`, if-with-statement |
| 05-loops | the single `for` keyword, `break`, `continue`, `range` |
| 06-functions | multiple return values, variadic, functions as values |
| 07-arrays-slices | fixed arrays vs dynamic slices, `append`, slicing |
| 08-maps | key/value stores, the "comma ok" idiom |
| 09-structs | grouping data, methods, value vs pointer receivers |
| 10-pointers | `&` and `*`, pass-by-value vs pass-by-pointer |
| 11-errors | the `if err != nil` pattern, sentinel errors |
| 12-interfaces | behavior-based typing, no `implements` keyword |
| 13-goroutines | `go`, `WaitGroup`, channels — Go's concurrency model |

---

## 5. After you finish these

You'll be ready for the roadmap's **Week 1 warm-up drill**: a concurrent
worker pool. It combines goroutines (13), channels (13), slices (07), and
functions (06) — everything here. That's your bridge from "basics" into
building **Conduit**, your headline project.

Next natural steps once these click:
- **Week 2:** `pprof` profiling & benchmarks (`testing.B`)
- **Week 3:** your first real service — the Auth Service with Gin
