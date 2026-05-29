## 2024-05-29 - Fixed Windows Reserved Name Bypass via Order of Operations
**Vulnerability:** A payload structured to trigger length truncation at the very end of path sanitization can bypass Windows reserved name filters. If a string like `CON` + 252 spaces + `x` is evaluated, it doesn't match `CON` because of the trailing `x`. When truncated, the `x` drops, leaving `CON` with trailing spaces.
**Learning:** Path sanitization logic must prioritize length bounding and truncation *before* evaluating semantic meaning or reserved names. Truncating after validation can inadvertently create string suffixes that were already validated against.
**Prevention:** Always perform truncation or length limits before stripping characters or checking for reserved device names.
