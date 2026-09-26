# CLI Parameters: Arguments and Options

Declaring CLI arguments (positional) and options (named) with function parameters. Covers `typer.Argument()`, `typer.Option()`, the Annotated style, defaults, envvar, prompts, and requiredness.

> **Source of truth**: <https://typer.tiangolo.com/tutorial/> (arguments, options). For how types map to parsing/validation, see [parameter-types.md](parameter-types.md).

## The Core Rule

Typer infers CLI parameters from function signature type hints:

- **Positional parameters without a default** → CLI **arguments**
- **Parameters with a default value** → CLI **options** (except `bool` defaults, which become flags)

```python
import typer

def main(name: str, lastname: str = "", formal: bool = False):
    ...
```

- `name` → argument `NAME` (required)
- `lastname` → option `--lastname`
- `formal` → flag `--formal / --no-formal`

## Two Declaration Styles

Both styles work; **Annotated is the recommended modern style**.

### Annotated style (recommended)

```python
from typing import Annotated
from typing_extensions import Annotated  # for Python < 3.9

import typer

def main(
    name: Annotated[str, typer.Argument(help="Who to greet")],
    lastname: Annotated[str, typer.Option("--last")] = "",
    formal: Annotated[bool, typer.Option(help="Say it formally")] = False,
):
    ...
```

The type is the real annotation; `typer.Argument()`/`typer.Option()` carry the extra configuration. **The default value stays the parameter's default**, not inside the factory.

### Old style: factory as the default value

```python
import typer

def main(
    name: str = typer.Argument(..., help="Who to greet"),  # Ellipsis = required
    lastname: str = typer.Option("", "--last"),
):
    ...
```

Here the first positional argument of `Argument()`/`Option()` is the default; `...` (`Ellipsis`) marks it required. Mixing styles is allowed but confusing; prefer Annotated.

## `typer.Argument()`

Configures a positional argument. Common parameters:

| Parameter | Description |
|-----------|-------------|
| `default` | Default value (`...` = required) — only meaningful in old style |
| `help` | Help text |
| `metavar` | Custom placeholder shown in help |
| `show_default` | Show default in help (`True`/`False`/custom string) |
| `hidden` | Hide from help |
| `envvar` | Env var fallback (str or list of str) |
| `show_envvar` | Show env var in help |
| `rich_help_panel` | Group under a named help panel |
| `case_sensitive` | Case-insensitive choices (Enums) |

Optional arguments get a default; `list[str]` arguments are variadic (nargs=-1).

## `typer.Option()`

Configures a named option. Everything from `Argument()` plus:

| Parameter | Description |
|-----------|-------------|
| first positional args | Option names: `typer.Option("--name", "-n")` (Annotated: passed as positional args to the factory) |
| `prompt` | Prompt interactively if not provided (`True` or a custom string) |
| `confirmation_prompt` | Ask twice (for secrets/passwords) |
| `hide_input` | Hide typed input (passwords) |
| `count` | Repeated counter option (`-v -v -v`) |
| `allow_from_autoenv` | Allow env var auto-detection |

### Option names

```python
def main(
    # Annotated style: names are positional args to Option()
    name: Annotated[str, typer.Option("--name", "-n")] = "",
    # Only a short name
    force: Annotated[bool, typer.Option("-f")] = False,
    # Short names can be combined: -fn
    full: Annotated[bool, typer.Option("-fn")] = False,
):
    ...
```

For flags, Typer generates both `--flag` and `--no-flag` variants (so a `True` default can be turned off with `--no-flag`). `bool` options with a `False` default can be used as `--flag` alone.

## Required and Optional

| Goal | Annotated style | Old style |
|------|----------------|-----------|
| Required argument | `Annotated[str, typer.Argument()]` (no default) | `= typer.Argument(...)` |
| Optional argument | `Annotated[str, typer.Argument()] = "x"` | `= typer.Argument("x")` |
| Required option | `Annotated[str, typer.Option()]` (no default) | `= typer.Option(...)` (or bare `= typer.Option()`) |
| Optional option | `Annotated[str, typer.Option()] = "x"` | `= typer.Option("x")` |

In Annotated style, requiredness comes purely from whether the parameter has a default value; the docs don't use `Ellipsis` there — leave the default off instead. In old style, a bare `typer.Option()` with no arguments also means required (Ellipsis is its implicit default).

## Environment Variables

```python
def main(
    name: Annotated[str, typer.Argument(envvar="MY_NAME")] = "World",
    token: Annotated[str, typer.Option(envvar=["TOKEN", "API_TOKEN"], show_envvar=False)] = "",
):
    ...
```

`envvar` accepts a string or a list (tried in order). CLI values take precedence over env vars. `show_envvar=False` hides it from help.

## Prompts

```python
def main(
    name: Annotated[str, typer.Option(prompt=True)],
    password: Annotated[str, typer.Option(prompt="Password", hide_input=True, confirmation_prompt=True)] = "",
):
    ...
```

Prompted options must not have a default (in Annotated style) or must use `...` (old style), otherwise the default short-circuits the prompt.

## Help Text and Defaults

- `help=` adds per-parameter help; it is combined with (and overrides) the docstring mention.
- `show_default=True|False|"custom text"` controls default display.
- `metavar="text"` customizes the placeholder (argument names in usage).

## Callbacks for Validation (Parameter-Level)

Any parameter can take a custom callback that runs during parsing:

```python
import typer

def validate_name(value: str):
    if value != "Camila":
        raise typer.BadParameter("Only Camila is allowed")
    return value

def main(
    name: Annotated[str, typer.Option(callback=validate_name)] = "",
):
    ...
```

Callbacks receive the value; raise `typer.BadParameter` for validation errors. The callback signature can also accept `typer.Context` and the `CallbackParam`; details in [context.md](context.md).

## Gotchas

- In the old style, `typer.Option(...)` with no name generates the name from the parameter (`--lastname`); explicit names always win.
- `bool` arguments (positional) are not flags; flags only come from `bool` *options*.
- `count=True` requires `int` type.
- Typer raises `AnnotatedParamWithDefaultValueError` if you put a default inside `Annotated[..., typer.Option(...)]` instead of using `=`, and `MixedAnnotatedAndDefaultStyleError` if you use both styles on one parameter.
- A parameter annotated `typer.Context` is not a CLI parameter at all — it injects the context (see [context.md](context.md)).

## References

- <https://typer.tiangolo.com/tutorial/arguments/>
- <https://typer.tiangolo.com/tutorial/options/>
- <https://typer.tiangolo.com/reference/parameters/>

## Related

- Type conversion and validation: [parameter-types.md](parameter-types.md)
- Parameter-level callbacks and context: [context.md](context.md)
- Help rendering: [help-output.md](help-output.md)
