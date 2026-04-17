## 2024-04-17 - Windows Reserved Name Evasion via Multiple Extensions
**Vulnerability:** The path sanitizer incorrectly used `filepath.Ext(filename)` to extract the base name of a file to check against Windows reserved names (e.g., `CON`). `filepath.Ext()` only extracts the final extension (e.g., `.gz` in `CON.tar.gz`), meaning the check was performed against `CON.tar`, allowing attackers to evade the sanitizer by using multiple extensions to access forbidden device files.
**Learning:** `filepath.Ext()` is unsafe for base name extraction when malicious actors can chain extensions.
**Prevention:** Split the filename at the first dot using `strings.SplitN(filename, ".", 2)[0]` to extract the true base name, and preserve all extensions when reconstructing the sanitized filename.
