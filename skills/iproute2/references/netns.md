# ip netns — Network Namespace Management

Manage named network namespaces.

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `list` | `show` | List named network namespaces (default) |
| `add` | — | Create a new named network namespace |
| `attach` | — | Create named namespace from a process PID |
| `del` | `delete` | Delete named namespace(s) |
| `set` | — | Assign an id to a peer network namespace |
| `identify` | — | Report network namespace names for a process |
| `pids` | — | Report processes in the named namespace |
| `exec` | — | Run a command in the named namespace |
| `monitor` | — | Report as namespace names are added/deleted |
| `list-id` | — | List network namespace ids (nsid) |

## Syntax

```
ip netns list
ip netns add NETNSNAME
ip netns attach NETNSNAME PID
ip [-all] netns del [ NETNSNAME ]
ip netns set NETNSNAME NETNSID
ip netns identify [ PID ]
ip netns pids NETNSNAME
ip [-all] netns exec [ NETNSNAME ] command...
ip netns monitor
ip netns list-id [ target-nsid POSITIVE-INT ] [ nsid POSITIVE-INT ]
```

## NETNSID

| Value | Description |
|-------|-------------|
| `auto` | Automatically assign namespace ID |
| `POSITIVE-INT` | Explicit numeric namespace ID |

## -all flag

When `-all` is used with `del` or `exec`, the operation applies to all named network namespaces.

## References

- [ip-netns(8)](https://man7.org/linux/man-pages/man8/ip-netns.8.html)
