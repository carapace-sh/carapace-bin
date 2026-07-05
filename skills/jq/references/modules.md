# Modules, I/O, Environment, and SQL Operators

jq's module/library system, I/O builtins, environment variable access, and SQL-style operators. These are the advanced integration features beyond pure filter processing.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#advanced-features> (modules and I/O sections) and [SQL-Style Operators](https://jqlang.github.io/jq/manual/#sql-style-operators). For **command-line options** that interact with modules (`-L`, `--slurpfile`, `--rawfile`), see [cli.md](cli.md).

## Environment Variables

### `$ENV`, `env`

| Binding | Description |
|---------|-------------|
| `$ENV` | Object of environment variables as set when jq started |
| `env` | Object representing jq's current environment |

Both produce an object mapping variable names to their string values. There is currently no builtin for **setting** environment variables.

```jq
$ENV.PAGER              # → "less" (or whatever PAGER is set to)
env.HOME                # → "/home/user" (or similar)
```

| Program | Input | Output |
|---------|-------|--------|
| `$ENV.PAGER` | `null` | `"less"` |
| `env.PAGER` | `null` | `"less"` |

### `$JQ_BUILD_CONFIGURATION`

A builtin binding showing the jq executable's build configuration. No particular format — at least the `./configure` command-line arguments. May be enriched in future. Can be overridden with `--arg` and related options.

### `$ARGS`

Contains command-line arguments passed via `--arg`/`--argjson`/`--args`/`--jsonargs`:

| Field | Description |
|-------|-------------|
| `$ARGS.named` | Object of named arguments (`--arg name value` → `{"name": "value"}`) |
| `$ARGS.positional` | Array of positional arguments (`--args` or `--jsonargs`) |

See [cli.md](cli.md) for the command-line options.

## I/O Builtins

jq has minimal I/O support — mostly control over when inputs are read, plus debug output.

### `input`

Outputs one new input. Generally requires `-n` (null input) or the first entity is lost:

```bash
echo 1 2 3 4 | jq '[., input]'    # → [1,2] [3,4]
```

### `inputs`

Outputs all remaining inputs, one by one. Primarily useful for reductions. Requires `-n`:

```bash
echo 1 2 3 | jq -n 'reduce inputs as $i (0; . + $i)'    # → 6
```

### `debug`, `debug(msgs)`

Like `.` (passes input through) but writes debug messages to stderr. The `debug` filter outputs:

```
["DEBUG:",<input-value>]
```

`debug(msgs)` is defined as `(msgs | debug | empty), .` — allowing custom message content:

```jq
1 as $x | 2 | debug("Entering function foo with $x == \($x)", .) | (.+1)
# stderr: ["DEBUG:","Entering function foo with $x == 1"]
# stderr: ["DEBUG:",2]
# output: 3
```

### `stderr`

Prints input to stderr in raw, compact mode with no decoration (not even a newline).

### `input_filename`

Returns the name of the file whose input is currently being filtered. May not work well outside a UTF-8 locale.

### `input_line_number`

Returns the line number of the input currently being filtered.

## Modules

jq has a library/module system. Module files end in `.jq`.

### Module Search Path

Modules are searched for in this order:

1. Paths given to `-L` command-line option (replaces default search list)
2. Default search path: `["~/.jq", "$ORIGIN/../lib/jq", "$ORIGIN/../lib"]`

### Path Substitutions

| Prefix | Substitution |
|--------|-------------|
| `~/` | User's home directory |
| `$ORIGIN/` | Directory where the jq executable is located |
| `./` or `.` | Path of the including file (current directory for command-line programs) |

### Module Resolution

A dependency with relative path `foo/bar` is searched for as:

- `foo/bar.jq`
- `foo/bar/bar.jq`

This allows both single-file modules and modules in directories (with README, version control, etc.).

Consecutive components with the same name are not allowed (e.g., `foo/foo`) to avoid ambiguities.

### `~/.jq` Auto-Loading

If `~/.jq` exists as a file (not a directory), it is automatically sourced into the main program.

### `import RelativePathString as NAME [<metadata>];`

Imports a module. A `.jq` suffix is added to the path. The module's symbols are prefixed with `NAME::`:

```jq
import "foo/bar" as Bar;     # Bar::function_name to access
```

The optional metadata is a constant jq expression (object). jq uses the `search` key (string or array of strings) as a search path prefix. Metadata is available via `modulemeta`.

### `include RelativePathString [<metadata>];`

Imports a module **in place** — symbols are imported into the caller's namespace directly (no `NAME::` prefix):

```jq
include "foo/bar";           # functions available without prefix
```

### `import RelativePathString as $NAME [<metadata>];`

Imports a **JSON file** (`.json` suffix added). The file's data is available as `$NAME::NAME`:

```jq
import "data/config" as $Config;     # $Config::Config holds the parsed JSON
```

### `module <metadata>;`

Optional directive providing metadata readable via `modulemeta`. Must be a constant jq expression (object with keys like `homepage`). jq doesn't use this metadata, but it's available to users.

### `modulemeta`

Takes a module name as input, outputs metadata as an object:

| Key | Value |
|-----|-------|
| `deps` | Array of imports (including their metadata) |
| `defs` | Array of the module's defined functions |

Programs can use this to query module metadata — e.g., to search for, download, and install missing dependencies.

## SQL-Style Operators

jq provides a few SQL-style operators for joining and indexing:

### `INDEX(stream; index_expression)`

Produces an object whose keys are computed by applying `index_expression` to each value from `stream`:

```jq
INDEX(.[]; .id)    # object keyed by .id of each element
```

### `JOIN($idx; stream; idx_expr; join_expr)`

Joins values from `stream` to the given index. Index keys computed by `idx_expr`. An array of `[stream_value, index_value]` is fed to `join_expr`:

```jq
JOIN($idx; .[]; .id; .[1].name)    # join on .id, output matching name
```

### `JOIN($idx; stream; idx_expr)`

Same as `JOIN($idx; stream; idx_expr; .)` — join expression defaults to identity.

### `JOIN($idx; idx_expr)`

Joins the input `.` to the index. `idx_expr` applied to `.` computes the index key:

```jq
JOIN($idx; .id)
```

### `IN(s)`

Outputs `true` if `.` appears in the given stream, otherwise `false`:

```jq
.[] | IN(1, 2, 3)    # true if element is 1, 2, or 3
```

### `IN(source; s)`

Outputs `true` if any value in `source` appears in `s`, otherwise `false`.

## References

- [jq manual — Advanced features](https://jqlang.github.io/jq/manual/#advanced-features) (I/O and modules)
- [jq manual — SQL-Style Operators](https://jqlang.github.io/jq/manual/#sql-style-operators)
- For `-L`, `--slurpfile`, `--rawfile` command-line options, see [cli.md](cli.md)
- For `now` and date functions, see [dates.md](dates.md)
