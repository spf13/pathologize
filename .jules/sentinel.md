## 2024-03-17 - DoS Vulnerability via Long Filenames
**Vulnerability:** The `Clean` function successfully sanitized invalid characters and reserved names, but failed to call the existing `truncateFilename` helper, allowing excessively long filenames to pass through.
**Learning:** Even when security utilities (like length limiters) are implemented in a codebase, they must be explicitly wired into the primary execution flow. The presence of a test for `truncateFilename` provided a false sense of security that length was being bounded in the main path.
**Prevention:** Always verify that the entry points of a library or API actually invoke the underlying security controls, regardless of the test suite's coverage of those helper functions.
