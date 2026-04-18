## 2024-05-18 - Path Sanitization Vulnerability with Multiple Extensions
**Vulnerability:** Windows reserved names (e.g., CON, PRN, AUX) could be bypassed if a file had multiple extensions (e.g., `CON.tar.gz`).
**Learning:** `filepath.Ext` only strips the final extension, which is insufficient for Windows validation, as Windows evaluates everything up to the *first* period as the base name. A file named `CON.tar.gz` would be evaluated as `CON.tar` by the code, passing the reserved name check, but the underlying OS would still treat it as the `CON` device.
**Prevention:** When evaluating base names for device driver reservations on Windows, split the filename string at the *first* period (e.g., `strings.SplitN(filename, ".", 2)[0]`) rather than using `filepath.Ext`.
