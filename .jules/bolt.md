## 2024-05-19 - Expensive operations in go pathologize
**Learning:** `regexp.MustCompile` and `strings.Split` within repetitive string manipulations were identified as huge memory bottlenecks causing thousands of unnecessary allocations.
**Action:** Avoid repetitive inline instantiations or regex use for simple string trims. Use `strings.TrimRightFunc` for whitespace/trailing characters and move static computations like array construction to global package scope initialized variables.
