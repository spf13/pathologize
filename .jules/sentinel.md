## 2024-04-08 - Windows Path Sanitization Bypass via Multiple Extensions

**Vulnerability:** The application attempts to sanitize Windows reserved names (e.g. `CON`, `PRN`), but the validation could be bypassed by appending multiple extensions, such as `CON.tar.gz`. The application used `filepath.Ext(filename)` which only stripped the last extension (`.gz`), resulting in `CON.tar` which does not match `CON`.

**Learning:** On Windows, file extensions are ignored when evaluating if a base name is a reserved device name (DOS legacy). Therefore, `CON`, `CON.txt`, `CON.tar.gz` are all treated identically as the underlying `CON` device. Path sanitization logic must look at the true base name (before the very first dot), not just the text before the final extension.

**Prevention:** When validating Windows reserved names, use `strings.SplitN(filename, ".", 2)[0]` to extract the true base name instead of relying on `filepath.Ext()`. After sanitizing, restore the full extension substring using `strings.TrimPrefix(filename, basefilename)` rather than appending only the final extension.
