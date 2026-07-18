# GOLANG

My journey to becoming a Go backend developer — following a structured
**6-month roadmap**. This repo tracks everything from language basics to,
eventually, a full distributed workflow platform.

> Starting point: comfortable with another language, some Go exposure.
> Commitment: 4–6 hrs/day.

---

## 📁 Repository structure

```
GOLANG/
├── BASICS/                         # Go fundamentals — start here
│   ├── 01-hello/ … 13-goroutines/  # one runnable topic per folder
│   ├── go.mod                      # Go module definition
│   └── README.md                   # detailed guide for the basics
├── GO_ROADMAP.docx                 # the full 6-month plan
└── CEX_Companion_Project_Guide.docx
```

Each learning area lives in its own top-level folder. Right now that's
`BASICS/`; larger project modules (the "Conduit" platform from the roadmap)
will be added as separate folders as I progress.

---

## 🚀 Getting started

**Prerequisites:** [Go](https://go.dev/dl/) 1.22+ installed
(`go version` to check).

Clone and run any example:

```bash
git clone https://github.com/SrijitGyawali/GOLANG.git
cd GOLANG/BASICS
go run ./01-hello
```

See [`BASICS/README.md`](BASICS/README.md) for the full topic list, the
recommended learning order, and how the module is organized.

---

## 🗺️ Roadmap progress

The plan builds one cohesive platform (**Conduit** — a cloud workflow
engine) rather than scattered toy projects.

| Phase | Focus | Status |
|-------|-------|--------|
| Month 1 · Wk 1 | Go idioms & concurrency core | 🟡 In progress |
| Month 1 · Wk 2 | Go internals & performance (pprof) | ⚪ Not started |
| Month 1 · Wk 3 | REST APIs with Gin — Auth Service | ⚪ Not started |
| Month 1 · Wk 4 | PostgreSQL for real applications | ⚪ Not started |
| Month 2 | Testing, architecture & Job Service | ⚪ Not started |
| Month 3 | Microservices & distributed systems | ⚪ Not started |
| Month 4 | Cloud-native & DevOps | ⚪ Not started |
| Month 5 | Observability, security & workflow engine | ⚪ Not started |
| Month 6 | Polish, system design & job hunt | ⚪ Not started |

Legend: ✅ Done · 🟡 In progress · ⚪ Not started

---

## 📚 What I'm learning here

- **Go fundamentals:** variables, control flow, functions, structs,
  interfaces, error handling, and concurrency (goroutines & channels).
- **Engineering habits:** reading production source code, writing about
  what I build, and keeping a clean, meaningful Git history.
- **Systems foundations:** Linux, networking, and advanced Git, applied to
  real code rather than throwaway examples.

---

## 📄 License

Personal learning repository. Feel free to read and learn from it.
