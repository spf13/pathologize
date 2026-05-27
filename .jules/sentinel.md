## 2024-05-24 - Path Sanitization Order of Operations
**Vulnerability:** Path sanitization bypass via length truncation.
**Learning:** If truncation happens after semantic checks (like stripping trailing spaces/dots or checking reserved names), an attacker can supply an overlong string that, when truncated, results in a restricted pattern (like "CON" with trailing spaces, which Windows resolves to the restricted "CON" device).
**Prevention:** Always perform length bounding/truncation before evaluating semantic meaning or trimming.
