## 2024-05-24 - Pre-compile Regex and Arrays for Validation Lookups
**Learning:** Initializing regular expressions and splitting string lists repeatedly inside loop bodies or frequently called functions (like `removeTrailing` and `removeReservedNames`) causes measurable performance bottlenecks due to repeated parsing and allocations.
**Action:** Always extract static regular expressions and lists to package-level variables and parse them at initialization rather than inside the function body.
