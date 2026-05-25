## 2024-05-25 - Windows Path Sanitization Bypass
**Vulnerability:** Truncating filenames after checking for Windows reserved names allows a bypass where trailing spaces or dots reveal a reserved name (e.g., `CON   `), which Windows ignores when validating, and also bypasses the max length check when reserved names are expanded (e.g. `CON_`).
**Learning:** Always perform length bounding/truncation before evaluating semantic meaning, checking for reserved names, or trimming trailing spaces/dots. Truncating after these validation steps can create new string suffixes that bypass the filter logic.
**Prevention:** Move `truncateFilename` to immediately after invalid character removal and before checking for reserved names and trimming trailing characters.
