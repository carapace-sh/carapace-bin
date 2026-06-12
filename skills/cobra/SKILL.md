---
name: cobra
description: >
  Use when building CLI applications with spf13/cobra — creating commands, defining flags,
  argument validation, lifecycle hooks, help/usage customization, shell completions, or
  Viper integration. Triggers on: "cobra", "cobra command", "cobra flag", "cobra subcommand",
  "cobra completion", "cobra help", "cobra usage", "cobra hook", "cobra viper",
  "AddCommand", "PersistentFlags", "LocalFlags", "MarkFlagRequired", "RunE", "PreRun",
  "PostRun", "PersistentPreRun", "ValidArgs", "ValidArgsFunction", "RegisterFlagCompletionFunc",
  "ShellCompDirective", "ExactArgs", "NoArgs", "MinimumNArgs", "MatchAll",
  "MarkFlagsRequiredTogether", "MarkFlagsMutuallyExclusive", "cobra-cli",
  "cobra completion", "cobra active help", "cobra plugin".
user-invocable: true
---

# Cobra CLI Application Reference

Reference for building CLI applications with [cobra](https://github.com/spf13/cobra). Covers the public API, common patterns, and best practices.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| Command, AddCommand, subcommand, root command, project structure, command fields, Use, Short, Long, Example, Aliases, Deprecated, Hidden, GroupID, Run, RunE | [references/commands.md](references/commands.md) |
| flag, PersistentFlags, LocalFlags, required flag, flag group, required together, mutually exclusive, one required, repeated flag, count flag, string slice, TraverseChildren, MarkFlagRequired, MarkPersistentFlagRequired | [references/flags.md](references/flags.md) |
| argument, positional arg, Args, NoArgs, ArbitraryArgs, MinimumNArgs, MaximumNArgs, ExactArgs, RangeArgs, OnlyValidArgs, NoDuplicateArgs, MatchAll, ValidArgs, ArgAliases, custom validator | [references/arguments.md](references/arguments.md) |
| hook, lifecycle, PreRun, PostRun, PersistentPreRun, PersistentPostRun, PreRunE, PostRunE, EnableTraverseRunHooks, execution order, inherited hooks | [references/hooks.md](references/hooks.md) |
| help, usage, template, SetHelpFunc, SetHelpTemplate, SetUsageFunc, SetUsageTemplate, version, SetVersionTemplate, help command, command grouping, GroupID, SetHelpCommandGroupId, SetCompletionCommandGroupId, ErrPrefix | [references/help-usage.md](references/help-usage.md) |
| completion, shell completion, ValidArgs, ValidArgsFunction, CompletionFunc, RegisterFlagCompletionFunc, ShellCompDirective, NoFileCompletions, FixedCompletions, active help, AppendActiveHelp, GetActiveHelpConfig, CompletionOptions, completion command, bash, zsh, fish, powershell | [references/completions.md](references/completions.md) |
| viper, config, OnInitialize, BindPFlag, config file, 12-factor, initConfig, AutomaticEnv | [references/viper-integration.md](references/viper-integration.md) |

## Quick Guide

- **How do I create a command or subcommand?** → [references/commands.md](references/commands.md)
- **How do I add flags (persistent, local, required, groups)?** → [references/flags.md](references/flags.md)
- **How do I validate positional arguments?** → [references/arguments.md](references/arguments.md)
- **How do lifecycle hooks work (PreRun, PostRun, etc.)?** → [references/hooks.md](references/hooks.md)
- **How do I customize help/usage/version output?** → [references/help-usage.md](references/help-usage.md)
- **How do I add shell completions?** → [references/completions.md](references/completions.md)
- **How do I integrate with Viper for config?** → [references/viper-integration.md](references/viper-integration.md)
- **How do I make a kubectl-style plugin?** → [references/commands.md](references/commands.md)
- **How do I add Active Help messages during completion?** → [references/completions.md](references/completions.md)
- **How do I make flags required-together or mutually-exclusive?** → [references/flags.md](references/flags.md)

## Cross-Project References

- For cobra internals (Command struct fields, ExecuteC dispatch, Find vs Traverse, flag resolution, completion engine, shell script generation, global state), use the **cobra-dev** skill.
- For pflag internals (flag set implementation, POSIX compliance, non-POSIX modes), use the **carapace-dev** skill → `references/pflag.md`.
