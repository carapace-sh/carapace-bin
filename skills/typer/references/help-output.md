# Help, Rich Output, and Printing

Help text customization (docstrings, Rich markup, panels, deprecation) and printing output (print, Rich, typer.echo, colors, stderr).

> **Source of truth**: <https://typer.tiangolo.com/tutorial/commands/help/> and <https://typer.tiangolo.com/tutorial/printing/>. For parameter-level help flags, see [parameters.md](parameters.md).

## Help Text Sources

Help text for commands and apps resolves from (highest priority first):

1. `help=` argument on `@app.command()` / `typer.Typer(help=...)` / `app.add_typer(help=...)`
2. Function / module **docstring**

```python
import typer

app = typer.Typer()

@app.command()
def create(name: str):
    """Create a user.          # <- this is the command help

    Some more explanation in the docstring.
    """
    ...
```

The first line becomes the short help (command list); the full docstring becomes the long help.

## Rich Markup Modes

`typer.Typer(rich_markup_mode="rich")` (or `"markdown"`) sets how help text is rendered:

```python
app = typer.Typer(rich_markup_mode="rich")

@app.command()
def create():
    """
    Create something with [red]red text[/red].

    Some **markdown-like emphasis** and [bold]bold[/bold].
    """
```

- `"rich"` mode uses Rich markup: `[bold]`, `[red]`, `[italic]`, etc.
- `"markdown"` mode accepts standard Markdown syntax in help.
- In recent versions the **default is already `"rich"`** (when Rich is available), so markup renders out of the box; older versions without a mode set showed markup literally as plain text.

## Help Panels

Group commands or parameters under titled sections:

```python
app = typer.Typer()

@app.command(rich_help_panel="Customization and configuration")
def config():
    ...

@app.command()
def main_app():  # stays in the default panel
    ...
```

Parameters can also be grouped: `typer.Option(..., rich_help_panel="Extra options")` — grouped options get their own options section in help.

## Other Help Features

- `deprecated=True` on `@app.command()` marks a command as deprecated.
- `epilog="..."` (in `typer.Typer()` or `@app.command()`) adds trailing help text.
- `short_help="..."` overrides the one-liner in the commands list.
- `suggest_commands=True` on the app prints "did you mean" suggestions when a subcommand is mistyped.

## Printing: Standard Python

`print()` works fine for simple output. For stderr:

```python
import sys
print("error!", file=sys.stderr)
```

Rich gives better terminal output:

```python
from rich import print
print(":apple: [bold red]error![/bold red]")   # rich markup + emoji
print({"key": "value"})                         # pretty-printed structures
```

For explicit console control:

```python
from rich.console import Console
stderr_console = Console(stderr=True)
stderr_console.print("[red]error[/red]")
```

## typer.echo / typer.style / typer.secho

Typer's own output helpers (Click heritage, no Rich dependency required):

```python
import typer

typer.echo("plain message")               # to stdout
typer.echo("diagnostic", err=True)        # to stderr
typer.echo(typer.style("warn", fg=typer.colors.YELLOW))
typer.secho("bold red", fg=typer.colors.RED, bold=True)  # print + style
```

Colors constants live in `typer.colors` (`RED`, `GREEN`, `YELLOW`, `BLUE`, ...). Style accepts `fg`, `bg`, `bold`, `dim`, `underline`, `blink`, `reverse`, `reset`.

## Exit Codes and Errors

```python
import typer

if not valid:
    typer.echo("Invalid", err=True)
    raise typer.Exit(code=1)
# or
raise typer.Abort()   # "Aborted!" with exit code 1
```

`typer.Exit` exits cleanly; `typer.Abort` prints an abort message. For validation errors in parameter callbacks, use `typer.BadParameter` (see [context.md](context.md)).

## Gotchas
- Plain `typer` includes Rich and shellingham by default; `typer-slim` is deprecated (since v0.22 it just depends on `typer`) and `typer[all]` is no longer needed.
- `TYPER_USE_RICH=false` (env var) **disables** Rich entirely — help falls back to plain text and errors to plain tracebacks. Rich is enabled by default (`TYPER_USE_RICH` defaults to true).
- `typer.echo` strips styles when not attached to a TTY, unlike raw `print` with ANSI codes.
- Emoji support depends on the terminal font/encoding.

## References

- <https://typer.tiangolo.com/tutorial/commands/help/>
- <https://typer.tiangolo.com/tutorial/printing/>
- <https://typer.tiangolo.com/tutorial/exceptions/>

## Related

- Parameter help flags: [parameters.md](parameters.md)
- Docstring vs `help=` precedence for groups: [commands.md](commands.md)
- `typer.Exit`/`typer.Abort` internals: [internals.md](internals.md)
