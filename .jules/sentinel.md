## 2024-04-09 - Pathologize Path Sanitization

**Vulnerability:** Path sanitization logic relied on `filepath.Ext()` to separate extensions from base names before checking for Windows/DOS reserved names (e.g., `CON`). `filepath.Ext()` only extracts the final extension. Therefore, if a file had multiple extensions (e.g., `CON.tar.gz`), `filepath.Ext()` would only remove `.gz`, leaving `CON.tar` as the base name, which circumvents the check for the `CON` reserved name. The original behavior allowed potentially restricted reserved names to slip through because `CON.tar` != `CON`.

**Learning:** When validating Windows reserved names and extracting base filenames from arbitrary strings, do not rely on standard library functions that extract only the final extension suffix. Windows treats any extension added to a DOS reserved name as the reserved device itself.

**Prevention:** To robustly extract the base name regardless of the number of extensions, split the string at the first period (`strings.SplitN(filename, ".", 2)[0]`). To extract and preserve the entirety of the trailing extensions, remove the discovered base name from the original string as a prefix (`strings.TrimPrefix(filename, basefilename)`).
