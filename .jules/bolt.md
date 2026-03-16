## 2024-05-24 - Pre-computing and fast-pathing in Go

**Learning:** Compiling regexes (`regexp.MustCompile`) inside functions that are called frequently or in loops causes massive allocation overhead in Go. `strings.Split` and string concatenation inside loops do the same. Checking if a string *needs* modification with `strings.ContainsAny` is substantially faster than blindly passing it to a regex replacer.
**Action:** Always move regex compilation and static string splitting/manipulation to package-level vars or `init()` blocks. Use fast string scanning (`strings.ContainsAny`, `strings.IndexByte`) to bypass complex replacement logic for strings that don't need changes.
