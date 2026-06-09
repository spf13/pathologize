## 2024-05-24 - Path Sanitization Order of Operations
**Vulnerability:** Windows reserved name bypass via length truncation.
**Learning:** Truncating after semantic validation steps can create new string suffixes (like trailing spaces revealing a hidden Windows reserved device name) that bypass the filter logic.
**Prevention:** Always perform length bounding/truncation before evaluating semantic meaning, checking for reserved names, or trimming trailing spaces/dots.
