
## 2024-03-26 - Regex replacement in Go for trailing space/dots

**Learning:** Using `regexp.MustCompile` and `re.ReplaceAllString` for simple trailing character removal (like `[.\s]+$`) is notoriously slow in Go because of Regex engine overhead. `strings.TrimRight` with a set of trailing characters (`". \t\n\v\f\r"`) is significantly faster and uses 0 allocations, taking only ~23 ns/op compared to ~2500 ns/op with regex.
**Action:** Always consider standard string trimming operations like `strings.TrimRight` before defaulting to regex for simple character matching tasks at boundaries.
