## 2024-05-30 - Regex Bottleneck in Path Sanitization
**Learning:** In highly called sanitization functions like `removeInvalidCharacters`, using `regexp.ReplaceAllString` introduces significant allocation and CPU overhead.
**Action:** Pair `strings.IndexFunc` for a zero-allocation fast path with `strings.Map` for replacement to dramatically reduce CPU cycles and heap allocations on mostly clean inputs.
