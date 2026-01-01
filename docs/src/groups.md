> unreleased

# Groups

Completers are organized into **groups**.

- `android` termux completers
- `bridge` bridged completers
- `bsd` bsd-like completers
- `common` common completers
- `darwin` macos completers
- `linux` linux completers
- `unix` unix-like completers
- `user` user specs
- `system` system specs
- `windows` windows completers

You can **list** available completers of a **group** with `carapace --list @{group}`.

![](./groups/group.cast)

> Binaries only contain **relevant groups** unless built with the [build tag](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags) `force_all`.

## Priority

Multiple **groups** providing a completer for a command are ordered by **priority**.

- darwin
  1. `user`
  1. `system`
  1. `darwin`
  1. `bsd`
  1. `unix`
  1. `common`
  1. `bridge`

- linux
  1. `user`
  1. `system`
  1. `linux`
  1. `unix`
  1. `common`
  1. `bridge`

- termux
  1. `user`
  1. `system`
  1. `android`
  1. `linux`
  1. `unix`
  1. `common`
  1. `bridge`

- windows
  1. `user`
  1. `system`
  1. `windows`
  1. `common`
  1. `bridge`


You can **list** available completers of a **command** with `carapace --list {command}`.

![](./groups/priority.cast)
