## 2024-04-15 - Inline Regexp Compilation Bottleneck

**Learning:** In Go, calling `regexp.MustCompile` inside a frequently executed function (like `removeTrailing` during path sanitization) causes massive performance degradation and unnecessary heap allocations on every call. Native string manipulation functions, such as `strings.TrimRight(filename, ".\t\n\f\r ")`, are perfectly equivalent to regex operations like `[.\s]+$` but execute orders of magnitude faster (e.g., dropping from ~3400 ns/op to ~21 ns/op).

**Action:** Always scan for inline `regexp.MustCompile` calls in hot paths and either hoist them to global package-level variables or, ideally, replace them entirely with highly optimized native `strings` functions when dealing with simple character classes.
