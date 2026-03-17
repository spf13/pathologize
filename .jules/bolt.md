## 2024-03-17 - Precompute static arrays in Go to prevent huge allocation bottlenecks
**Learning:** In Go, splitting a static string like `strings.Split("A B C", " ")` inside a function dynamically allocates memory on *every* call, causing substantial GC pressure and latency for hot path functions (like those parsing path elements repeatedly).
**Action:** Always precompute dynamically calculated static variables (like slices derived from splits) as package-level variables or inside `init()` blocks instead of recalculating them inside the function body.
