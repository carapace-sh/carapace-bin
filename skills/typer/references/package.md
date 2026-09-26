# Packaging and Distributing a Typer CLI

Making a Typer app installable: `pyproject.toml` scripts, `python -m` support, wheels, and publishing.

> **Source of truth**: <https://typer.tiangolo.com/tutorial/package/>. For project structure basics, see [commands.md](commands.md).

## Project Structure

```text
myproject/
├── pyproject.toml
├── README.md
└── src/
    └── myproject/
        ├── __init__.py
        ├── main.py        # app = typer.Typer() lives here
        └── __main__.py    # for python -m support
```

The tutorial uses `uv init --package` to scaffold this layout.

## The Entry Point Script

Map a command name to the app object in `pyproject.toml`:

```toml
[project]
name = "myproject"
dependencies = ["typer"]

[project.scripts]
myproject = "myproject.main:app"

[build-system]
requires = ["setuptools>=61"]
build-backend = "setuptools.build_meta"
```

After `uv pip install -e .` (or `pip install -e .`), the `myproject` command invokes the app. The target can be a `typer.Typer()` instance or any callable — Typer apps are callable.

## `python -m` Support

```python
# src/myproject/__main__.py
from myproject.main import app

app()
```

Then `python -m myproject` works identically to the installed script.

## Building and Publishing

```bash
uv build            # or: python -m build  → creates dist/*.whl and *.tar.gz
uv publish          # or: twine upload dist/*
```

Users install with `pip install myproject` (or `uv pip install myproject`) and get the CLI command plus automatic `--install-completion` (see [completion.md](completion.md)). To install a built wheel as an isolated tool: `uv tool install dist/*.whl`.

## Generating Docs

The `typer` command can generate Markdown documentation for your app (it detects the `typer.Typer` app in the module automatically):

```bash
typer myproject.main utils docs --output README.md --name myproject
```

- `myproject.main` — the module (positional PATH_OR_MODULE)
- `utils docs` — the docs subcommand of the `typer` CLI
- `--output` — write Markdown to a file (omit to print to stdout)
- `--name` — program name used in the generated usage lines

## Gotchas

- With `src` layout, editable installs require a modern build backend; `uv` handles this by default.
- `[project.scripts]` points at the **app object**, not a `main()` function — calling `app()` is what runs the CLI. Pointing at a function works too if you use `typer.run`.
- Shell completion does **not** work via `python -m myproject` — completion is tied to the program name, so call the installed script directly.
- Version bumps for publishing: update `version` in `pyproject.toml` (single source of truth).

## References

- <https://typer.tiangolo.com/tutorial/package/>

## Related

- App structure: [commands.md](commands.md)
- Completion installed with the app: [completion.md](completion.md)
