
## 2024-05-24 - Zero-Allocation Path Cleaning in Go
**Learning:** The `removeInvalidCharacters` function in path sanitization is extremely hot path and gets called on every `Clean` invocation. Regex-based filtering `regexp.MustCompile(...).ReplaceAllString` creates massive overhead (~800ns and 3 allocations just for a clean string).
**Action:** When filtering characters, pair `strings.IndexFunc` (to create a zero-allocation fast path for clean strings) with `strings.Map` (for replacement). Combined with `strings.IndexByte` for extension extraction over `strings.SplitN`, this drops CPU overhead by up to 10x and eliminates most allocations.
