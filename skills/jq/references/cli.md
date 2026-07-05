# Invoking jq

Command-line options, input/output modes, exit codes, and environment variables. Covers the full CLI surface of jq 1.7.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#invoking-jq>. For **filter language syntax**, see [filters.md](filters.md). For **passing variables into programs**, see the `--arg` / `--argjson` sections below and [variables-functions.md](variables-functions.md).

## Synopsis

```
jq [<options>...] <filter> [<files>...]
```

jq reads a stream of whitespace-separated JSON values from stdin (default) or from one or more files. Each value is passed through the filter independently. Output is a sequence of newline-separated JSON texts, pretty-printed by default.

## Input Mode Options

These control how jq reads input:

| Option | Short | Effect |
|--------|-------|--------|
| `--null-input` | `-n` | Don't read any input. Run the filter once with `null` as input. Useful for calculators or constructing JSON from scratch. |
| `--raw-input` | `-R` | Don't parse input as JSON. Each line of text is passed as a string. With `--slurp`, the entire input becomes one long string. |
| `--slurp` | `-s` | Read the entire input stream into a single array, run the filter once instead of per-value. |
| `--stream` | | Parse input in streaming fashion, emitting `[path, leaf_value]` arrays. For very large inputs. |
| `--stream-errors` | | Like `--stream`, but invalid JSON yields `[error_string, path]` arrays. Implies `--stream`. |
| `--seq` | | Use `application/json-seq` MIME scheme — ASCII RS before each output, ASCII LF after. Input parse errors are ignored (with warning) until next RS. |

### `--stream` Output Format

In stream mode, jq outputs arrays of path and leaf values:

```
"a"          → [[],"a"]
[[],"a",["b"]] → [[0],[]], [[1],"a"], [[2,0],"b"]
```

Useful with `reduce` and `foreach` for incremental processing of large inputs.

### When to Use `-n` (Null Input)

The most common gotcha: when using `input` or `inputs` builtins, you must pass `-n` or the first input value is consumed implicitly by jq's main loop:

```bash
echo 1 2 3 | jq -n 'reduce inputs as $i (0; . + $i)'   # → 6
echo 1 2 3 | jq 'reduce inputs as $i (0; . + $i)'      # → 5 (first value consumed)
```

## Output Formatting Options

| Option | Short | Effect |
|--------|-------|--------|
| `--compact-output` | `-c` | One JSON object per line, no pretty-printing. |
| `--raw-output` | `-r` | If the result is a string, output it directly without JSON quotes. For non-JSON output. |
| `--raw-output0` | | Like `-r` but NUL-separated instead of newline. Error if output contains NUL. |
| `--join-output` | `-j` | Like `-r` but no newline after each output. |
| `--ascii-output` | `-a` | Escape all non-ASCII as `\uXXXX` instead of UTF-8. |
| `--sort-keys` | `-S` | Output object keys in sorted order. |
| `--tab` | | Use one tab per indentation level instead of two spaces. |
| `--indent n` | | Use `n` spaces for indentation (max 7). |
| `--unbuffered` | | Flush output after each JSON object. Useful for slow piped input. |
| `--color-output` | `-C` | Force colored output even when writing to a pipe/file. |
| `--monochrome-output` | `-M` | Disable colored output. |

### Color and `NO_COLOR`

By default jq colors output when writing to a terminal. When the `NO_COLOR` environment variable is set (non-empty), jq disables color by default — use `-C` to re-enable.

Colors are configured via `JQ_COLORS` (see below).

### `--raw-output` vs `--join-output` vs `--raw-output0`

```
echo '"hello"' | jq .              # "hello"     (JSON-quoted, newline)
echo '"hello"' | jq -r .           # hello       (raw, newline)
echo '"hello"' | jq -j .           # hello       (raw, no newline)
echo '"hello"' | jq --raw-output0 .# hello<NUL>  (raw, NUL separator)
```

Use `--raw-output0` when output values may contain newlines and you need reliable parsing downstream.

## Variable Binding Options

These pass values from the command line into the jq program:

| Option | Arguments | Binds | Example |
|--------|-----------|-------|---------|
| `--arg name value` | name, string | `$name` = `"value"` (string) | `jq --arg foo bar -n '$foo'` → `"bar"` |
| `--argjson name JSON` | name, JSON text | `$name` = parsed JSON value | `jq --argjson foo 123 -n '$foo'` → `123` |
| `--slurpfile name file` | name, filename | `$name` = array of all JSON values in file | `jq --slurpfile foo data.json -n '$foo'` |
| `--rawfile name file` | name, filename | `$name` = file contents as string | `jq --rawfile foo text.txt -n '$foo'` |

### `--arg` vs `--argjson`

