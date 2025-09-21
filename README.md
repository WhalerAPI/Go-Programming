# 🧠 Go Programming

This repository demonstrates core **data structures** and **algorithms** implemented in **Go (Golang)** — with a focus on performance, clarity, and idiomatic code.

---

## 📚 Contents

| Feature | Description |
|--------|-------------|
| `tree/` | Binary trees, insertion, traversal, TreeSort |
| `sort/` | Implementations of Merge Sort, Quick Sort, Bubble Sort |
| `structures/` | Stacks, Queues, Linked Lists, and Circular Buffers |
| `utils/` | Helper functions, benchmarking tools |
| `web/`  | Handlers, Route-mapping, HTML/TEXT Templates
| `network/`  | TCP/UDP, Web Socket Programming, UNIX, Transports, Clients
| `data/` | Marshalling data, pointers, *DB type, Handlers

---

## 🚀 Example: TreeSort

```go
values := []int{5, 1, 9, 2, 6}
TreeSort(values)
fmt.Println("Sorted:", values)
```
```mermaid
---
config:
  theme: redux-dark
---
flowchart TD
 subgraph Core["Core"]
        errors["errors"]
        builtin["builtin"]
        fmt["fmt"]
        strconv["strconv"]
        strings["strings"]
        bytes["bytes"]
        unicode["unicode"]
        utf8["utf8"]
  end
 subgraph IO["IO"]
        io["io"]
        os["os"]
        bufio["bufio"]
        textTemplate["text/template"]
        htmlTemplate["html/template"]
  end
 subgraph Data["Data"]
        encodingJSON["encoding/json"]
        encodingCSV["encoding/csv"]
        encodingXML["encoding/xml"]
        encodingBase64["encoding/base64"]
        encodingHex["encoding/hex"]
  end
 subgraph Net["Net"]
        netHTTP["net/http"]
        net["net"]
        netURL["net/url"]
  end
 subgraph Concurrency["Concurrency"]
        context["context"]
        sync["sync"]
        atomic["atomic"]
  end
    builtin --> errors & fmt & strconv & strings & bytes & unicode & utf8
    os --> io
    io --> bufio & textTemplate & htmlTemplate
    net --> netHTTP & netURL
    sync --> context & atomic
    context --> netHTTP
    Core --> IO
    IO --> Data
    Net --> Data
```


