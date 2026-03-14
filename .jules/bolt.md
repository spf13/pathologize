
## 2024-05-28 - Precomputing RegExp and Static Slices
**Learning:** Found an anti-pattern in `pathological.go` where static structures (RegExp, constant string splits) were defined inside high-frequency functions. Go's `regexp.MustCompile` can add significant overhead, as does repeatedly allocating array slices from static strings.
**Action:** When performing string modifications or validation checks inside hot-path functions, extract static resources (`regexp.MustCompile` results, constant maps, or static slices) into package-level variables initialized in `var` blocks or `init()` functions to avoid per-call allocation and setup overhead.
