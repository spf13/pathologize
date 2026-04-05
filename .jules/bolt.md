## 2024-05-24 - Avoid Inline Regex Compilation and Repeated String Allocations in Go
**Learning:** Found two common performance anti-patterns in Go:
1) Compiling regular expressions inline inside frequently called functions (e.g., `regexp.MustCompile` in `removeTrailing`) which introduces significant performance bottlenecks and unnecessary heap allocations on every single call.
2) Repeatedly executing `strings.Split` on static constants inside frequently called functions (e.g., in `removeReservedNames`), causing repeated memory allocations.

**Action:**
1) Always replace inline `regexp.MustCompile` with native `strings` functions like `strings.TrimRight` where applicable (e.g., replacing `re.ReplaceAllString` for simple trailing character removal with `strings.TrimRight`), which reduces allocations to zero and significantly improves speed.
2) Extract repeatedly calculated values from static/constant strings into package-level variables initialized once. For simple character class matchings (like `[.\s]+$`), check exactly what character sets they match in Go (`\s` corresponds to `[\t\n\f\r ]` in the `regexp` package) to ensure exact behavioral parity when switching to `strings` functions.
