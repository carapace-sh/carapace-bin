# CLI Parameter Types

How Python type annotations map to CLI parsing, conversion, and validation.

> **Source of truth**: <https://typer.tiangolo.com/tutorial/parameter-types/>. For declaration syntax (Annotated etc.), see [parameters.md](parameters.md).

## Conversion Overview

| Annotation | CLI behavior |
|------------|--------------|
| `str` | Plain string (default) |
| `int`, `float` | Numeric conversion with validation errors |
| `bool` | Flag (`--flag/--no-flag`) for options |
| `Enum` subclass | Choices with validation |
| `Literal["a", "b"]` | Choices without a class |
| `pathlib.Path` | Path with validation flags |
| `typer.FileText` etc. | File handles |
| `datetime` | Parsed date/time |
| `uuid.UUID` | Parsed UUID |
| `list[X]` | Multiple values (repeated option or variadic argument) |
| `tuple[X, Y]` | Fixed-size tuple values |

Invalid input produces a friendly validation error and non-zero exit before your code runs.

## Numbers

```python
def main(
    age: Annotated[int, typer.Option(min=0, max=120)] = 0,
    level: Annotated[float, typer.Option(clamp=True)] = 0,
    verbosity: Annotated[int, typer.Option(count=True)] = 0,
):
    ...
```

- `min`/`max` reject out-of-range values (use `Annotated[int, typer.Option(min=..., max=...)]`).
- `clamp=True` silently clamps to the range instead of erroring.
- `count=True` turns an int option into a counter (`-vvv` → 3).

## Enum and Literal Choices

```python
from enum import Enum
from typing import Literal

class Network(str, Enum):
    simplex = "simplex"
    full = "full"

def main(
    network: Annotated[Network, typer.Option(case_sensitive=False)] = Network.simplex,
    mode: Annotated[Literal["fast", "slow"], typer.Argument()] = "fast",
):
    ...
```

- Subclass `str, Enum` to get better editor support and string values.
- `case_sensitive=False` allows `FULL`, `Simplex`, etc.
- `list[Network]` accepts multiple choice values.
- `Literal` gives the same choices UI without declaring a class.
- Invalid values show the valid choices in the error.

## Path

`pathlib.Path` parameters accept validation flags:

```python
import typer
from pathlib import Path

def main(
    config: Annotated[Path, typer.Option(exists=True, file_okay=True, dir_okay=False, readable=True)],
    output: Annotated[Path, typer.Option(writable=True, resolve_path=True)] = Path("."),
):
    ...
```

Key flags: `exists`, `file_okay`, `dir_okay`, `writable`, `readable`, `executable`, `resolve_path`, `allow_dash` (treat `-` as stdin/stdout), `path_type` (e.g. `path_type=typer.FileText` to open the path).

## Files

Typer provides four file markers that produce open file-like objects:

| Type | Purpose |
|------|---------|
| `typer.FileText` | Read text (`mode="r"`) |
| `typer.FileTextWrite` | Write text (`mode="w"`) |
| `typer.FileBinaryRead` | Read binary |
| `typer.FileBinaryWrite` | Write binary |

```python
import typer

def main(infile: typer.FileText, outfile: typer.FileTextWrite = typer.FileTextWrite()):
    for line in infile:
        outfile.write(line.upper())
```

`typer.FileText` handles `-` as stdin/stdout automatically. Extra configuration (`mode`, `encoding`, `lazy`, `atomic`) can be passed to the marker types: `typer.FileText(encoding="utf-8")`. Typer closes the files when the CLI exits.

## DateTime

```python
from datetime import datetime

def main(
    start: Annotated[datetime, typer.Option()],                          # required, ISO 8601
    end: Annotated[datetime, typer.Option(formats=["%Y-%m-%d"])],        # required, custom format
):
    ...
```

Default parsing accepts ISO 8601 and several common formats; `formats=[...]` restricts parsing to custom strftime formats.

## UUID

`uuid.UUID` parses standard UUID strings and passes a real `uuid.UUID` object.

## Lists, Tuples

```python
import typer
from pathlib import Path

def main(
    files: Annotated[list[Path], typer.Argument(help="All files")],  # variadic positional
    tags: Annotated[list[str], typer.Option("--tag")] = [],          # repeated option
    point: tuple[int, int] = (0, 0),                                  # fixed tuple
):
    ...
```

- `list[X]` **arguments** are variadic (they consume all remaining positional values, so a list argument must be the last argument); `list[X]` **options** are repeatable (`--tag a --tag b`).
- `tuple[X, Y]` consumes exactly the annotated number of values; **variable-length tuples (`tuple[X, ...]`) are not supported**.
- Tuples can carry defaults (shown in the error message, e.g. `Argument 'names' takes 3 values`).
- `set` and other non-list iterables are **not supported** — Typer raises `Type not yet supported`.

## Custom Types and Annotated Metadata

Anything the underlying (vendored) Click `ParamType` supports can be used. Note how `Annotated` metadata is handled: Typer keeps only `typer.Argument()`/`typer.Option()` instances inside `Annotated[...]` and **silently ignores any other metadata** (per PEP 593). Putting more than one Typer marker in one annotation raises `MultipleTyperAnnotationsError`.

## Gotchas

- `Optional[X]` / `X | None` parameters are still CLI parameters (default `None`); `None` does not mean "no parameter".
- Typer silently ignores non-Typer metadata in `Annotated[...]` — a mistyped marker (e.g. forgetting the parentheses in `typer.Option`) disappears without an error.

## References

- <https://typer.tiangolo.com/tutorial/parameter-types/>
- <https://typer.tiangolo.com/tutorial/parameter-types/enum/>
- <https://typer.tiangolo.com/tutorial/parameter-types/file/>
- <https://typer.tiangolo.com/tutorial/parameter-types/path/>

## Related

- Declaration syntax: [parameters.md](parameters.md)
- How conversion is wired at runtime: [internals.md](internals.md)
