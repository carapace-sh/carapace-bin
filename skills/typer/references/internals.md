# Typer Internals

How Typer converts function signatures into a Click command tree at runtime. Source layout, the build pipeline, and the key classes.

> **Source of truth**: <https://github.com/fastapi/typer> (`typer/` package, v0.27.x). Covers behavior since v0.26.0 where Click is vendored.

## Package Layout

```text
typer/
├── __init__.py          # public API exports, __version__
├── main.py              # Typer class, get_command(), run(), launch()
├── core.py              # TyperArgument, TyperOption, TyperCommand, TyperGroup
├── models.py            # Context, ParameterInfo/OptionInfo/ArgumentInfo, CommandInfo, TyperInfo, ...
├── params.py            # public Option() and Argument() factories
├── utils.py             # type-hint introspection, get_params_from_function()
├── _types.py            # TyperChoice (Choice with enum support)
├── _typing.py           # typing compat shim (Annotated, get_origin, ...)
├── cli.py               # the `typer` CLI command (run scripts without packaging)
├── completion.py        # shell completion logic
├── rich_utils.py        # Rich help/error formatting
├── testing.py           # CliRunner
├── exceptions.py        # Abort, Exit, TyperException
├── colors.py            # color constants
└── _click/              # vendored Click (complete copy, since v0.26.0)
```

## The Registration-then-Build Pattern

The `Typer` class is a **registrar**; no Click objects are built until the app is called:

```python
class Typer:
    self.registered_commands: list[CommandInfo]   # from @app.command()
    self.registered_groups: list[TyperInfo]       # from app.add_typer()
    self.registered_callback: TyperInfo | None    # from @app.callback()
```

Decorators append lightweight info records. `app()` → `Typer.__call__` → `get_command(self)(*args)`, which builds the whole Click tree, then runs it.

```text
app() → get_command(app) → TyperGroup | TyperCommand → _main() → command callback
```

## `get_command()` — Shape Decision

In `main.py`, `get_command(typer_instance)` decides the CLI shape:

- **Multiple commands, groups, or a callback** → build a `TyperGroup` (via `get_group_from_info()`).
- **Exactly one command and no callback** → build a single `TyperCommand` (via `get_command_from_info()`).

If `add_completion=True`, `--install-completion`/`--show-completion` params are injected at the top level. `get_group_from_info()` recurses through `registered_groups`, building nested `TyperGroup`s for sub-apps — this is how `add_typer()` composes trees.

## From Signature to Parameters

1. `get_params_from_function(func)` (`utils.py`) — `inspect.signature()` + `get_type_hints()` produce `ParamMeta(name, default, annotation)` per parameter. Handles `Annotated[type, Option()]`, old-style factories-as-defaults, and `default_factory`.
2. `get_click_param(meta)` (`main.py`) — inspects the annotation to route:
   - `typer.Context` → context injection (not a CLI param)
   - `typer.CallbackParam` → callback-parameter injection
   - Everything else → `TyperOption` or `TyperArgument` built from the `OptionInfo`/`ArgumentInfo` and the type
3. `get_callback(func)` (`main.py`) — wraps the user function to inject converted values, plus `Context`/`CallbackParam`/raw `args` where annotated.

## Type Mapping

`get_click_param` maps annotations to vendored-Click types:

| Annotation | Click type |
|------------|-----------|
| `str` | STRING |
| `int` / `float` | INT/FLOAT or IntRange/FloatRange (with min/max/clamp) |
| `bool` | boolean flag with secondary `--no-*` opt |
| `Enum` | TyperChoice (custom Choice with enum support) |
| `Literal[...]` | TyperChoice |
| `pathlib.Path` | TyperPath (models.py, adapted Click Path) |
| `FileText`/`FileTextWrite`/`FileBinaryRead`/`FileBinaryWrite` | types.File with matching mode |
| `datetime` | DateTime (with `formats`) |
| `uuid.UUID` | UUID |
| `list[X]` | multiple=True (options) / nargs=-1 (arguments) |
| `tuple[X, Y]` | tuple with per-position types |

## Core Classes (`core.py`)

| Class | Base | Notable overrides |
|-------|------|-------------------|
| `TyperArgument` | Click Parameter | `make_metavar()`, argument-specific help rendering, `show_envvar`, `rich_help_panel` |
| `TyperOption` | Click Parameter | prompt/confirmation support, boolean `--x/--no-x` secondary opts, envvar resolution, `consume_value()` prompting |
| `TyperCommand` | Click Command | `format_options()` splits **Arguments** and **Options** sections (Click mixes them), Rich help via `rich_utils.rich_format_help()` |
| `TyperGroup` | Click Command (not `click.Group`) | custom `commands` dict, `resolve_command()` with `suggest_commands` (difflib close matches), insertion-order `list_commands()`, `shell_complete()` |

`_main()` in `core.py` replaces Click's `Command.main()`: it handles completion requests, Rich-formatted errors, Abort/Exit exceptions, and EPIPE.

## Models (`models.py`)

- `Context` — thin subclass of vendored Click `Context` (pass-through).
- `OptionInfo` / `ArgumentInfo` (base `ParameterInfo`) — carry every configuration field (`prompt`, `count`, `envvar`, `rich_help_panel`, ...). `params.py` factories `Option()`/`Argument()` just construct these.
- `CommandInfo` / `TyperInfo` — registration records for commands and sub-apps.
- `DefaultPlaceholder` / `Default()` — sentinel distinguishing "user explicitly passed None" from "not set".
- `CallbackParam` — Click `Parameter` subclass injected into parameter callbacks.
- `FileText`/`FileTextWrite`/`FileBinaryRead`/`FileBinaryWrite` — markers converted to Click `File` types.

## Rich Integration (`rich_utils.py`)

- Rich is enabled by default; `core.py` computes `HAS_RICH = parse_boolean_env_var(os.getenv("TYPER_USE_RICH"), default=True)` — `TYPER_USE_RICH=false` disables it.
- When Rich is available, `DEFAULT_MARKUP_MODE = "rich"`, so help renders Rich markup without configuration.
- Rich-formatted help (`rich_format_help`), errors (`rich_format_error`), abort messages, and pretty tracebacks with local variables.
- Controlled further by the `pretty_exceptions_*` constructor flags.

## Vendored Click (`_click/`)

Since v0.26.0 the full Click source lives in `typer/_click/`:

- No external Click dependency at runtime.
- Typer can modify Click internals without upstream coordination (e.g. help layout, completion flow).
- Typer's public API intentionally does **not** expose `typer._click`; import Click-compatible things through `typer.models`/`typer.core` where offered.

## Testing Internals (`testing.py`)

`CliRunner.invoke(app, args, ...)` calls `get_command(app)` then runs it inside `isolation()`, which patches `sys.stdin/stdout/stderr`, `os.environ`, and terminal I/O. `Result` wraps captured output bytes and exit code (see [testing.md](testing.md)).

## Gotchas

- Because building happens at `app()` time, decorating modules after calling the app (import-order bugs) silently drops commands.
- `TyperGroup` is not a `click.Group` subclass — code doing `isinstance(cmd, click.Group)` against real Click fails; use the vendored classes.
- `DefaultPlaceholder` means `param.default is None` checks can be misleading; use Click's `ParamDefault`-aware helpers or Typer's own defaults handling.

## References

- <https://github.com/fastapi/typer/tree/master/typer>
- <https://typer.tiangolo.com/reference/>

## Related

- Public behavior of commands/apps: [commands.md](commands.md)
- Parameter declaration: [parameters.md](parameters.md)
- CliRunner usage: [testing.md](testing.md)
