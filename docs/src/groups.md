# Groups

Completers are organized into **groups**.

- `android` termux completers
- `bash` bash completers
- `bridge` bridged completers
- `bsd` bsd-like completers
- `cmd` cmd completers
- `common` common completers
- `darwin` macos completers
- `elvish` elvish completers
- `fish` fish completers
- `linux` linux completers
- `unix` unix-like completers
- `user` user specs
- `system` system specs
- `windows` windows completers
- `zsh` zsh completers

You can **list** available completers of a **group** with `carapace --list @{group}`.

![](./groups/group.cast)

> Binaries only contain **relevant groups** unless built with the [build tag](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags) `force_all`.

> Shell groups (`bash`, `bash-ble`, `cmd`, `elvish`, `fish`, `nushell`, `oil`, `powershell`, `tcsh`, `xonsh`, `zsh`) contain completers for commands and functions of the respective shell. They are **experimental** and only active if enabled via [`CARAPACE_BUILTINS`](./setup/environment.md#carapace_builtinss). The one matching `CARAPACE_SHELL` is boosted in [priority](#priority) above the other enabled shell groups.

## Priority

Multiple **groups** providing a completer for a command are ordered by **priority**.

- darwin
  1. `user`
  1. `system`
  1. shell matching `CARAPACE_SHELL`
  1. `darwin`
  1. `bsd`
  1. `unix`
  1. `common`
  1. `bridge`
  1. other enabled shells

- linux
  1. `user`
  1. `system`
  1. shell matching `CARAPACE_SHELL`
  1. `linux`
  1. `unix`
  1. `common`
  1. `bridge`
  1. other enabled shells

- termux
  1. `user`
  1. `system`
  1. shell matching `CARAPACE_SHELL`
  1. `android`
  1. `linux`
  1. `unix`
  1. `common`
  1. `bridge`
  1. other enabled shells

- windows
  1. `user`
  1. `system`
  1. shell matching `CARAPACE_SHELL`
  1. `windows`
  1. `common`
  1. `bridge`
  1. other enabled shells


You can **list** available completers of a **command** with `carapace --list {command}`.

![](./groups/priority.cast)
