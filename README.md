# 🧠 Golang Data Structures & Algorithms 🚧

This repository demonstrates core **data structures** and **algorithms** implemented in **Go (Golang)** — with a focus on performance, clarity, and idiomatic code.

It includes (or will include at some point):
- Binary Search Trees (BST)
- Tree-based sorting (TreeSort)
- Binary Search
- Custom data structures (linked lists, stacks, queues)
- Sorting algorithms (merge, quick, bubble, etc.)
- Efficient use of Go slices, maps, and pointers

---

## 📚 Contents

| Feature | Description |
|--------|-------------|
| `trees/` | Binary trees, insertion, traversal, TreeSort |
| `search/` | Binary search (iterative and recursive) |
| `sort/` | Implementations of Merge Sort, Quick Sort, Bubble Sort |
| `structures/` | Stacks, Queues, Linked Lists, and Circular Buffers |
| `graphs/` | (coming soon) Graph representation & traversal (DFS, BFS) |
| `utils/` | Helper functions, benchmarking tools |

---

## 🚀 Example: TreeSort

```go
values := []int{5, 1, 9, 2, 6}
TreeSort(values)
fmt.Println("Sorted:", values)

