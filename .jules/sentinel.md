## 2024-04-04 - Fix path sanitization bypass for Windows reserved names
**Vulnerability:** Path sanitization bypass on Windows reserved names when multiple extensions are used (e.g., `CON.tar.gz`).
**Learning:** `filepath.Ext()` only extracts the final extension, so `removeReservedWithExtension` would check `CON.tar` against reserved names (which is not reserved), allowing `CON.tar.gz` to bypass sanitization. Windows treats any extension on a reserved name as the reserved device itself.
**Prevention:** Extract the true base name by splitting at the first dot (`strings.SplitN(filename, ".", 2)[0]`) and append the remaining suffix (`strings.TrimPrefix(filename, basefilename)`) if the base name is modified.
