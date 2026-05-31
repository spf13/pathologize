## 2024-05-31 - Optimize String Filtering Fast Paths
**Learning:** Combining `strings.IndexFunc` (for a fast zero-allocation pass on clean strings) with `strings.Map` drastically reduces CPU cycles and heap allocations compared to `regexp` for single-character filtering tasks.
**Action:** Always prefer `strings.IndexByte`, `strings.IndexFunc`, and slice operations over `regexp` and `strings.Split` in hot path string manipulation.
