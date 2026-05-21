## 2024-05-21 - Replace RegExp with IndexFunc/Map for fast character filtering
**Learning:** Regular expressions in Go introduce significant CPU overhead and heap allocations for simple character filtering tasks. Replacing `regexp.ReplaceAllString` with a combination of `strings.IndexFunc` (fast path for clean strings) and `strings.Map` provides massive performance gains (e.g. 10x faster) and zero allocations.
**Action:** Prefer `strings.IndexFunc` and `strings.Map` over regular expressions for filtering or replacing characters when performance is critical.
