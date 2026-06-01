
## 2024-05-24 - Order of Operations Bypass in Path Sanitization
**Vulnerability:** Path sanitization validation (reserved names) could be bypassed because truncation happened after the reserved name checks. A path like "CON" + 252 spaces + "x" would pass the `removeReservedNames` check, but then get truncated to "CON" + 252 spaces, which Windows still treats as the reserved device "CON".
**Learning:** String length bounding and truncation must occur before evaluating semantic meaning, checking for reserved words, or trimming trailing characters. Truncating later can create new string suffixes that bypass earlier filter logic.
**Prevention:** Always perform length bounding/truncation as one of the first steps in a sanitization pipeline, immediately after removing invalid characters.
