# Defaults Files

Pandoc defaults files: YAML or JSON files that specify persistent option sets, invoked with `-d`/`--defaults`.

> **Source of truth**: <https://pandoc.org/MANUAL.html#defaults-files>. For the full CLI option reference, see [cli-options.md](cli-options.md).

## Overview

A defaults file is a YAML or JSON file that maps option names to values. Instead of typing long command lines, you store common option sets in a defaults file and invoke it:

```bash
pandoc -d mydefaults.yaml input.md
```

Defaults files are looked up in the `defaults` subdirectory of the user data directory (`~/.local/share/pandoc/defaults/` on Unix). You can also specify a full path.

```bash
# Looks for ~/.local/share/pandoc/defaults/professional.yaml
pandoc -d professional input.md

# Full path
pandoc -d /path/to/myconfig.yaml input.md
```

## Structure

A defaults file is a YAML map where top-level keys correspond to pandoc option categories. Each key maps to a sub-map of option names and values.

```yaml
from: markdown
to: html5
standalone: true
toc: true
toc-depth: 3
css:
  - styles/main.css
  - styles/print.css
metadata:
  title: My Document
  author: Jane Doe
variables:
  mainfont: "Source Sans Pro"
  fontsize: 12pt
  geometry:
    - margin=1in
filters:
  - pandoc-citeproc
  - behead.lua
```

## Environment Variable Interpolation

Defaults files support environment variable interpolation using `${...}` syntax:

| Variable | Resolves to |
|----------|------------|
| `${HOME}` | User home directory |
| `${USERDATA}` | Pandoc user data directory |
| `${.}` | Directory containing the defaults file |

```yaml
include-before-body:
  - ${.}/header.html
include-in-header:
  - ${HOME}/.pandoc/templates/custom.sty
```

## Field Mapping

All CLI options can be specified in a defaults file. The mapping follows the option names (without the `--` prefix):

| CLI Flag | Defaults Key | Notes |
|----------|--------------|-------|
| `--from FORMAT` | `from: FORMAT` | |
| `--to FORMAT` | `to: FORMAT` | |
| `--output FILE` | `output: FILE` | |
| `--standalone` | `standalone: true` | Boolean |
| `--template FILE` | `template: FILE` | |
| `--variable KEY=VAL` | `variables: {KEY: VAL}` | Map or list |
| `--metadata KEY=VAL` | `metadata: {KEY: VAL}` | Map or list |
| `--metadata-file FILE` | `metadata-files: [FILE]` | List |
| `--toc` | `toc: true` | Boolean |
| `--toc-depth N` | `toc-depth: N` | |
| `--number-sections` | `number-sections: true` | Boolean |
| `--top-level-division DIV` | `top-level-division: DIV` | |
| `--include-in-header FILE` | `include-in-header: [FILE]` | List |
| `--include-before-body FILE` | `include-before-body: [FILE]` | List |
| `--include-after-body FILE` | `include-after-body: [FILE]` | List |
| `--resource-path PATH` | `resource-path: [PATH]` | List |
| `--css URL` | `css: [URL]` | List |
| `--highlight-style STYLE` | `highlight-style: STYLE` | |
| `--pdf-engine PROG` | `pdf-engine: PROG` | See [pdf.md](pdf.md) |
| `--citeproc` | `citeproc: true` | Boolean |
| `--bibliography FILE` | `bibliography: [FILE]` | List |
| `--csl FILE` | `csl: FILE` | |
| `--shift-heading-level-by N` | `shift-heading-level-by: N` | |
| `--file-scope` | `file-scope: true` | Boolean |
| `--wrap MODE` | `wrap: MODE` | |
| `--columns N` | `columns: N` | |
| `--eol MODE` | `eol: MODE` | |
| `--embed-resources` | `embed-resources: true` | Boolean |
| `--reference-doc FILE` | `reference-doc: FILE` | |
| `--slide-level N` | `slide-level: N` | |
| `--split-level N` | `split-level: N` | |
| `--epub-cover-image FILE` | `epub-cover-image: FILE` | |
| `--epub-embed-font FILE` | `epub-embed-font: [FILE]` | List |
| `--epub-metadata FILE` | `epub-metadata: FILE` | |
| `--ipynb-output MODE` | `ipynb-output: MODE` | |
| `--track-changes MODE` | `track-changes: MODE` | |
| `--extract-media DIR` | `extract-media: DIR` | |
| `--tab-stop N` | `tab-stop: N` | |
| `--ascii` | `ascii: true` | Boolean |
| `--fail-if-warnings` | `fail-if-warnings: true` | Boolean |
| `--sandbox` | `sandbox: true` | Boolean |
| `--verbose` | `verbose: true` | Boolean |
| `--quiet` | `quiet: true` | Boolean |
| `--log FILE` | `log: FILE` | |
| `--request-header HEADER` | `request-headers: [HEADER]` | List |
| `--no-check-certificate` | `no-check-certificate: true` | Boolean |

## Filters in Defaults Files

The `filters` key specifies a list of filters. Pandoc determines the type from the file extension:

| Extension | Filter Type |
|-----------|-------------|
| `.lua` | Lua filter (uses `--lua-filter` semantics) |
| Other | JSON filter (uses `--filter` semantics) |
| `citeproc` | Built-in citeproc |

```yaml
filters:
  - pandoc-citeproc        # JSON filter (external program)
  - behead.lua             # Lua filter
  - citeproc               # Built-in citeproc
  - ./my-filter.py         # JSON filter (external program)
```

Filters are applied in list order, matching the CLI semantics. For details on filter types, see [filters.md](filters.md).

## Input Files in Defaults

Defaults files can specify input files:

```yaml
input-files:
  - chapter01.md
  - chapter02.md
  - chapter03.md
```

Command-line input files are appended to those from the defaults file.

## Multiple Defaults Files

Multiple `-d` options can be combined; later defaults files override earlier ones:

```bash
pandoc -d base.yaml -d overrides.yaml input.md
```

Values are merged: scalar values are replaced, lists are extended, maps are merged.

## Related References

- [cli-options.md](cli-options.md) — Full CLI option reference
- [filters.md](filters.md) — Filter types and pipeline
- [templates.md](templates.md) — Template variables (used in defaults)
- [citations.md](citations.md) — Citation-related options
