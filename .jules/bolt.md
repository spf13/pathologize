## 2024-04-22 - Replacing Regex with Native Strings Functions
**Learning:** Inline regex compilation (e.g., `regexp.MustCompile`) in frequently called functions causes massive performance bottlenecks and heap allocations. The `\s` character class in Go's regex engine strictly equates to `[\t\n\f\r ]` (excluding `\v`).
**Action:** Always prefer native `strings` functions like `strings.TrimRight` with exact character sets over `regexp` for simple character trimming to avoid allocations and improve execution time by orders of magnitude.
## 2024-04-23 - Pre-computing Constant String Splits
**Learning:** Calling `strings.Split` on constant strings inside frequently called functions (like `removeReservedNames` in path sanitization) causes unnecessary heap memory allocations and CPU overhead on every invocation.
**Action:** Always pre-compute and store the results of splitting constant strings in package-level variables during initialization to avoid repeated allocations in critical paths.
