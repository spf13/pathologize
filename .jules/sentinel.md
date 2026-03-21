## 2024-05-18 - [Missing Path Length Enforcement]
**Vulnerability:** The library claims to restrict paths to a safe length (255 characters), but the internal `truncateFilename` function was never actually called in the primary `Clean` function. This led to potentially unconstrained filename lengths.
**Learning:** Security controls explicitly documented as features must be directly validated in testing to ensure they are invoked.
**Prevention:** Ensure integration or boundary tests invoke functions through their primary public API to catch when an internal enforcement function isn't called.
