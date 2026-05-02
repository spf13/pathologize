## 2024-04-22 - Replacing Regex with Native Strings Functions
**Learning:** Inline regex compilation (e.g., `regexp.MustCompile`) in frequently called functions causes massive performance bottlenecks and heap allocations. The `\s` character class in Go's regex engine strictly equates to `[\t\n\f\r ]` (excluding `\v`).
**Action:** Always prefer native `strings` functions like `strings.TrimRight` with exact character sets over `regexp` for simple character trimming to avoid allocations and improve execution time by orders of magnitude.
## 2024-05-18 - Pre-calculate static lists to avoid allocations
**Learning:** `strings.Split` inside frequently called functions (e.g., `removeReservedNames`, called multiple times per `Clean` invocation) causes significant, unnecessary memory allocations (704 B/op in this case).
**Action:** Extract operations like `strings.Split` on static constant strings to package-level variables so they are computed only once during initialization.
