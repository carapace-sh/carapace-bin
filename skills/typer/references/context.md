# The Context and CallbackParam

`typer.Context` — passing state between callbacks and commands, inspecting invocation, and handling unknown arguments. Also covers `typer.CallbackParam`.

> **Source of truth**: <https://typer.tiangolo.com/reference/context/> and <https://typer.tiangolo.com/tutorial/commands/context/>.

## Getting the Context

Annotate a parameter as `typer.Context` (Annotated or default style both work):

```python
import typer

app = typer.Typer()

@app.command()
def hello(ctx: typer.Context, name: str):
    print(f"params: {ctx.params}")
    print(f"raw args: {ctx.args}")
    ctx.obj["greeted"] = name
```

The context parameter is not a CLI parameter; it is injected at invocation time. A command callback may take it positionally anywhere in the signature.

## Key Context Attributes

| Attribute | Description |
|-----------|-------------|
| `ctx.params` | Parsed parameters of the current command (dict) |
| `ctx.args` | Remaining raw args after parsing (needs `allow_extra_args`) |
| `ctx.invoked_subcommand` | Name of the subcommand being invoked (in callbacks) |
| `ctx.resilient_parsing` | `True` during shell completion — don't execute side effects |
| `ctx.obj` | Arbitrary shared dict; propagate state callback → command |
| `ctx.parent` | Parent context (walk the chain with `ctx.parent.parent...`) |
| `ctx.command` | The Click command being invoked |

## Passing State: `ctx.obj`

```python
import typer

app = typer.Typer()

@app.callback()
def callback(ctx: typer.Context, verbose: bool = False):
    ctx.obj = {"verbose": verbose}

@app.command()
def create(ctx: typer.Context):
    print(f"verbose: {ctx.obj['verbose']}")
```

Each command gets its own context; `ctx.obj` is inherited from the parent if unset.

## Detecting the Subcommand in Callbacks

```python
@app.callback()
def callback(ctx: typer.Context):
    if ctx.invoked_subcommand is None:
        typer.echo("No command executed")  # app called without a subcommand
    elif ctx.invoked_subcommand == "hello":
        typer.echo("About to say hello")
```

## Executable Callback (no subcommand required)

By default a callback with `ctx.invoked_subcommand` set runs only as a pre-command hook. To make the callback itself executable:

```python
@app.callback(invoke_without_command=True)
def callback(ctx: typer.Context):
    if ctx.invoked_subcommand is None:
        typer.echo("Running standalone callback")
```

## Configuring the Context

`@app.callback()` and `@app.command()` accept Click context settings:

```python
@app.command(context_settings={"allow_extra_args": True, "ignore_unknown_options": True})
def main(ctx: typer.Context):
    for arg in ctx.args:
        print(arg)
```

- `allow_extra_args=True` → unmatched args land in `ctx.args` instead of erroring.
- `ignore_unknown_options=True` → unknown `--options` are not treated as errors.
- `help_option_names=["-h", "--help"]` → additional help flags.

## Context Methods

| Method | Description |
|--------|-------------|
| `ctx.fail(message)` | Abort with a usage error (exit code 2) |
| `ctx.abort()` | Abort without message |
| `ctx.exit(code)` | Exit with a code |
| `ctx.get_help()` | Get help text for the current command |
| `ctx.invoke(other_cmd, **kwargs)` | Invoke another command with the same context chain |
| `ctx.find_root()` | Find the topmost context |
| `ctx.find_object(cls)` / `ctx.ensure_object(cls)` | Object lookup/creation up the context chain |
| `ctx.get_parameter_source(name)` | Where a param came from: `COMMANDLINE`, `ENVIRONMENT`, `DEFAULT`, `PROMPT` |

## CallbackParam

Inside a **parameter callback** (validation callback of an option/argument), Typer injects the `typer.CallbackParam` object — the Click `Parameter` that declared the callback:

```python
import typer

def name_callback(ctx: typer.Context, param: typer.CallbackParam, value: str):
    if param.name == "name" and value == "evil":
        raise typer.BadParameter("Evil name")
    if ctx.resilient_parsing:  # don't break shell completion
        return value
    return value

def main(
    name: Annotated[str, typer.Option(callback=name_callback)] = "",
):
    ...
```

`CallbackParam` exposes `name`, `opts` (e.g. `["--name"]`), `required`, and other Click parameter metadata.

## Gotchas

- `ctx.resilient_parsing` is `True` while generating completions; parameter callbacks that query services must return early (see [completion.md](completion.md)).
- `ctx.params` in the app callback contains only callback-declared parameters; command params arrive later.
- `get_parameter_source` values are Click enums, compare against `click.core.ParameterSource` or the strings.

## References

- <https://typer.tiangolo.com/reference/context/>
- <https://typer.tiangolo.com/tutorial/commands/context/>
- <https://typer.tiangolo.com/tutorial/options/callback-and-context/>

## Related

- App-level callbacks: [commands.md](commands.md)
- Parameter-level callbacks: [parameters.md](parameters.md)
- Resilient parsing during completion: [completion.md](completion.md)
