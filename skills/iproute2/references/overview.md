# ip Command Overview

The top-level structure, global options, and object types of the `ip` command from iproute2.

## Top-Level Syntax

```
ip [ OPTIONS ] OBJECT { COMMAND | help }
ip [ -force ] -batch filename
```

When no COMMAND is given, the default is usually `list` (or `show`), or `help` if listing is not applicable.

## Global Options

| Short | Long / Variant | Value | Description |
|-------|----------------|-------|-------------|
| `-V` | `-Version` | — | Print version and exit |
| `-h` | `-human`, `-human-readable` | — | Human-readable stats with suffixes |
| `-b` | `-batch` | `<FILENAME>` | Read commands from file or stdin |
| — | `-force` | — | Don't terminate on batch errors |
| `-s` | `-stats`, `-statistics` | — | More info (repeatable) |
| `-d` | `-details` | — | More detailed info |
| `-l` | `-loops` | `<COUNT>` | Max flush loops (default 10, 0 = infinite) |
| `-f` | `-family` | `inet`/`inet6`/`bridge`/`mpls`/`link` | Protocol family |
| `-4` | — | — | Shortcut for `-family inet` |
| `-6` | — | — | Shortcut for `-family inet6` |
| `-B` | — | — | Shortcut for `-family bridge` |
| `-M` | — | — | Shortcut for `-family mpls` |
| `-0` | — | — | Shortcut for `-family link` |
| `-o` | `-oneline` | — | One record per line, LF → `\` |
| `-r` | `-resolve` | — | Use DNS to resolve hostnames |
| `-n` | `-netns` | `<NETNS>` | Switch to network namespace |
| `-N` | `-Numeric` | — | Print numeric values, no name conversion |
| `-a` | `-all` | — | Execute over all objects |
| `-c` | `-color` | `always`/`auto`/`never` | Color output (optional value) |
| `-t` | `-timestamp` | — | Timestamp with monitor |
| `-ts` | `-tshort` | — | Short timestamp with monitor |
| `-rc` | `-rcvbuf` | `<SIZE>` | Netlink socket receive buffer size (default 1MB) |
| — | `-iec` | — | Print rates in IEC units (1Ki = 1024) |
| `-br` | `-brief` | — | Tabular basic output |
| `-j` | `-json` | — | JSON output |
| `-p` | `-pretty` | — | Pretty-printed JSON |
| — | `-echo` | — | Request kernel to echo back applied config |

### Option Aliases

| Canonical | Aliases |
|-----------|---------|
| `-human-readable` | `-h`, `-human` |
| `-statistics` | `-s`, `-stats` |
| `-family` | `-f` |
| `-oneline` | `-o` |
| `-resolve` | `-r` |
| `-netns` | `-n` |
| `-Numeric` | `-N` |
| `-all` | `-a` |
| `-color` | `-c` |
| `-timestamp` | `-t` |
| `-tshort` | `-ts` |
| `-rcvbuf` | `-rc` |
| `-brief` | `-br` |
| `-json` | `-j` |
| `-pretty` | `-p` |
| `-details` | `-d` |
| `-loops` | `-l` |
| `-batch` | `-b` |

### Mutually Exclusive Groups

- `-human-readable` / `-human` — same option
- `-statistics` / `-stats` — same option

## Object Types

| Object | Aliases | Description |
|--------|---------|-------------|
| `address` | `addr`, `a` | Protocol (IP/IPv6) address on a device |
| `addrlabel` | — | Label configuration for protocol address selection |
| `fou` | — | Foo-over-UDP receive port configuration |
| `ila` | — | Identifier locator addresses |
| `ioam` | — | IOAM namespaces and schemas |
| `l2tp` | — | Tunnel ethernet over IP (L2TPv3) |
| `link` | `l` | Network device |
| `macsec` | — | MACsec device configuration |
| `maddress` | `maddr` | Multicast address |
| `monitor` | — | Watch for netlink messages |
| `mptcp` | — | Manage MPTCP path manager |
| `mroute` | — | Multicast routing cache entry |
| `mrule` | — | Rule in multicast routing policy database |
| `neighbour` | `neigh`, `n`, `neighbor` | ARP or NDISC cache entries |
| `netconf` | — | Network configuration monitoring |
| `netns` | — | Manage network namespaces |
| `nexthop` | — | Manage nexthop objects |
| `ntable` | `ntbl` | Manage the neighbor cache's operation |
| `route` | `r`, `ro` | Routing table entry |
| `rule` | — | Rule in routing policy database |
| `sr` | — | IPv6 segment routing |
| `stats` | — | Manage and show interface statistics |
| `tcp_metrics` | `tcpmetrics` | Manage TCP Metrics |
| `token` | — | Manage tokenized interface identifiers |
| `tunnel` | `tunl` | Tunnel over IP |
| `tuntap` | — | Manage TUN/TAP devices |
| `vrf` | — | Manage virtual routing and forwarding devices |
| `xfrm` | — | Manage IPSec policies |

## Default Commands

When no COMMAND is specified, the default behavior depends on the object:

| Object | Default Command |
|--------|----------------|
| `address` | `show` |
| `addrlabel` | `list` |
| `link` | `show` |
| `maddress` | `show` |
| `mroute` | `show` |
| `neighbour` | `show` |
| `netconf` | `show` |
| `nexthop` | `show` |
| `ntable` | `show` |
| `route` | `show` |
| `rule` | `show` |
| `tcp_metrics` | `show` |
| `token` | `list` |
| `tunnel` | `show` |
| `tuntap` | `show` |
| `vrf` | `show` |

## References

- [ip(8) man page](https://man7.org/linux/man-pages/man8/ip.8.html)