`--arg` always binds a **string**, even if the value looks numeric:

```bash
jq --arg foo 123 -n '$foo'        # "123" (string)
jq --argjson foo 123 -n '$foo'    # 123   (number)
```

### `$ARGS`

Named arguments (`--arg`, `--argjson`) are also available as `$ARGS.named` (an object). Positional arguments (`--args` or `--jsonargs`) are available as `$ARGS.positional[]` (an array).

## Positional Argument Options

| Option | Effect |
|--------|--------|
| `--args` | Remaining arguments are positional **string** arguments, not files. Available as `$ARGS.positional[]`. |
| `--jsonargs` | Remaining arguments are positional **JSON** arguments, not files. Available as `$ARGS.positional[]`. |
| `--` | Terminates argument processing. Remaining arguments are positional (strings, JSON, or filenames depending on above). |

## Program Source Options

| Option | Short | Effect |
|--------|-------|--------|
| `--from-file filename` | `-f` | Read the filter from the file instead of the command line. Supports `#` comments. |
| `-L directory` | | Prepend `directory` to the module search path. Replaces the default search list. |

## Execution Control Options

| Option | Short | Effect |
|--------|-------|--------|
| `--exit-status` | `-e` | Exit 0 if last output is neither `false` nor `null`; 1 if it is; 4 if no valid result. |
| `--binary` | `-b` | Windows (WSL/MSYS2/Cygwin) with native jq.exe: prevents LF→CRLF conversion. |
| `--run-tests [filename]` | | Run tests from file or stdin. Must be last option. Format is non-stable. |

## Information Options

| Option | Short | Effect |
|--------|-------|--------|
| `--version` | `-V` | Print jq version and exit 0. |
| `--build-configuration` | | Print build configuration and exit 0. No stable format. |
| `--help` | `-h` | Print help and exit 0. |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Program ran successfully (normal output) |
| 1 | With `-e`: last output was `false` or `null` |
| 2 | Usage problem or system error |
| 3 | jq program compile error |
| 4 | With `-e`: no valid result was ever produced |
| 5 | Default exit code for `halt_error` |

`halt_error(exit_code)` can set a custom exit code (default 5).

## Environment Variables

| Variable | Effect |
|----------|--------|
| `JQ_COLORS` | Configures output colors (see below). |
| `NO_COLOR` | If non-empty, disables colored output by default. Override with `-C`. |

### `JQ_COLORS`

A colon-separated list of ANSI escape sequences for different JSON value types, in this order:

```
null  false  true  numbers  strings  arrays  objects  object keys
```

Each entry is an ANSI escape code. Default (on most systems):

```
JQ_COLORS="1;30:0;37:0;37:0;37:0;32:1;37:1;37:1;34"
```

| Position | Default | Represents |
|----------|---------|------------|
| 1 | `1;30` | `null` |
| 2 | `0;37` | `false` |
| 3 | `0;37` | `true` |
| 4 | `0;37` | numbers |
| 5 | `0;32` | strings |
| 6 | `1;37` | arrays |
| 7 | `1;37` | objects |
| 8 | `1;34` | object keys |

## Shell Quoting

jq programs contain many characters that are shell metacharacters (`|`, `.`, `[]`, `{}`, `()`, `?`, `//`). Always quote the filter:

| Shell | Quoting | Example |
|-------|---------|---------|
| Unix shells (bash, zsh) | Single quotes | `jq '.["foo"]'` |
| PowerShell | Single quotes, escape `"` inside as `\"` | `jq '.[\"foo\"]'` |
| Windows cmd.exe | Double quotes, escape `"` inside as `\"` | `jq ".[\"foo\"]"` |

The most common mistake: `jq "foo"` fails on Unix shells because the shell strips the quotes and jq sees `foo` (an undefined function). Use `jq '.foo'` or `jq '"foo"'` (the latter produces the string "foo").

## Input from Multiple Files

When multiple files are given, jq reads them in order as a single concatenated stream:

```bash
jq '.user' file1.json file2.json    # processes each value from both files
jq -s '.' file1.json file2.json     # slurps all values from both into one array
```

With `--raw-input`, each line from each file is a separate string input.

## `--run-tests` Format

The test file format (non-stable, may change):

- Comment lines (starting with `#`) and empty lines are ignored
- Program line, then one input line, then expected output lines (one per output), then a blank line
- Compilation failure tests: line `%%FAIL`, then program, then expected error message

## References

- [jq manual — Invoking jq](https://jqlang.github.io/jq/manual/#invoking-jq)
- For variable usage inside programs, see [variables-functions.md](variables-functions.md)
- For the `input`/`inputs` I/O builtins, see [modules.md](modules.md)
