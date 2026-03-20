## 2024-11-06 - Missing Filename Truncation Enforcement
**Vulnerability:** The library implemented a truncation function for safety but failed to invoke it in the main `Clean` path, potentially allowing bypass of intended length limits (DoS risk).
**Learning:** Unused security functions are common when features are implemented but the main entrypoint is missed.
**Prevention:** Always write integration tests that hit the main entry point to verify constraints like length.
