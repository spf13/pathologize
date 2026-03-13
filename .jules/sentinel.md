## 2024-05-18 - Missing Length Restriction on Paths
**Vulnerability:** Path truncation was defined but not utilized in `Clean` function
**Learning:** Functions designed to restrict potentially unbounded inputs (like paths) can be inadvertently bypassed if not connected to the main execution path. This represents a minor DoS or application crash risk if very long paths are encountered on restricted filesystems.
**Prevention:** Ensure all input sanitization steps are correctly pipelined in the main entrypoints.
