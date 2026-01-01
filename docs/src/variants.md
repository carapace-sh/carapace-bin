> unreleased

# Variants

[Groups](./groups.md) can contain multiple completers for the same command.


You can **list** available **variants** for a command with `carapace --list {name}[/{variant}][@{group}]`.

![](./variants/variants.cast)

## Priority

Multiple **variants** within the same group are ordered by **name**.

- tldr
  1. `tealdeer`
  1. `tldr-python-client`
