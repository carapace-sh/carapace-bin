# ip tunnel — Tunnel Configuration

Manage IP tunnels (ipip, gre, sit, vti, etc.).

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Add a new tunnel |
| `change` | — | Change an existing tunnel |
| `del` | `delete` | Destroy a tunnel |
| `show` | `list` | List tunnels (default) |
| `prl` | — | Potential router list (ISATAP only) |
| `6rd` | — | 6rd (IPv6 Rapid Deployment) configuration |

## Syntax

```
ip tunnel { add | change | del | show | prl | 6rd }
    [ NAME ] [ mode MODE ] [ remote ADDR ] [ local ADDR ]
    [ [i|o]seq ] [ [i|o]key KEY ] [ [i|o]csum ]
    [ encaplimit ELIM ] [ ttl|hoplimit TTL ]
    [ tos TOS ] [ flowlabel FLOWLABEL ]
    [ prl-default ADDR ] [ prl-nodefault ADDR ] [ prl-delete ADDR ]
    [ 6rd-prefix ADDR ] [ 6rd-relay_prefix ADDR ] [ 6rd-reset ]
    [ [no]pmtudisc ] [ [no]ignore-df ] [ [no]allow-localremote ]
    [ dev PHYS_DEV ]
```

## MODE

| Mode | Description |
|------|-------------|
| `ipip` | IPv4-in-IPv4 |
| `gre` | GRE over IPv4 |
| `sit` | IPv6-in-IPv4 |
| `isatap` | Intra-Site Automatic Tunnel Addressing Protocol |
| `vti` | Virtual Tunnel Interface (IPv4) |
| `ip6ip6` | IPv6-in-IPv6 |
| `ipip6` | IPv4-in-IPv6 |
| `ip6gre` | GRE over IPv6 |
| `vti6` | Virtual Tunnel Interface (IPv6) |
| `any` | Any tunnel mode |

## Value-Taking Keywords

| Keyword | Value Type | Description |
|---------|------------|-------------|
| `mode` | `MODE` | Tunnel mode |
| `remote` | `ADDR` (IP or `any`) | Remote endpoint address |
| `local` | `ADDR` (IP or `any`) | Local endpoint address |
| `ttl` | `1..255` or `inherit` | TTL (IPv4) |
| `hoplimit` | `1..255` or `inherit` | Hop limit (IPv6) |
| `tos` | `STRING`/`00..ff`/`inherit` | TOS/DS field |
| `dsfield` | (alias for `tos`) | — |
| `tclass` | (IPv6 TOS equivalent) | Traffic class |
| `key` | `DOTTED_QUAD`/`NUMBER` | GRE key (both in/out) |
| `ikey` | `DOTTED_QUAD`/`NUMBER` | Input GRE key |
| `okey` | `DOTTED_QUAD`/`NUMBER` | Output GRE key |
| `encaplimit` | `none`/`0..255` | IPv6 encapsulation limit |
| `flowlabel` | `NUMBER` | IPv6 flow label |
| `dev` | `PHYS_DEV` | Physical device |

## Boolean/Negatable Keywords

| Keyword | Description |
|---------|-------------|
| `pmtudisc` / `nopmtudisc` | Path MTU discovery on/off |
| `ignore-df` / `no-ignore-df` | Ignore DF bit on/off |
| `allow-localremote` / `no-allow-localremote` | Allow local remote endpoint |
| `iseq` / `oseq` | Sequence numbers (in/out) |
| `icsum` / `ocsum` | Checksums (in/out) |
| `csum` | Checksums (both) |

## prl (ISATAP only)

```
ip tunnel prl NAME [ prl-default ADDR ] [ prl-nodefault ADDR ] [ prl-delete ADDR ]
```

## 6rd

```
ip tunnel 6rd NAME [ 6rd-prefix ADDR ] [ 6rd-relay_prefix ADDR ] [ 6rd-reset ]
```

## References

- [ip-tunnel(8)](https://man7.org/linux/man-pages/man8/ip-tunnel.8.html)
