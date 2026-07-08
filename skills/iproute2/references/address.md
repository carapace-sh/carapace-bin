# ip address — Protocol Address Management

Manage IP and IPv6 addresses on network devices.

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Add new protocol address |
| `change` | — | Change an existing address |
| `replace` | — | Add new or change existing address |
| `del` | `delete` | Delete protocol address |
| `show` | `list` | Look at protocol addresses (default) |
| `flush` | — | Flush protocol addresses matching criteria |
| `save` | — | Save address table to stdout (raw) |
| `restore` | — | Restore address table from stdin |
| `showdump` | — | Dump saved address data |

## IFADDR Syntax

```
IFADDR := PREFIX | ADDR peer PREFIX
           [ broadcast ADDR ] [ anycast ADDR ]
           [ label LABEL ] [ scope SCOPE-ID ] [ proto ADDRPROTO ]
```

| Component | Description |
|-----------|-------------|
| `PREFIX` | IP address with prefix length (e.g. `192.168.1.1/24`) |
| `ADDR peer PREFIX` | Point-to-point address with remote end |
| `broadcast ADDR` | Broadcast address (auto-detected if omitted) |
| `anycast ADDR` | Anycast address |
| `label LABEL` | Address label (e.g. `eth0:0`) |
| `scope SCOPE-ID` | Address scope |
| `proto ADDRPROTO` | Address protocol |

### SCOPE-ID

| Value | Description |
|-------|-------------|
| `host` | Address is valid only inside this host |
| `link` | Address is valid only on this link |
| `global` | Address is globally valid |
| `site` | Address is valid within the site (IPv6) |
| `NUMBER` | Numeric scope value |

### ADDRPROTO

Protocol that assigned the address. Can be a name or number. Common values: `kernel`, `dhcp`, `static`, `ra`.

## CONFFLAG-LIST

Configuration flags applied at add/change/replace time:

| Flag | Scope | Description |
|------|-------|-------------|
| `home` | IPv6 | Designates this address the "home address" |
| `nodad` | IPv6 | Do not perform Duplicate Address Detection |
| `mngtmpaddr` | IPv6 | Make kernel manage temporary addresses from this template |
| `optimistic` | IPv6 | Use address in optimistic DAD phase |
| `noprefixroute` | both | Do not auto-create route for the network prefix |
| `autojoin` | both | Enable autojoin (for multicast groups) |

## FLAG-LIST (for show/flush filtering)

| Flag | Description |
|------|-------------|
| `permanent` | Address is permanent (not temporary) |
| `dynamic` | Address is dynamically assigned |
| `secondary` | Address is secondary |
| `primary` | Address is primary |
| `tentative` | Address is tentative (DAD in progress) |
| `deprecated` | Address is deprecated |
| `dadfailed` | Duplicate Address Detection failed |
| `temporary` | Address is temporary (IPv6 privacy) |

Flags can be negated with a leading `-` (e.g. `-temporary`).

## LIFETIME

```
LIFETIME := [ valid_lft LFT ] [ preferred_lft LFT ]
LFT := [ forever | SECONDS ]
```

| Field | Description |
|-------|-------------|
| `valid_lft` | Valid lifetime in seconds (or `forever`) |
| `preferred_lft` | Preferred lifetime in seconds (or `forever`) |

## show/flush Additional Filters

| Filter | Description |
|--------|-------------|
| `dev IFNAME` | Device name |
| `scope SCOPE-ID` | Filter by scope |
| `to PREFIX` | Filter by prefix |
| `label PATTERN` | Filter by label pattern |
| `up` | Only addresses on up devices |
| `master DEVICE` | Filter by master device |
| `nomaster` | Only addresses without master |
| `type TYPE` | Filter by device type |
| `vrf NAME` | Filter by VRF |
| `metric METRIC` | Filter by metric |
| `proto ADDRPROTO` | Filter by address protocol |

## TYPE Values (device types for `type` filter)

See [link.md](link.md) for the full TYPE list. The address `type` filter accepts the same device types as `ip link add`.

## References

- [ip-address(8)](https://man7.org/linux/man-pages/man8/ip-address.8.html)
