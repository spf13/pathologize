## 2024-05-27 - Zero-Allocation Fast Paths
**Learning:** For string character filtering and truncation, `strings.Map` combined with `strings.IndexFunc`, and `strings.IndexByte` drastically outperform compiled regex (`regexp.ReplaceAllString`) and array allocation (`strings.SplitN`).
**Action:** Use `IndexFunc` + `Map` for char filtering, and `IndexByte` for simple delimiter splitting instead of `regexp` and `SplitN` in hot paths.
