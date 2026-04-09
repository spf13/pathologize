## 2024-05-24 - Inline Regex and Repeated Allocations
**Learning:** Compiling regular expressions inline within frequently called functions (e.g., using `regexp.MustCompile`) introduces massive performance bottlenecks. Repeatedly running operations like `strings.Split` on constant strings inside loops also creates unnecessary heap allocations.
**Action:** Pre-compile regular expressions as package-level variables or replace them with native `strings` functions. Extract repeated operations on constants into package-level variables.
