# ip rule — Routing Policy Database

Manage routing policy rules — control which routing table is consulted for a given packet.

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `show` | `list`, `lst` | List rules (default) |
| `add` | — | Insert a new rule |
| `del` | `delete` | Delete a rule |
| `flush` | — | Flush all rules |
| `save` | — | Save rules table to stdout (raw) |
| `restore` | — | Restore rules table from stdin |

## SELECTOR (usable in show, add, del)

```
SELECTOR := [ not ] [ from PREFIX ] [ to PREFIX ] [ tos TOS ]
            [ dscp DSCP[/MASK] ] [ fwmark FWMARK[/MASK] ] [ iif STRING ]
            [ oif STRING ] [ priority PREFERENCE ] [ l3mdev ]
            [ uidrange NUMBER-NUMBER ] [ ipproto PROTOCOL ]
            [ sport [ NUMBER[/MASK] | NUMBER-NUMBER ] ]
            [ dport [ NUMBER[/MASK] | NUMBER-NUMBER ] ]
            [ tun_id TUN_ID ] [ flowlabel FLOWLABEL[/MASK] ]
```

| Keyword | Aliases | Value Type | Description |
|---------|---------|------------|-------------|
| `not` | — | *(prefix modifier)* | Negate the following selector (no value) |
| `from` | — | `PREFIX` | Source prefix to match |
| `to` | — | `PREFIX` | Destination prefix to match |
| `tos` | `dsfield` | `TOS` | TOS/DSCP value |
| `dscp` | — | `DSCP[/MASK]` | DSCP value with optional mask |
| `fwmark` | — | `FWMARK[/MASK]` | Firewall mark (optional mask) |
| `iif` | — | `STRING` | Incoming interface name |
| `oif` | — | `STRING` | Outgoing interface name |
| `priority` | `preference`, `order` | `NUMBER` | Rule priority |
| `l3mdev` | — | *(boolean)* | Match on L3 master device (no value) |
| `uidrange` | — | `NUMBER-NUMBER` | UID range |
| `ipproto` | — | `PROTOCOL` | IP protocol (name or number) |
| `sport` | — | `NUMBER[/MASK]` or `NUMBER-NUMBER` | Source port, port with mask, or range |
| `dport` | — | `NUMBER[/MASK]` or `NUMBER-NUMBER` | Destination port, port with mask, or range |
| `tun_id` | — | `TUN_ID` | Tunnel ID |
| `flowlabel` | — | `FLOWLABEL[/MASK]` | Flow label with optional mask |

### ipproto Values

IP protocol names from `/etc/protocols`:

```
ah | dccp | egp | eigrp | esp | gre | hopopt | icmp | icmpv6 | igmp |
igp | ipcomp | ipip | ipv6 | ipv6-frag | ipv6-icmp | ipv6-nonxt |
ipv6-opts | ipv6-route | l2tp | ospf | pim | rsvp | rspd | sctp |
st | tcp | udp | udplite | vrrp
```

## ACTION (only in add/del, NOT in show)

```
ACTION := [ table TABLE_ID ] [ protocol PROTO ] [ nat ADDRESS ]
          [ realms [SRCREALM/]DSTREALM ] [ goto NUMBER ] SUPPRESSOR
```

| Keyword | Aliases | Value Type | Description |
|---------|---------|------------|-------------|
| `table` | `lookup` | `TABLE_ID` | Routing table to consult |
| `protocol` | `proto` | `PROTO` | Routing protocol |
| `nat` | `map-to` | `ADDRESS` | Source NAT address |
| `realms` | — | `[SRCREALM/]DSTREALM` | Routing realms |
| `goto` | — | `NUMBER` | Jump to rule by priority |

### TABLE_ID

`local` | `main` | `default` | `all` | `NUMBER`

### protocol Values

`kernel` | `boot` | `static` | `NUMBER`

## SUPPRESSOR (part of ACTION, not a selector)

```
SUPPRESSOR := [ suppress_prefixlength NUMBER ] [ suppress_ifgroup GROUP ]
```

| Keyword | Value Type | Description |
|---------|------------|-------------|
| `suppress_prefixlength` | `NUMBER` | Reject routes with prefix length <= N |
| `suppress_ifgroup` | `GROUP` | Reject routes using device in interface group |

## type (rule type, part of ACTION)

| Value | Description |
|-------|-------------|
| `unicast` | Normal rule (default) |
| `blackhole` | Silently drop |
| `unreachable` | Drop with ICMP unreachable |
| `prohibit` | Drop with ICMP prohibited |
| `nat` | Network address translation |

## Selector vs Action Summary

| Category | Keywords |
|----------|----------|
| **Selectors** (show/add/del) | `not`, `from`, `to`, `tos`, `dsfield`, `dscp`, `fwmark`, `iif`, `oif`, `priority`, `preference`, `order`, `l3mdev`, `uidrange`, `ipproto`, `sport`, `dport`, `tun_id`, `flowlabel` |
| **Actions** (add/del only) | `table`, `lookup`, `protocol`, `proto`, `nat`, `map-to`, `realms`, `goto`, `type`, `suppress_prefixlength`, `suppress_ifgroup` |
| **Boolean** (no value) | `not` (prefix modifier), `l3mdev` (flag) |

## References

- [ip-rule(8)](https://man7.org/linux/man-pages/man8/ip-rule.8.html)
