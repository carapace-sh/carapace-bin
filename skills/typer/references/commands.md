# Commands, Apps, and Subcommands

How to structure a Typer application: `typer.run()`, the `typer.Typer()` app object, `@app.command()`, `@app.callback()`, and composing sub-apps with `app.add_typer()`.

> **Source of truth**: <https://typer.tiangolo.com/tutorial/> (first-steps, typer-app, commands). For how these registrations become Click commands at runtime, see [internals.md](internals.md).

## Two Ways to Start

### `typer.run()` — the simple script path

```python
import typer

def main(name: str, formal: bool = False):
    if formal:
        print(f"Goodbye Ms. {name}. Have a good day.")
    else:
        print(f"Bye {name}!")

if __name__ == "__main__":
    typer.run(main)
```

`typer.run(func)` creates an ephemeral `Typer()` app, registers `func` as its single command, and invokes it. Introspection of `func`'s type hints generates the CLI parameters (see [parameters.md](parameters.md)).

### Explicit `typer.Typer()` app

```python
import typer

app = typer.Typer()

@app.command()
def hello(name: str):
    print(f"Hello {name}")

@app.command()
def goodbye(name: str, formal: bool = False):
    print(f"Goodbye {name}{'.' if formal else '!'}")

if __name__ == "__main__":
    app()
```

Calling `app()` runs the CLI (Typer's `__call__` builds the Click tree and executes it, see [internals.md](internals.md)).

## Single vs Multiple Commands

The number of registered commands changes the CLI's shape:

- **One command** → that command *is* the app: `python main.py NAME --formal` (no subcommand name).
- **Multiple commands** → they become subcommands: `python main.py hello NAME`.

A **callback** (or group settings) forces single-command shape: with an explicit `@app.callback()`, the single command becomes a subcommand again.

```python
@app.command()
def create():
    ...

@app.callback()
def callback():
    """A callback forces subcommand mode."""
```

Now `create` is invoked as `python main.py create` even though it is the only command.

## `@app.command()`

Registers a function as a CLI command. Parameters:

| Parameter | Type | Description |
|-----------|------|-------------|
| `name` | `str` | Command name (defaults to function name with `_` → `-`) |
| `help` | `str` | Help text (defaults to docstring) |
| `cls` | `type` | Custom Click command class |
| `short_help` | `str` | Short help for the commands list |
| `hidden` | `bool` | Hide from help and completion |
| `deprecated` | `bool` | Mark command as deprecated |
| `epilog` | `str` | Trailing help text |
| `short_help` | `str` | One-liner in the commands list (overrides docstring first line) |
| `rich_help_panel` | `str` | Group under a named help panel |
| `add_help_option` | `bool` | Add a `--help` option (default `True`) |
| `no_args_is_help` | `bool` | Show help when the command gets no args (default `False`) |
| `options_metavar` | `str` | Placeholder for options in usage line |
| `context_settings` | `dict` | Forwarded to Click (e.g. `ignore_unknown_options`) |

## `@app.callback()`

The app-level callback runs **before** any subcommand. Use it to define CLI parameters that apply to all commands (`--verbose`), to control single-command behavior, or for documentation.

```python
import typer

app = typer.Typer()

@app.callback()
def callback(verbose: bool = False):
    if verbose:
        print("Will write verbose output")
```

The callback's parameters become options of the *main* command (`python main.py --verbose hello`), not of the subcommands.

The pretty exception flags (`pretty_exceptions_enable`, `pretty_exceptions_show_locals`, `pretty_exceptions_short`) belong to `typer.Typer()`, not `@app.command()`.

It can also be provided at construction: `typer.Typer(callback=callback)`; a later `@app.callback()` replaces it.

For `invoke_without_command=True` and executable callbacks, see [context.md](context.md).

## Sub-Apps: `app.add_typer()`

Compose larger CLIs by nesting Typer apps — each sub-app becomes a command group:

```python
import typer

app = typer.Typer()
items_app = typer.Typer()
users_app = typer.Typer()

app.add_typer(items_app, name="items")
app.add_typer(users_app, name="users")

@items_app.command("create")
def items_create(name: str):
    ...
```

This creates `python main.py items create NAME`. Sub-apps nest arbitrarily deep.

### `add_typer()` help precedence

Help text for a group resolves in this order (highest wins):

1. `name`/`help` passed to `app.add_typer(sub, name=..., help=...)`
2. `help` passed to `typer.Typer(help=...)` of the sub-app itself
3. `help` on the sub-app's callback (lowest priority)

A `@app.callback()` on the sub-app overrides the callback set in `typer.Typer()`, but does not override `add_typer(help=...)`.

## `typer.Typer()` Constructor

Key parameters:

| Parameter | Description |
|-----------|-------------|
| `name` | App name shown in help |
| `help` | Main app help text (overrides docstring) |
| `epilog` | Text after help |
| `callback` | App-level callback |
| `add_completion` | Inject `--install-completion`/`--show-completion` (default `True`) |
| `add_help_option` | Add `--help` option (default `True`) |
| `no_args_is_help` | Show help when called without args |
| `invoke_without_command` | Make the callback executable without a subcommand |
| `rich_markup_mode` | `"rich"` or `"markdown"` help rendering |
| `rich_help_panel` | Default help panel for parameters |
| `suggest_commands` | Suggest similar commands on typos (default `True` in recent versions) |

## Gotchas

- The `@app.command()` decorator runs at import time; a function decorated twice or called directly still works as a plain Python function.
- `typer.run()` and a single `@app.command()` without a callback do **not** create a subcommand — this surprises people migrating between the two styles.
- Command names convert underscores to dashes (`do_work` → `do-work`); pass `name=` explicitly to control this.
- Registration order determines the order commands appear in help (TyperGroup preserves insertion order, see [internals.md](internals.md)).

## References

- <https://typer.tiangolo.com/tutorial/typer-app/>
- <https://typer.tiangolo.com/tutorial/commands/one-or-multiple/>
- <https://typer.tiangolo.com/tutorial/commands/callback/>
- <https://typer.tiangolo.com/tutorial/subcommands/add-typer/>
- <https://typer.tiangolo.com/tutorial/subcommands/name-and-help/>

## Related

- Parameter declaration: [parameters.md](parameters.md)
- Context sharing between callback and commands: [context.md](context.md)
- Runtime build process: [internals.md](internals.md)
