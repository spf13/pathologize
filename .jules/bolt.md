## 2024-06-07 - Avoid strings.SplitN for prefix extraction
**Learning:** Using `strings.SplitN(str, delim, 2)[0]` to extract a prefix up to a delimiter causes unnecessary heap allocations for the slice. In Go, it's significantly faster and avoids allocations to use `strings.IndexByte(str, delim)` combined with string slicing (`str[:idx]`).
**Action:** When extracting a substring up to a specific single-byte delimiter (like a file extension), always prefer `strings.IndexByte` and slicing over `strings.Split` to achieve zero-allocation performance.
