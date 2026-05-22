## 2024-05-22 - Order of Operations in Path Sanitization
**Vulnerability:** Truncating length AFTER semantic evaluation bypassed reserved name filters. A malicious string like `CON` + 252 spaces + `a` would pass reserved checks (since it's not `CON`) but get truncated to exactly 255 bytes, ending in spaces. Windows OS would later trim these trailing spaces, treating the result as the reserved device `CON`, bypassing sanitization.
**Learning:** Path sanitization must always perform structural changes like length bounding/truncation BEFORE evaluating semantic meaning, checking for reserved names, or trimming spaces/dots.
**Prevention:** Always perform length bounding/truncation before evaluating semantic meaning, checking for reserved names, or trimming trailing spaces/dots.
