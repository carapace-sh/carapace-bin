# Date and Time Functions

jq provides basic date handling — high-level ISO 8601 functions and low-level C library time functions. All date builtins operate in **UTC**.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#dates>.

## High-Level Date Functions

| Function | Description |
|----------|-------------|
| `fromdateiso8601` | Parse ISO 8601 datetime string → seconds since Unix epoch |
| `todateiso8601` | Seconds since epoch → ISO 8601 datetime string |
| `fromdate` | Parse datetime string → seconds since epoch (currently ISO 8601 only) |
| `todate` | Alias for `todateiso8601` |
| `now` | Current time in seconds since Unix epoch |

### `fromdate` / `todate`

```jq
fromdate              # "2015-03-05T23:51:47Z" → 1425599507
todate                # 1425599507 → "2015-03-05T23:51:47Z"
```

| Program | Input | Output |
|---------|-------|--------|
| `fromdate` | `"2015-03-05T23:51:47Z"` | `1425599507` |

`fromdate` currently only supports ISO 8601 format. Future versions may parse more formats.

## Low-Level Time Functions

These interface to the C library time functions:

| Function | Description |
|----------|-------------|
| `gmtime` | Seconds since epoch → "broken down time" array (UTC) |
| `localtime` | Like `gmtime` but using local timezone |
| `mktime` | "Broken down time" array → seconds since epoch |
| `strptime(fmt)` | Parse string with format → "broken down time" array |
| `strftime(fmt)` | Format time (GMT) with format string |
| `strflocaltime(fmt)` | Like `strftime` but using local timezone |

### "Broken Down Time" Array

`gmtime` and `localtime` output an array of 8 numbers:

| Index | Field | Notes |
|-------|-------|-------|
| 0 | Year | |
| 1 | Month | Zero-based (0 = January) |
| 2 | Day of month | One-based |
| 3 | Hour of day | |
| 4 | Minute of hour | |
| 5 | Second of minute | |
| 6 | Day of week | Zero-based (Sunday = 0) |
| 7 | Day of year | Zero-based (January 1 = 0) |

The day-of-week may be wrong on some systems for dates before March 1, 1900 or after December 31, 2099.

### `strptime` and `strftime` Format Strings

Format strings follow C library conventions. The ISO 8601 format is:

```
%Y-%m-%dT%H:%M:%SZ
```

| Program | Input | Output |
|---------|-------|--------|
| `strptime("%Y-%m-%dT%H:%M:%SZ")` | `"2015-03-05T23:51:47Z"` | `[2015,2,5,23,51,47,4,63]` |
| `strptime("%Y-%m-%dT%H:%M:%SZ") | mktime` | `"2015-03-05T23:51:47Z"` | `1425599507` |

### Chaining `strptime` → `mktime`

To parse a date string and get epoch seconds:

```jq
strptime("%Y-%m-%dT%H:%M:%SZ") | mktime
```

### Chaining `gmtime` → `strftime`

To format epoch seconds as a string:

```jq
1425599507 | gmtime | strftime("%Y-%m-%dT%H:%M:%SZ")
# → "2015-03-05T23:51:47Z"
```

## Platform Limitations

- jq may not support all date functionality on some systems
- `%u` and `%j` specifiers for `strptime` are not supported on macOS
- Localization behavior is not necessarily stable

## References

- [jq manual — Dates](https://jqlang.github.io/jq/manual/#dates)
- For `now` (current time), see this page
- For environment access (`$ENV`, `env`), see [modules.md](modules.md)
