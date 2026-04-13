## 2024-04-13 - Path Sanitization Bypass for Windows Reserved Names
**Vulnerability:** Windows reserved names with multiple extensions (e.g., CON.tar.gz) bypassed sanitization because filepath.Ext only extracts the final extension, causing the base name check to fail.
**Learning:** When checking for DOS/Windows reserved names, base name extraction must consider everything before the FIRST dot, not the last dot.
**Prevention:** Use strings.SplitN(filename, ".", 2) to correctly isolate the base name from any arbitrary sequence of extensions.
