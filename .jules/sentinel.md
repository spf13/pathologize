## 2024-03-24 - Unused Length Truncation Vulnerability
**Vulnerability:** The `pathologize` package defines a `truncateFilename` function and tests it but never actually calls it in the main `Clean` function. This leaves the library vulnerable to processing overly long strings which could lead to length-based DoS or file-system constraint errors.
**Learning:** Even when security functions are written and tested, they can be accidentally omitted from the main execution path. Always verify that security controls are actively enforced.
**Prevention:** Ensure integration tests cover boundary conditions (like max length) on the primary public API (`Clean`), not just testing internal helper functions in isolation.
