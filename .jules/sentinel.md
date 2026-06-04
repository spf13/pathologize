## 2025-02-27 - Fix Windows Reserved Name Validation Bypass
**Vulnerability:** Windows reserved name validation bypass via path truncation sequencing.
**Learning:** Path truncation must happen *before* validating against reserved names or removing trailing spaces. If truncation happens last, a string like `"CON" + 252 spaces + "x"` bypasses checks because it isn't "CON" or trailing spaces. When truncated, the "x" drops, leaving trailing spaces which Windows then strips, resolving to the forbidden "CON" device.
**Prevention:** Always perform length bounding/truncation before evaluating semantic meaning, checking for reserved names, or trimming trailing spaces/dots in path sanitization pipelines.
