# Testing Typer Apps

Testing CLI applications with pytest and `typer.testing.CliRunner`.

> **Source of truth**: <https://typer.tiangolo.com/tutorial/testing/>. For `Result` internals, see [internals.md](internals.md).

## Basic Pattern

```python
from typer.testing import CliRunner
from myapp.main import app  # your typer.Typer() app

runner = CliRunner()

def test_app():
    result = runner.invoke(app, ["hello", "Camila"])
    assert result.exit_code == 0
    assert "Hello Camila" in result.output
```

`runner.invoke(app, args)` builds the Click command via `get_command(app)` and executes it with the given args, capturing stdout/stderr.

## Checking Results

`Result` exposes:

| Attribute | Description |
|-----------|-------------|
| `exit_code` | Process exit code (0 on success, 2 on usage errors) |
| `output` | Combined stdout+stderr as str (matches what users see) |
| `stdout`, `stderr` | Separately captured streams |
| `exception` | Unhandled exception, if `catch_exceptions=True` (default) |
| `exc_info` | Exception traceback info |

With `catch_exceptions=True` (default) an unhandled exception is captured into `result.exception` instead of failing the test; assert on `result.exit_code != 0` and `result.exception`.

## Testing Errors and Edge Cases

```python
def test_missing_arg():
    result = runner.invoke(app, ["hello"])
    assert result.exit_code != 0
    assert "Missing argument" in result.output
```

## Testing Input (Prompts)

```python
def test_prompt():
    result = runner.invoke(app, ["hello"], input="Camila\n")
    assert "Name:" in result.output
```

`input=` simulates stdin; end each simulated line with `\n`.

## Testing a Bare Function

You can invoke a plain function (created with `typer.run`-style signatures) without an explicit app — `CliRunner.invoke` accepts functions that Typer can convert:

```python
from typer.testing import CliRunner
from myapp.main import main  # plain function

runner = CliRunner()

def test_function():
    result = runner.invoke(main, ["Camila"])
    assert result.exit_code == 0
```

Typer wraps the function in an ephemeral app under the hood.

## Isolated Filesystem and Environment

`CliRunner` (Click heritage) supports:

- `runner.invoke(app, args, env={"MY_VAR": "x"})` — temporarily set env vars.
- `with runner.isolated_filesystem():` — run in a temp directory that is cleaned up.

## Gotchas

- `@app.command()` registers into the app's global state; importing the module multiple times in tests can double-register commands. Import once (pytest caches modules).
- Rich-formatted output may include ANSI codes and wrapping; prefer substring assertions on stable words, or set `env={"NO_COLOR": "1"}` / `TERM=dumb` to reduce formatting.
- `result.output` merges stdout and stderr; use `result.stdout`/`result.stderr` when the destination matters (requires the app to write via `typer.echo(err=True)` or a Rich stderr console).
- Because `invoke` catches exceptions by default, a test asserting success should also check `result.exception is None` when unsure.

## References

- <https://typer.tiangolo.com/tutorial/testing/>
- Click's CliRunner docs (vendored implementation)

## Related

- Building the Click tree being invoked: [internals.md](internals.md)
- Output helpers: [help-output.md](help-output.md)
