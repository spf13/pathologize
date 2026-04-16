## 2024-04-16 - Windows Reserved Names Path Sanitization Bypass
**Vulnerability:** Path sanitization logic (`pathologize.Clean`) bypassed when given a reserved filename with multiple extensions (e.g., `CON.tar.gz`).
**Learning:** `filepath.Ext()` only extracts the final extension. Windows treats any extension on a reserved base name (e.g., `CON`) as the device itself, so `CON.tar.gz` and `CON.txt.txt` are both reserved and bypass a check that relies solely on `filepath.Ext()`.
**Prevention:** When validating Windows reserved names, extract the true base name by splitting at the first dot (`strings.SplitN(filename, ".", 2)[0]`) and preserve the full extension suffix (`strings.TrimPrefix(filename, basefilename)`) instead of relying on `filepath.Ext()`.
