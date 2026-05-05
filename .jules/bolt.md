## 2024-04-22 - Replacing Regex with Native Strings Functions
**Learning:** Inline regex compilation (e.g., `regexp.MustCompile`) in frequently called functions causes massive performance bottlenecks and heap allocations. The `\s` character class in Go's regex engine strictly equates to `[\t\n\f\r ]` (excluding `\v`).
**Action:** Always prefer native `strings` functions like `strings.TrimRight` with exact character sets over `regexp` for simple character trimming to avoid allocations and improve execution time by orders of magnitude.
## 2024-05-24 - Pre-calculating strings.Split in hot paths
**Learning:** Calling `strings.Split` on a static string inside a frequently called function (like `removeReservedNames` which is called multiple times per `Clean` invocation) causes unnecessary memory allocations and CPU overhead on every call.
**Action:** Pre-calculate and store the result of `strings.Split` in a package-level variable (e.g. `var reservedNamesList = ...`) to significantly reduce memory allocations and improve execution speed.
