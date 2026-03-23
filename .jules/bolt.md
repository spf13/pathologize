## 2024-05-24 - Extract dynamic operations from critical path
**Learning:** `strings.Split` and `regexp.MustCompile` inside simple functions executed often (like `Clean`) significantly block execution and waste memory with unneeded allocations. Go creates standard arrays inside a function each time instead of retaining them.
**Action:** Extract any static sets like reserved filenames and common regex patterns from function bodies into package-level variables.
