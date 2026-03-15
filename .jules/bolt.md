## 2024-05-24 - Regexp Compilation Inside Functions
**Learning:** In Go, calling `regexp.MustCompile` inside a function compiles the regular expression on every invocation, causing significant CPU and memory overhead.
**Action:** Always compile regular expressions as package-level variables using `var regex = regexp.MustCompile(...)`, or replace them with simpler string/rune scanning functions like `strings.TrimRightFunc` when applicable for better performance.
