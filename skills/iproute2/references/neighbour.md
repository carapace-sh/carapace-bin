# ip neighbour — ARP/NDISC Cache Management

Manage neighbour (ARP for IPv4, NDISC for IPv6) cache entries.

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Add a new neighbour entry |
| `del` | `delete` | Delete a neighbour entry |
| `change` | — | Change an existing entry |
| `replace` | — | Add new or change existing entry |
| `show` | `list` | List neighbour entries (default) |
| `flush` | — | Flush neighbour entries |
| `get` | — | Lookup a neighbour entry for a destination |

## add/del/change/replace Syntax

```
ip neigh { add | del | change | replace }
    { ADDR [ lladdr LLADDR ] [ nud STATE ] | proxy ADDR }
    [ dev DEV ] [ router ] [ use ] [ managed ]
    [ extern_learn ] [ extern_valid ]
```

| Keyword | Value Type | Description |
|---------|------------|-------------|
| `ADDR` | IP address | Protocol address (positional) |
| `lladdr` | `LLADDR` | Link-layer address (MAC) |
| `nud` | `STATE` | Neighbour state |
| `proxy` | `ADDR` | Proxy neighbour (no lladdr) |
| `dev` | `DEV` | Device name |
| `vrf` | `NAME` | VRF name (show/flush only) |

### Boolean flags (no value)

`router`, `use`, `managed`, `extern_learn`, `extern_valid`, `proxy`, `nomaster`, `unused`

## NUD State

| State | Description |
|-------|-------------|
| `permanent` | Entry is permanently valid (manually set) |
| `noarp` | No ARP needed (e.g. for point-to-point) |
| `reachable` | Entry is reachable (confirmed recently) |
| `stale` | Entry is valid but suspicious |
| `none` | Entry has no state (pending) |
| `incomplete` | Address resolution in progress |
| `delay` | Delayed resolution |
| `probe` | Probing to confirm reachability |
| `failed` | Address resolution failed |

## show/flush Filters

```
ip neigh { show | flush } [ proxy ] [ to PREFIX ] [ dev DEV ]
    [ nud STATE ] [ vrf NAME ] [ nomaster ] [ unused ]
```

| Filter | Description |
|--------|-------------|
| `proxy` | Show only proxy entries |
| `to PREFIX` | Filter by destination prefix |
| `dev DEV` | Filter by device |
| `nud STATE` | Filter by NUD state |
| `vrf NAME` | Filter by VRF |
| `nomaster` | Only entries without master |
| `unused` | Only unused entries (show only) |

## get

```
ip neigh get ADDR dev DEV
```

## References

- [ip-neighbour(8)](https://man7.org/linux/man-pages/man8/ip-neighbour.8.html)
