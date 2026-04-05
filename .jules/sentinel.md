
## 2024-11-20 - Path Sanitization Bypass via Multiple Extensions
**Vulnerability:** Path sanitization bypass on Windows reserved names (e.g., `CON.tar.gz`). `filepath.Ext()` only removes the final extension, so files with multiple extensions bypass the reserved name filter if the remaining extension isn't stripped.
**Learning:** In Go, `filepath.Ext()` only extracts the last extension. When checking for a Windows reserved base name, extracting it using `filepath.Ext()` will miss cases with multiple extensions. Windows processes the entire path differently and strips any extension before identifying the reserved device.
**Prevention:** Always extract the true base name by splitting at the first dot (`strings.SplitN(filename, ".", 2)[0]`) when filtering out Windows/DOS reserved names, and subsequently preserve the entire suffix to form the safe file name.
