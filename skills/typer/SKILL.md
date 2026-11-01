---
name: typer
description: >
  Use when building CLI applications with fastapi/typer — creating commands, defining CLI
  arguments and options, type-hint-driven parameter parsing, callbacks, context, help
  customization, shell completion, testing, packaging, or Typer internals. Triggers on:
  "typer", "typer.run", "typer.Typer", "typer app", "typer command", "typer argument",
  "typer option", "typer callback", "typer context", "Annotated", "typer.Argument",
  "typer.Option", "typer.Context", "CliRunner", "add_typer", "rich_markup_mode",
  "typer completion", "typer autocompletion", "typer echo", "typer test", "typer package",
  "typer click", "get_command".
user-invocable: true
---

# Typer CLI Application Reference

Reference for building CLI applications with [typer](https://github.com/fastapi/typer). Covers the public API, type-driven parameter declaration, and how Typer wraps Click internally.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| Typer, typer.run, app, @app.command, @app.callback, add_typer, subcommand, group, one or multiple commands, TyperInfo, CommandInfo, suggest_commands, invoke_without_command | [references/commands.md](references/commands.md) |
| typer.Argument, typer.Option, Annotated, default value style, Ellipsis, required, optional, envvar, prompt, confirmation_prompt, short name, metavar, show_default, help, OptionInfo, ArgumentInfo, ParamMeta | [references/parameters.md](references/parameters.md) |
| type conversion, parameter type, str, int, float, bool, Enum, Literal, choices, case_sensitive, FileText, FileTextWrite, FileBinaryRead, FileBinaryWrite, pathlib.Path, exists, min, max, clamp, count, DateTime, UUID, list, tuple | [references/parameter-types.md](references/parameter-types.md) |
| typer.Context, ctx, CallbackParam, invoked_subcommand, resilient_parsing, allow_extra_args, ignore_unknown_options, find_root, find_object, ensure_object, fail, abort, exit, get_parameter_source | [references/context.md](references/context.md) |
| help, docstring, rich, rich_markup_mode, rich_help_panel, markdown help, deprecated, epilog, printing, echo, style, secho, stderr, rich markup, suggest_commands | [references/help-output.md](references/help-output.md) |
| completion, shell completion, install-completion, show-completion, autocompletion, incomplete, shellingham, custom completion, completion function | [references/completion.md](references/completion.md) |
| testing, pytest, CliRunner, Result, invoke, exit_code, input, isolated filesystem | [references/testing.md](references/testing.md) |
| package, pyproject.toml, project.scripts, wheel, python -m, __main__.py, publish, docs generation | [references/package.md](references/package.md) |
| internals, source code, main.py, core.py, models.py, utils.py, _click, vendored click, get_command, get_group_from_info, get_command_from_info, get_callback, TyperGroup, TyperCommand, TyperOption, TyperArgument, DefaultPlaceholder, rich_utils | [references/internals.md](references/internals.md) |

## Quick Guide

- **How do I create a CLI app, commands, or subcommands?** → [references/commands.md](references/commands.md)
- **How do I declare CLI arguments and options (Annotated vs default-value style)?** → [references/parameters.md](references/parameters.md)
- **Which Python types can I use for CLI parameters (Enum, Path, File, etc.)?** → [references/parameter-types.md](references/parameter-types.md)
- **How do I use the context (typer.Context) in callbacks and commands?** → [references/context.md](references/context.md)
- **How do I customize help text, use Rich markup, or print colored output?** → [references/help-output.md](references/help-output.md)
- **How do I add shell completion or custom value completion?** → [references/completion.md](references/completion.md)
- **How do I test my CLI app with pytest?** → [references/testing.md](references/testing.md)
- **How do I package my CLI app and publish it?** → [references/package.md](references/package.md)
- **How does Typer work internally (Click conversion, module layout)?** → [references/internals.md](references/internals.md)
- **Why is my parameter an option instead of an argument (or vice versa)?** → [references/parameters.md](references/parameters.md) and [references/commands.md](references/commands.md)
- **How do I make subcommands share CLI parameters via a callback?** → [references/commands.md](references/commands.md) and [references/context.md](references/context.md)

## Cross-Project References

- For the completion-aware pflag/cobra ecosystem used in this repo, see the **cobra** skill (in this repo, `skills/cobra/`). Typer is Click-based, not cobra-based; Typer's own internals are covered in [references/internals.md](references/internals.md).
- For FastAPI itself, Typer's design mirrors it (dependency injection, declarative type hints) but there is no FastAPI skill here.

## Note

Typer versions covered: 0.12+ through 0.27.x. Since v0.26.0 Click is vendored inside `typer/_click/`, so no external Click dependency exists at runtime (see [references/internals.md](references/internals.md)).
