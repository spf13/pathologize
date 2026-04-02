## 2024-05-18 - Path Sanitization Bug in Pathologize

**Vulnerability:** The pathologize library did not correctly handle path sanitization for Windows reserved names that contained multiple extensions (e.g., `CON.tar.gz`). `filepath.Ext()` only extracts the final extension (e.g. `.gz`), meaning it evaluated the base name as `CON.tar`, missing the fact that the underlying filename is `CON` which is an invalid reserved name.

**Learning:** When validating Windows reserved names (e.g., 'CON.tar.gz'), do not use `filepath.Ext()` to find the base name, as it only extracts the final extension. Instead, extract the true base name by splitting at the first dot. The implementation of pathologize attempts to find the "base name" without any extensions, evaluate it for invalid characters, and append back the extension.

**Prevention:** Always extract the true base name by splitting at the first dot (e.g., `strings.SplitN(filename, ".", 2)[0]`) and preserve the full extension set using `strings.TrimPrefix(filename, basefilename)`.
