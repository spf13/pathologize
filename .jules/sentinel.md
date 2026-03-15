## 2023-11-20 - Windows Reserved Names Bypass via Multiple Extensions
**Vulnerability:** Windows reserved device names (e.g., CON, PRN) could bypass sanitization if the filename had multiple extensions (e.g., `CON.tar.gz`).
**Learning:** `filepath.Ext()` only extracts the last extension (`.gz`). This leaves the base name as `CON.tar`, which fails exact string matching against the reserved names list (`CON`). However, Windows still treats `CON.tar.gz` as the `CON` device, leading to potential path traversal or denial of service issues.
**Prevention:** To safely sanitize base filenames against reserved names, always split at the *first* dot instead of relying on standard extension-extraction libraries that look at the *last* dot.
