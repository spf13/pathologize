## 2024-04-14 - Windows Reserved Name Path Traversal Bypass via Compound Extensions
**Vulnerability:** Path sanitization bypass for Windows reserved names (e.g., `CON.tar.gz`) due to `filepath.Ext()` only removing the final extension. This allowed malicious files to bypass validation and hit device handles.
**Learning:** `filepath.Ext()` is insufficient for correctly validating base names against restricted patterns when dealing with multi-dotted compound extensions.
**Prevention:** Split filenames on the first dot (`strings.SplitN(filename, ".", 2)`) to reliably extract the true base name when performing reserved name validation.
