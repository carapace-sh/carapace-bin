# Shell Completion and Custom Autocompletion

Built-in shell completion installation and custom value completion for CLI parameters.

> **Source of truth**: <https://typer.tiangolo.com/tutorial/options-autocompletion/> and <https://typer.tiangolo.com/tutorial/install/>. For completion-time guards, see [context.md](context.md).

## Built-In Shell Completion

Every Typer app (with `add_completion=True`, the default) gets two injected options:

```
myapp --install-completion   # install completion script for the current shell
myapp --show-completion      # print the completion script to stdout
```

Supported shells: Bash, Zsh, Fish, PowerShell. Shell detection uses **shellingham**. On install, completion is wired into the shell's config file (e.g. `~/.bashrc`, zsh fpath). The generated completion completes subcommands, option names, and Enum/Literal choices out of the box.

The underlying mechanism is vendored Click's `shell_completion` module invoked through `TyperGroup.shell_complete()` (see [internals.md](internals.md)).

## Custom Value Completion

Add a completion function for a parameter with `autocompletion=`:

```python
import typer
from typing import Annotated, Iterable

def complete_name(incomplete: str) -> Iterable[tuple[str, str]]:
    for name in ["Camila", "Carlos", "Sebastian"]:
        if name.startswith(incomplete):
            yield (name, f"help text for {name}")   # (value, help)

def main(
    name: Annotated[str, typer.Option(autocompletion=complete_name)] = "",
):
    print(f"Hello {name}")
```

The completion function receives the currently typed incomplete value (`incomplete: str`) and yields or returns:

- plain strings: `"value"`
- tuples with help: `("value", "help")`
- tuples with help + style: `("value", "help", "cyan")`

A generator with `yield` is the simplest form; a plain `list` return works too.

## Accessing the Context and Raw Args

Completion functions can take additional injected parameters (order-independent):

```python
from typing import Annotated
import typer

def complete_ctx(ctx: typer.Context, incomplete: str):
    # ctx.params holds already-parsed parameters of the command
    return [f"{ctx.params['lang']}-{incomplete}"]

def complete_args(ctx: typer.Context, args: list[str], incomplete: str):
    # args = raw CLI parameters before the current word
    return [a for a in ["a", "b"] if a.startswith(incomplete)]

def main(
    lang: Annotated[str, typer.Option("--lang")] = "en",
    name: Annotated[str, typer.Option(autocompletion=complete_ctx)] = "",
    extra: Annotated[list[str], typer.Argument(autocompletion=complete_args)] = [],
):
    ...
```

Declared parameters (`typer.Context`, `args: list[str]`) are detected and injected; everything else must be typed.

## Completion-Time Guards

During completion, Typer sets `ctx.resilient_parsing = True` and skips most validation/conversion:

```python
def name_callback(ctx: typer.Context, value: str):
    if ctx.resilient_parsing:
        return value  # don't hit the network while completing
    return validate(value)
```

Parameter callbacks that do heavy work (API calls, filesystem checks) must short-circuit on `resilient_parsing`, or completion breaks or slows down (see [context.md](context.md)).

## The `typer` Script and Completion

Running a script via the `typer` command uses the `run` subcommand:

```bash
typer main.py run hello --name Camila
```

Completion for the `typer` command itself is installed with `typer --install-completion`; it then completes your scripts' commands and options too.

## Gotchas

- Completion functions run *outside your app execution*; they cannot rely on parsed values being complete.
- `autocompletion` is set inside `typer.Option()`/`typer.Argument()` in both styles (old style: `typer.Option(default=..., autocompletion=func)`).
- Enum and `Literal` parameters complete automatically; custom completion is only needed for dynamic values.
- Broken completion functions are silently ignored by some shells — test with the shell directly.

## References

- <https://typer.tiangolo.com/tutorial/options-autocompletion/>
- <https://typer.tiangolo.com/tutorial/install/>
- <https://click.palletsprojects.com/shell-completion/> (underlying mechanism, vendored)

## Related

- Resilient parsing: [context.md](context.md)
- How `--install-completion` is injected at build time: [internals.md](internals.md)
