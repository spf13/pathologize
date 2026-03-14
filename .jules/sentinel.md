## $(date +%Y-%m-%d) - Missing filename truncation in Clean function
**Vulnerability:** The `Clean` function did not enforce a maximum filename length, although a `truncateFilename` helper existed.
**Learning:** Functions that validate and clean paths might overlook length limits, leading to potential DoS or writing errors on filesystems with strict length constraints (e.g., 255 bytes).
**Prevention:** Always ensure length limits are enforced early in the sanitization pipeline.
