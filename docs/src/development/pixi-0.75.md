# Pixi 0.75 CLI update checklist

This checklist records the differences found while comparing the local Pixi completer with Pixi `0.75.0`.

- Local CLI: `pixi 0.75.0`
- Reference: [Pixi changelog](https://pixi.prefix.dev/latest/CHANGELOG/)

## Commands

- Add `pixi auth token <HOST>`.
- Keep the existing `upload artifactory` and `workspace export conda-explicit-spec` commands; they are already present in the completer.

## Options to add or update

- Add `--environment` to `add`, `remove`, `task add`, `task alias`, `task remove`, `upgrade`, `workspace channel add`, `workspace channel remove`, `workspace platform add`, and `workspace platform remove`.
- Rename the Git dependency option from `--subdir` to `--subdirectory` in `add`, `remove`, `global add`, and `global install`. Keep `workspace platform edit --subdir` unchanged.
- Add `--build-backend` and `--package` to `global add` and `global install`.
- Add `publish --dry-run` and `publish --no-skip-existing`; replace the old `--skip-existing` option. Keep `--to` as an alias of `--target-channel`.
- Add `self-update --from-url`, `--offline`, `--config-file`, and `--no-config`.
- Add `--offline[=<true|false>]` to all relevant network, solve, install, and update commands:
  `add`, `exec`, `import`, `install`, `lock`, `publish`, `reinstall`, `remove`, `run`, `search`, `self-update`, `shell`, `shell-hook`, `update`, `upgrade`, `upload`, `global add`, `global install`, `global list`, `global remove`, `global sync`, `global update`, `global uninstall`, `global expose add`, `global expose remove`, `global shortcut add`, `global shortcut remove`, `workspace channel add`, `workspace channel remove`, and `workspace export conda-explicit-spec`.
- Add the missing common configuration options to `lock` and `search`: `--auth-file`, `--concurrent-downloads`, `--concurrent-solves`, `--no-hard-links`, `--no-ref-links`, `--no-symbolic-links`, `--pinning-strategy`, `--pypi-keyring-provider`, `--run-post-link-scripts`, `--tls-no-verify`, `--tls-root-certs`, and `--use-environment-activation-cache`.
- Add `--config-file` and `--no-config` to the top-level `upload` command.

## Dynamic completion

Pixi `0.75.0` generates completion values from the CLI. The completer should provide dynamic values for:

- Environment options with `pixi.ActionEnvironments()`.
- Platform options and positional platform names with `pixi.ActionPlatforms()`.
- Feature options and positional feature names with `pixi.ActionFeatures()`.
- Task dependencies with `pixi.ActionTasks()`.

In particular, check `task add` (`--environment`, `--default-environment`, `--depends-on`), `task alias` (environment and dependent tasks), `workspace environment remove`, `workspace feature remove`, and `workspace platform edit`/`move`.

The platform aliases should also be represented consistently: `--auto-detected` and `--current` are aliases of `workspace platform add --auto-detect`, and `--osx` is an alias of `--macos`.
