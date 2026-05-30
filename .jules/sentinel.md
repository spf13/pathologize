## 2024-05-24 - Path Sanitization Bypass via Truncation Order
**Vulnerability:** The `Clean` function performed path sanitization (trimming, reserved name checks) before length truncation. This allowed bypassing reserved name filters by appending spaces and a non-space character to push the non-space character past the 255 limit, leaving trailing spaces after truncation.
**Learning:** Always perform length bounding/truncation before evaluating semantic meaning, checking for reserved names, or trimming trailing spaces/dots. Truncating after validation can create new string suffixes that bypass the filter logic.
**Prevention:** Ensure `truncateFilename` executes immediately after initial character filtering, before higher-level validations and trimmings.
