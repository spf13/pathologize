## 2024-04-22 - Replacing Regex with Native Strings Functions
**Learning:** Inline regex compilation (e.g., `regexp.MustCompile`) in frequently called functions causes massive performance bottlenecks and heap allocations. The `\s` character class in Go's regex engine strictly equates to `[\t\n\f\r ]` (excluding `\v`).
**Action:** Always prefer native `strings` functions like `strings.TrimRight` with exact character sets over `regexp` for simple character trimming to avoid allocations and improve execution time by orders of magnitude.
