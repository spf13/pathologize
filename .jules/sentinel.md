## 2024-05-28 - Missing File Name Truncation Limit
**Vulnerability:** The `pathologize.Clean` and `pathologize.CleanPath` functions did not restrict the maximum length of filenames or paths, despite having a `truncateFilename` function defined in the codebase.
**Learning:** This could lead to Denial of Service (DoS) risks or file system errors when working with file systems that cannot handle extremely long file names. It highlights a surprising security gap where a defense mechanism was defined but not utilized in the main cleaning logic.
**Prevention:** Ensure that input length limits are consistently applied to all paths, and actively look for unused helper functions that might indicate missing security controls.
