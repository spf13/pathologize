## 2026-04-20 - Fix Windows Reserved Names Path Sanitization Bypass
**Vulnerability:** Path sanitization logic bypassed when handling Windows reserved names with multiple extensions (e.g., `CON.tar.gz`). `filepath.Ext` only strips the last extension, leaving the inner extensions as part of the base name, failing the reserved name check.
**Learning:** Windows treats any extension on a DOS reserved name as the reserved device itself. Path sanitization for Windows reserved names must consider the base name without *any* extensions.
**Prevention:** When validating Windows reserved names, extract the true base name by splitting at the first dot (e.g., `strings.SplitN(filename, ".", 2)[0]`) and preserve the full extension set.
