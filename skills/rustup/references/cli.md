# Top-Level Commands and Global Flags

The rustup CLI surface — top-level commands, management commands, and the global flags shared across all subcommands.

## Global Flags

These flags apply to most `rustup` subcommands:

| Flag | Short | Description |
|------|-------|-------------|
| `--verbose` | `-v` | Set log level to `DEBUG` (if `RUSTUP_LOG` is unset) |
| `--quiet` | `-q` | Suppress progress output; set log level to `WARN` (if `RUSTUP_LOG` is unset) |
| `--help` | `-h` | Print help for the command |
| `--version` | `-V` | Print rustup version (root command only) |
| `--no-self-update` | | Suppress automatic self-update on this invocation |

The `+<toolchain>` shorthand is a pseudo-flag accepted by proxy commands and `rustup run`. See [proxy.md](proxy.md) and [override.md](override.md).

## Top-Level Commands

| Command | Description |
|---------|-------------|
| `rustup install <toolchain>` | Install a toolchain (shorthand for `rustup toolchain install`) |
| `rustup uninstall <toolchain>` | Remove a toolchain (shorthand for `rustup toolchain uninstall`) |
| `rustup update [toolchain...]` | Update installed toolchains to latest on their channels; also checks for rustup self-update |
| `rustup check` | Check for available updates without downloading them |
| `rustup default [toolchain]` | Set or query the global default toolchain |
| `rustup toolchain <sub>` | Toolchain management — see [toolchain.md](toolchain.md) |
| `rustup target <sub>` | Cross-compilation target management — see [target.md](target.md) |
| `rustup component <sub>` | Component management — see [component.md](component.md) |
| `rustup override <sub>` | Directory overrides — see [override.md](override.md) |
| `rustup run <toolchain> <cmd> [args]` | Run a command with a specific toolchain — see [proxy.md](proxy.md) |
| `rustup which <cmd>` | Show path to the binary a proxy resolves to — see [proxy.md](proxy.md) |
| `rustup doc` | Open offline Rust documentation — see [proxy.md](proxy.md) |
| `rustup man <cmd>` | View the man page for a Rust tool command (Unix only) — see [proxy.md](proxy.md) |
| `rustup self <sub>` | Self-management — see below |
| `rustup set <sub>` | Configuration settings — see below |
| `rustup show [sub]` | Display toolchain/override info — see below |
| `rustup home` | Print the rustup home directory path |
| `rustup completions <shell>` | Generate shell completion scripts |
| `rustup help [sub...]` | Display help; accepts nested subcommand names as positional args |

## `rustup install`

Shorthand for `rustup toolchain install`. See [toolchain.md](toolchain.md) for flags.

```
rustup install stable
rustup install nightly-2024-01-15
rustup install 1.70.0
```

## `rustup uninstall`

Shorthand for `rustup toolchain uninstall`.

```
rustup uninstall nightly-2024-01-15
```

## `rustup update`

Update installed toolchains to the latest releases on their channels. With no argument, all installed toolchains are updated. Also checks whether rustup itself needs updating.

| Flag | Description |
|------|-------------|
| `--force` | Force update even if components are missing |
| `--no-self-update` | Skip updating rustup itself |
| `--force-non-host` | Allow updating non-host toolchains |

## `rustup check`

Check for available updates without downloading them. Reports the installed and latest versions for each channel and for rustup itself.

| Flag | Description |
|------|-------------|
| `--no-self-update` | Skip checking for rustup updates |

## `rustup default`

Set or query the global default toolchain. With no argument, prints the current default. With a toolchain name, sets it (and installs it if not present).

```
rustup default                  # print current default
rustup default stable           # set stable as default (installs if needed)
rustup default nightly-2024-01-15
```

| Flag | Description |
|------|-------------|
| `--profile <name>` | Installation profile when installing (`minimal`/`default`/`complete`) — see [component.md](component.md) |

## `rustup self` — Self-Management

| Subcommand | Description |
|------------|-------------|
| `rustup self update` | Update rustup itself to the latest version (does not download new toolchains) |
| `rustup self uninstall` | Remove rustup from the system |
| `rustup self upgrade-data` | Upgrade rustup's internal data format |

`rustup self uninstall` flags:

| Flag | Description |
|------|-------------|
| `-y` | Skip confirmation prompt |
| `--no-modify-path` | Do not modify the `PATH` environment variable |

## `rustup set` — Configuration Settings

| Subcommand | Description |
|------------|-------------|
| `rustup set default-host <triple>` | Set the default host target triple |
| `rustup set profile <name>` | Set the default installation profile (`minimal`/`default`/`complete`) |
| `rustup set auto-self-update <mode>` | Control auto self-update: `enable`/`disable`/`check-only` |
| `rustup set auto-install <mode>` | Control auto-install of missing toolchains: `enable`/`disable` |

These settings are persisted in `${RUSTUP_HOME}/settings.toml` — see [environment.md](environment.md).

## `rustup show`

Display information about installed toolchains, the active toolchain, and directory overrides.

| Variant | Description |
|---------|-------------|
| `rustup show` | Show all info (active toolchain, installed toolchains, overrides) |
| `rustup show active-toolchain` | Show only the currently active toolchain |
| `rustup show home` | Show the rustup home directory path (same as `rustup home`) |
| `rustup show profile` | Show the default profile setting |

| Flag | Description |
|------|-------------|
| `--verbose` / `-v` | More detailed output |

## `rustup home`

Print the absolute path to the rustup home directory (default: `~/.rustup`). Equivalent to `rustup show home`. Overridden by `RUSTUP_HOME` — see [environment.md](environment.md).

## `rustup completions`

Generate shell completion scripts for the specified shell. Output goes to stdout; redirect to the appropriate completion file. The optional second argument selects whether to generate completions for `rustup` (default) or `cargo`:

```
rustup completions <shell> [command]
```

| Shell | Install location |
|-------|------------------|
| `bash` | `~/.local/share/bash-completion/completions/rustup` |
| `fish` | `~/.config/fish/completions/rustup.fish` |
| `zsh` | `~/.zfunc/_rustup` |
| `powershell` | append to `$PROFILE` |
| `elvish` | `~/.elvish/completions/rustup.elv` (not officially documented) |

```
rustup completions bash > ~/.local/share/bash-completion/completions/rustup
rustup completions zsh  > ~/.zfunc/_rustup
rustup completions bash cargo > ~/.local/share/bash-completion/completions/cargo
```

Supported shells: `bash`, `fish`, `zsh`, `powershell`, `elvish`.

## `rustup help`

Display help for rustup commands. Accepts nested subcommand names as positional arguments, equivalent to `--help` but with a consistent invocation form.

```
rustup help                       # top-level help
rustup help toolchain             # help for toolchain subcommand
rustup help toolchain install     # help for nested subcommand
```

## References

- [rustup book — Commands](https://rust-lang.github.io/rustup/basics.html)
