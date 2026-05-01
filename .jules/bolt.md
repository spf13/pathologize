## 2024-04-22 - Replacing Regex with Native Strings Functions
**Learning:** Inline regex compilation (e.g., `regexp.MustCompile`) in frequently called functions causes massive performance bottlenecks and heap allocations. The `\s` character class in Go's regex engine strictly equates to `[\t\n\f\r ]` (excluding `\v`).
**Action:** Always prefer native `strings` functions like `strings.TrimRight` with exact character sets over `regexp` for simple character trimming to avoid allocations and improve execution time by orders of magnitude.
## 2024-05-01 - Prevent repeated memory allocations in hot paths
**Learning:** The `removeReservedNames` function in `github.com/spf13/pathologize` repeatedly performed a `strings.Split` operation on a constant string representing dos and windows reserved names. This function is called multiple times per `Clean` path evaluation, resulting in significant unnecessary heap allocations and garbage collection overhead.
**Action:** Pre-calculate the result of operations like `strings.Split` on constant strings and store them in package-level variables to be reused across function calls to optimize hot execution paths.
