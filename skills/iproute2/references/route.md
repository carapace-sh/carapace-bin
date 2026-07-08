# ip route — Routing Table Management

Manage routing table entries — add, delete, change, show, flush, and get routes.

## Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Add a new route |
| `del` | `delete` | Delete a route |
| `change` | — | Change a route |
| `append` | — | Append a route |
| `replace` | — | Change or add a new route |
| `show` | `list` | List routes (default) |
| `flush` | — | Flush routes matching criteria |
| `get` | — | Get a single route as kernel sees it |
| `save` | — | Save routing table to stdout (raw) |
| `restore` | — | Restore routing table from stdin |

## ROUTE Syntax

```
ROUTE := NODE_SPEC [ INFO_SPEC ]

NODE_SPEC := [ TYPE ] PREFIX [ tos TOS ] [ table TABLE_ID ]
             [ proto RTPROTO ] [ scope SCOPE ] [ metric METRIC ]
             [ ttl-propagate { enabled | disabled } ]

INFO_SPEC := { NH | nhid ID } OPTIONS FLAGS [ nexthop NH ] ...
```

## TYPE (route type)

| Value | Description |
|-------|-------------|
| `unicast` | Normal route to a destination (default) |
| `local` | Destination is assigned to this host |
| `broadcast` | Destination is a broadcast address |
| `multicast` | Special multicast route |
| `throw` | Route lookup fails (like unreachable without ICMP) |
| `unreachable` | Destination is unreachable (ICMP) |
| `prohibit` | Destination is prohibited (ICMP) |
| `blackhole` | Silently discard packets |
| `nat` | Network address translation route |

## TABLE_ID

| Value | Description |
|-------|-------------|
| `local` | Local routing table |
| `main` | Main routing table (default) |
| `default` | Default routing table |
| `all` | All routing tables |
| `NUMBER` | Numeric table ID |

## SCOPE

| Value | Description |
|-------|-------------|
| `host` | Route is on this host |
| `link` | Route is on this link |
| `global` | Route is globally valid |
| `NUMBER` | Numeric scope |

## RTPROTO

| Value | Description |
|-------|-------------|
| `kernel` | Route installed by kernel |
| `boot` | Route installed at boot |
| `static` | Route installed by administrator |
| `redirect` | Route installed by ICMP redirect |
| `NUMBER` | Numeric protocol ID |

## NH — Nexthop

```
NH := [ encap ENCAP ] [ via [ FAMILY ] ADDRESS ] [ dev STRING ]
      [ weight NUMBER ] NHFLAGS

FAMILY := inet | inet6 | mpls | bridge | link
NHFLAGS := onlink | pervasive
```

## OPTIONS (route attributes)

| Keyword | Value Type | Description |
|---------|------------|-------------|
| `mtu` | `NUMBER` | MTU for the route |
| `advmss` | `NUMBER` | Advertised MSS |
| `rtt` | `TIME` | Round-trip time estimate |
| `rttvar` | `TIME` | RTT variance |
| `reordering` | `NUMBER` | Packet reordering |
| `window` | `NUMBER` | TCP window |
| `cwnd` | `NUMBER` | Congestion window |
| `ssthresh` | `NUMBER` | Slow-start threshold |
| `realms` | `REALM` | Routing realms |
| `rto_min` | `TIME` | Minimum RTO |
| `initcwnd` | `NUMBER` | Initial congestion window |
| `initrwnd` | `NUMBER` | Initial receive window |
| `features` | `ecn` | ECN feature |
| `quickack` | `BOOL` | Quick ACK |
| `congctl` | `NAME` | Congestion control algorithm |
| `pref` | `low`/`medium`/`high` | Route preference |
| `expires` | `TIME` | Route expiry time |
| `fastopen_no_cookie` | `BOOL` | Fast open without cookie |

## ENCAP Types

| Type | Syntax |
|------|--------|
| `mpls` | `mpls [ LABEL ] [ ttl TTL ]` |
| `ip` | `ip id TUNNEL_ID dst REMOTE_IP [ src SRC ] [ tos TOS ] [ ttl TTL ]` |
| `bpf` | `bpf [ in PROG ] [ out PROG ] [ xmit PROG ] [ headroom SIZE ]` |
| `seg6` | `seg6 mode { encap \| encap.red \| inline \| l2encap \| l2encap.red } [ tunsrc ADDRESS ] segs SEGMENTS [ hmac KEYID ]` |
| `seg6local` | `seg6local action SEG6_ACTION [ SEG6_ACTION_PARAM ] [ count ]` |
| `ioam6` | `ioam6 [ freq K/N ] mode { inline \| encap \| auto } [ tunsrc ADDRESS ] [ tundst ADDRESS ] trace prealloc type IOAM6_TRACE_TYPE ns IOAM6_NAMESPACE size IOAM6_TRACE_SIZE` |
| `xfrm` | `xfrm` |

## SELECTOR (for show/flush)

```
SELECTOR := [ root PREFIX ] [ match PREFIX ] [ exact PREFIX ]
            [ table TABLE_ID ] [ vrf NAME ] [ proto RTPROTO ]
            [ type TYPE ] [ scope SCOPE ]
```

| Keyword | Value Type | Description |
|---------|------------|-------------|
| `root` | `PREFIX` | Root prefix to match |
| `match` | `PREFIX` | Prefix to match |
| `exact` | `PREFIX` | Exact prefix |
| `table` | `TABLE_ID` | Table to show |
| `vrf` | `NAME` | VRF name |
| `proto` | `RTPROTO` | Protocol |
| `type` | `TYPE` | Route type |
| `scope` | `SCOPE` | Route scope |

## route get

```
ip route get [ fibmatch ] ADDRESS [ from ADDRESS iif STRING ]
    [ oif STRING ] [ mark MARK ] [ tos TOS ] [ vrf NAME ]
    [ ipproto PROTOCOL ] [ sport NUMBER ] [ dport NUMBER ]
```

## Value-taking Keywords (add/del/change/append/replace)

| Keyword | Value Type |
|---------|------------|
| `to` / `PREFIX` | Subnet prefix |
| `via` | IP address |
| `dev` | Device name |
| `src` | IP address |
| `table` | Table name or number |
| `scope` | Scope name or number |
| `proto` | Protocol name or number |
| `metric` | Number |
| `mtu` | Number |
| `nhid` | Nexthop ID |
| `encap` | Encap type |
| `pref` | `low`/`medium`/`high` |
| `tos` | TOS value |
| `rtt` | Time |
| `window` | Number |
| `initcwnd` | Number |
| `congctl` | Algorithm name |
| `expires` | Time |

### Boolean keywords (no value)

`onlink`, `pervasive`, `ttl-propagate enabled`, `ttl-propagate disabled`

## References

- [ip-route(8)](https://man7.org/linux/man-pages/man8/ip-route.8.html)
