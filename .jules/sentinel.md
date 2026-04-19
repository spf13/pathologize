## 2024-04-19 - Path Sanitization Bypass on Windows
**Vulnerability:** Windows path validation bypass where reserved names with multiple extensions (e.g., `CON.tar.gz`) were not correctly identified and neutralized.
**Learning:** `filepath.Ext` only extracts the final extension. Windows treats any extension on a reserved DOS device name as the device itself, so `CON.tar.gz` resolves to `CON`. The validation must extract the true base name before the first dot.
**Prevention:** When validating against Windows reserved names, extract the base name using `strings.SplitN(filename, ".", 2)[0]` instead of relying on `filepath.Ext` and `TrimSuffix`.
