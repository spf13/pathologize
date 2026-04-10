## 2025-02-15 - Proper Extension Handling for Windows Reserved Names
**Vulnerability:** Path sanitization bypass for Windows reserved names (e.g., CON) when multiple extensions are present (e.g., CON.tar.gz).
**Learning:** `filepath.Ext()` only extracts the final extension (e.g., `.gz`), leaving intermediate extensions (e.g., `.tar`) attached to the base name. This fails to identify the true base name, allowing reserved names with multiple extensions to bypass sanitization.
**Prevention:** When validating base names against reserved words, isolate the true base name by splitting at the first dot (`strings.SplitN(filename, ".", 2)[0]`) and re-attach all extensions using `strings.TrimPrefix`.
