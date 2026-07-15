# Filters

Pandoc's filter system: JSON filters, Lua filters, citeproc, the filter pipeline, and running pandoc as a Lua interpreter.

> **Source of truth**: <https://pandoc.org/MANUAL.html> and <https://pandoc.org/filters.html>. For the internal AST and Walk typeclass, see [architecture.md](architecture.md).

## Overview

Filters sit between the reader and the writer, transforming the AST:

```
INPUT --reader--> AST --filter--> AST --writer--> OUTPUT
```

Filters are applied sequentially in the order specified on the command line. Multiple filters can be chained:

```bash
pandoc --lua-filter=filter1.lua --lua-filter=filter2.lua --filter=json_filter.py input.md -o output.html
```

## Filter Types

| Filter Type | CLI Flag | Implementation | Data Flow |
|-------------|----------|-----------------|-----------|
| **JSON Filter** | `--filter PROGRAM` | External executable | AST serialized to JSON on stdin, deserialized from stdout |
| **Lua Filter** | `--lua-filter SCRIPT` | Embedded Lua interpreter (`pandoc-lua-engine`) | Direct AST manipulation, no serialization overhead |
| **Citeproc** | `--citeproc` | Built-in Haskell module | Internal AST transformation via `processCitations` |

### JSON Filters

JSON filters are external programs that read the Pandoc AST as JSON from stdin and write a transformed AST as JSON to stdout. Any programming language can be used.

```bash
pandoc --filter pandoc-citeproc input.md -o output.html
pandoc --filter ./my-filter.py input.md -o output.html
```

The program receives JSON on stdin and must output valid Pandoc JSON on stdout. The filter runs once per invocation, processing the entire document at once.

**Example Python filter** (requires `pandocfilters` package):

```python
#!/usr/bin/env python3
from pandocfilters import toJSONFilter, Str, Emph

def caps(key, value, format, meta):
    if key == 'Str':
        return Str(value.upper())

if __name__ == "__main__":
    toJSONFilter(caps)
```

### Lua Filters

Lua filters use pandoc's embedded Lua interpreter, avoiding serialization overhead. They define functions that match AST element names.

```bash
pandoc --lua-filter=behead.lua input.md -o output.html
```

**Example Lua filter** (converts headers to bold text):

```lua
function Header(el)
    return pandoc.Para{pandoc.Strong(el.content)}
end
```

Lua filters have direct access to the pandoc AST through the `pandoc` Lua module. Key functions include:

| Function | Trigger |
|----------|---------|
| `Pandoc(doc)` | Called once on the full document |
| `Header(el)` | Called on each header element |
| `Para(el)` | Called on each paragraph |
| `CodeBlock(el)` | Called on each code block |
| `Link(el)` | Called on each link |
| `Image(el)` | Called on each image |
| `Math(el)` | Called on each math element |
| `Div(el)` | Called on each Div |
| `Span(el)` | Called on each Span |
| `Table(el)` | Called on each table |
| `Cite(el)` | Called on each citation |
| `Meta(meta)` | Called on metadata |

Each function receives the AST element and returns either a modified element, a list of elements, or `nil` (no change). Functions are applied in a specific order: `Pandoc` first, then per-element traversal (bottom-up), then `Meta`.

The `pandoc` Lua module provides constructors for all AST types (e.g., `pandoc.Header`, `pandoc.Para`, `pandoc.Str`, `pandoc.Emph`) and utility functions like `pandoc.walk` for recursive traversal.

### Citeproc

The `--citeproc` flag activates the built-in citation processor. It transforms citation elements (`Cite`) into rendered citations and adds a bibliography section. See [citations.md](citations.md) for details.

```bash
pandoc --citeproc --bibliography=refs.bib --csl=chicago.csl input.md -o output.html
```

## Filter Pipeline

Internally, filters are orchestrated by `Text.Pandoc.Filter.applyFilters`. The pipeline applies each filter sequentially using a monadic `foldM` over the filter list:

```
AST₀ → filter₁ → AST₁ → filter₂ → AST₂ → ... → ASTₙ → writer
```

After all external/Lua filters run, pandoc applies **internal transforms**:

| Transform | Effect |
|-----------|--------|
| `headerShift` | Applies `--shift-heading-level-by` |
| `eastAsianLineBreakFilter` | Handles East Asian line breaks |
| `filterToBeamer` | Adjusts for Beamer slides |
| `flatten` | Flattens nested constructs |

## Filter Ordering

The order of `--filter`, `--lua-filter`, and `--citeproc` flags determines the execution order. In defaults files, the `filters` list order determines execution order.

```bash
# citeproc runs after the Lua filter
pandoc --lua-filter=normalize.lua --citeproc input.md -o output.pdf
```

## Running Pandoc as a Lua Interpreter

Pandoc can be invoked as a standalone Lua interpreter:

```bash
pandoc-lua script.lua
pandoc lua script.lua
```

In this mode, the following are available:

| Global | Description |
|--------|-------------|
| `pandoc.*` | All pandoc Lua modules (AST constructors, utils, types) |
| `re` | Regular expression library |
| `lpeg` | LPeg parsing library |
| `PANDOC_VERSION` | Pandoc version string |
| `PANDOC_STATE` | Pandoc runtime state |
| `PANDOC_API_VERSION` | API version |

This mode is useful for running standalone scripts that manipulate pandoc documents without doing a full conversion.

## Custom Readers and Writers (Lua)

### Custom Lua Readers

A Lua script defining a `Reader(string)` function can be used as a custom reader. The function receives the raw input text and must return a Pandoc AST.

```lua
function Reader(input)
    local blocks = {}
    for line in input:lines() do
        table.insert(blocks, pandoc.Para{pandoc.Str(line)})
    end
    return pandoc.Pandoc(blocks)
end
```

Custom readers are sought in the `custom` subdirectory of the data directory. The `lpeg` parsing library is available by default.

### Custom Lua Writers

A custom Lua writer is a Lua script defining rendering functions for each AST element type. There is no default template for custom writers; one must be provided via `--template`.

Custom writers define functions like `Str(s)`, `Para(s)`, `Header(level, s, attr)`, etc., each returning a string fragment.

## Security Considerations

1. **Filters can perform arbitrary file operations**: JSON filters are external programs with full system access. Lua filters have access to the file system via `io` and `os` modules.
2. **`--sandbox`** limits file system and network access for Lua filters.
3. **Use `PandocPure` monad** in Haskell library mode for full isolation (no I/O at all).
4. **Parsers can exhibit pathological performance** on untrusted input. Use `commonmark` reader for untrusted Markdown. Use timeouts.
5. **HTML output is not guaranteed safe** (XSS). Sanitize before serving.

For the full security discussion, see [architecture.md](architecture.md).

## Related References

- [architecture.md](architecture.md) — AST types, Walk typeclass, PandocMonad, filter modules
- [citations.md](citations.md) — Citeproc and citation processing
- [cli-options.md](cli-options.md) — `--filter`, `--lua-filter`, `--citeproc` flags
- [defaults.md](defaults.md) — Specifying filters in defaults files
