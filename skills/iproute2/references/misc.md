# Miscellaneous ip Objects

Reference for the less complex ip object types: addrlabel, maddress, mroute, mrule, ntable, token, tcp_metrics, ioam, l2tp, mptcp, vrf, tuntap, monitor, fou, ila, macsec, netconf, nexthop, sr, stats.

## addrlabel

Label configuration for protocol address selection (RFC 3484).

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Add an address label entry |
| `del` | `delete` | Delete an address label entry |
| `list` | `show` | List address label entries (default) |
| `flush` | — | Flush all address labels |

### Syntax

```
ip addrlabel add prefix PREFIX [ dev DEV ] [ label NUMBER ]
ip addrlabel del prefix PREFIX [ dev DEV ] [ label NUMBER ]
ip addrlabel list
ip addrlabel flush
```

## maddress

Multicast address management.

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Add a static link-layer multicast address |
| `del` | `delete` | Delete a static multicast address |
| `show` | `list` | List multicast addresses (default) |

### Syntax

```
ip maddress add MULTIADDR dev NAME
ip maddress del MULTIADDR dev NAME
ip maddress show [ dev NAME ]
```

## mroute

Multicast routing cache entry. Display only — no add/delete.

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `show` | `list` | List multicast routing cache entries (default) |

### Syntax

```
ip mroute show [ [ to ] PREFIX ] [ from PREFIX ] [ iif DEVICE ] [ table TABLE_ID ]
```

## mrule

Rule in multicast routing policy database. Mirrors `ip rule` structure.

## ntable

Neighbor cache operation (per-device parameters for ARP/NDISC tables).

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `show` | `list` | List neighbour table parameters (default) |
| `change` | — | Modify neighbour table parameters |

### Syntax

```
ip ntable change name NAME [ dev DEV ] [ thresh1 VAL ] [ thresh2 VAL ]
    [ thresh3 VAL ] [ gc_int MSEC ] [ base_reachable MSEC ]
    [ retrans MSEC ] [ gc_stale MSEC ] [ delay_probe_time MSEC ]
    [ queue LEN ] [ app_probes VAL ] [ ucast_probes VAL ]
    [ mcast_probes VAL ] [ mcast_reprobes VAL ]
    [ anycast_delay MSEC ] [ proxy_delay MSEC ]
    [ proxy_queue LEN ] [ locktime MSEC ]

ip ntable show [ dev DEV ] [ name NAME ]
```

## token

Tokenized interface identifier support (IPv6).

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `set` | — | Set an interface token |
| `del` | — | Delete an interface token |
| `get` | — | Get the interface token for a device |
| `list` | `show` | List all interface tokens (default) |

### Syntax

```
ip token set TOKEN dev DEV
ip token del dev DEV
ip token get [ dev DEV ]
ip token [ list ]
```

## tcp_metrics / tcpmetrics

TCP Metrics management.

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `show` | `list` | Show cached TCP metrics entries (default) |
| `delete` | `del` | Delete a single TCP metrics entry |
| `flush` | — | Flush TCP metrics entries |

### Syntax

```
ip tcp_metrics show [ [ address ] PREFIX ]
ip tcp_metrics delete [ address ] ADDRESS
ip tcp_metrics flush [ [ address ] PREFIX ]
```

## ioam

IPv6 In-situ OAM (IOAM) configuration.

### Subcommands

| Subcommand | Commands |
|------------|----------|
| `namespace` | `show`, `add ID [ data DATA32 ] [ wide DATA64 ]`, `del ID`, `set ID schema { ID \| none }` |
| `schema` | `show`, `add ID DATA`, `del ID` |
| `monitor` | Display IOAM data received |

## l2tp

L2TPv3 static unmanaged tunnel configuration.

### Commands

| Command | Description |
|---------|-------------|
| `add tunnel` | Add a new L2TP tunnel |
| `del tunnel` | Destroy an L2TP tunnel |
| `show tunnel` | Show tunnel information |
| `add session` | Add a new session to a tunnel |
| `del session` | Destroy a session |
| `show session` | Show session information |

### Tunnel keywords

| Keyword | Value | Description |
|---------|-------|-------------|
| `tunnel_id` | `ID` | Local tunnel ID |
| `peer_tunnel_id` | `ID` | Peer tunnel ID |
| `remote` | `ADDR` | Remote IP (or `any`) |
| `local` | `ADDR` | Local IP (or `any`) |
| `encap` | `ip`/`udp` | Encapsulation type |
| `udp_sport` | `PORT` | UDP source port |
| `udp_dport` | `PORT` | UDP dest port |
| `udp_csum` | `on`/`off` | UDP checksums |
| `udp6_csum_tx` | `on`/`off` | IPv6 UDP TX checksums |
| `udp6_csum_rx` | `on`/`off` | IPv6 UDP RX checksums |

### Session keywords

| Keyword | Value | Description |
|---------|-------|-------------|
| `name` | `NAME` | Session name |
| `session_id` | `ID` | Local session ID |
| `peer_session_id` | `ID` | Peer session ID |
| `cookie` | `HEXSTR` | Cookie (8/16 hex digits) |
| `peer_cookie` | `HEXSTR` | Peer cookie |
| `l2spec_type` | `none`/`default` | L2 specific type |
| `seq` | `none`/`send`/`recv`/`both` | Sequence numbers |

## mptcp

MPTCP path manager configuration.

### Subcommands

| Subcommand | Commands |
|------------|----------|
| `endpoint` | `add`, `delete`, `change`, `show`, `flush` |
| `limits` | `set`, `show` |
| `monitor` | Display MPTCP connection events |

### endpoint

```
ip mptcp endpoint add IFADDR [ port PORT ] [ dev IFNAME ] [ id ID ] [ FLAG-LIST ]
ip mptcp endpoint delete id ID [ IFADDR ]
ip mptcp endpoint change [ id ID ] [ IFADDR ] [ port PORT ] CHANGE-OPT
ip mptcp endpoint show [ id ID ]
ip mptcp endpoint flush

FLAG := signal | subflow | laminar | backup | fullmesh | implicit
CHANGE-OPT := backup | nobackup | fullmesh | nofullmesh
```

### limits

```
ip mptcp limits set [ subflow SUBFLOW_NR ] [ add_addr_accepted ADD_ADDR_ACCEPTED_NR ]
ip mptcp limits show
```

## vrf

Virtual routing and forwarding device management.

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `show` | `list` | Show all configured VRF (default) |
| `exec` | — | Run a command against the named VRF |
| `identify` | — | Report VRF association for a process |
| `pids` | — | Report processes associated with a VRF |

### Syntax

```
ip vrf show [ NAME ]
ip vrf exec [ NAME ] command...
ip vrf identify [ PID ]
ip vrf pids NAME
```

## tuntap

TUN/TAP device management.

### Commands

| Command | Aliases | Description |
|---------|---------|-------------|
| `add` | — | Create a TUN or TAP device |
| `del` | `delete` | Remove a TUN/TAP device |
| `show` | `list` | List TUN/TAP devices (default) |

### Syntax

```
ip tuntap add [ dev NAME ] [ mode { tun | tap } ] [ user USERNAME ]
    [ group GROUPNAME ] [ one_queue ] [ pi ] [ vnet_hdr ] [ multi_queue ]
ip tuntap del [ dev NAME ] [ mode { tun | tap } ]
ip tuntap show
```

## monitor

Watch for netlink messages.

### Syntax

```
ip monitor [ all | OBJECT-LIST ] [ file FILENAME ] [ label ] [ all-nsid ]
    [ dev DEVICE ]

OBJECT-LIST := link | address | route | mroute | maddress | acaddress |
               prefix | neigh | netconf | rule | stats | nsid | nexthop
```

## fou

Foo-over-UDP receive port configuration.

### Commands

| Command | Description |
|---------|-------------|
| `add` | Add a FOU receive port |
| `del` | Delete a FOU receive port |
| `show` | Show FOU configuration |

### Syntax

```
ip fou add port PORT { gue | ipproto PROTO }
    [ local IFADDR ] [ peer IFADDR ] [ peer_port PORT ] [ dev IFNAME ]
ip fou del port PORT [ local IFADDR ] [ peer IFADDR ] [ peer_port PORT ] [ dev IFNAME ]
ip fou show
```

## ila

Identifier locator addresses.

## macsec

MACsec device configuration.

### Commands

| Command | Description |
|---------|-------------|
| `add` / `set` / `del` (tx/rx) | Manage send/receive associations |
| `offload` | Configure offload |
| `show` | Show MACsec configuration |

## netconf

Network configuration monitoring. Display only.

### Syntax

```
ip netconf show [ dev NAME ]
```

## nexthop

Nexthop object management.

### Commands

| Command | Description |
|---------|-------------|
| `add` | Add a nexthop |
| `replace` | Replace a nexthop |
| `del` | Delete a nexthop |
| `get` | Get a nexthop |
| `show` | Show nexthops (default) |
| `flush` | Flush nexthops |
| `bucket list` | List nexthop buckets |
| `bucket get` | Get a nexthop bucket |

### NH

```
NH := blackhole | [ via ADDRESS ] [ dev DEV ] [ onlink ] [ encap ENCAP ] [ fdb ]
      | group GROUP [ hw_stats { on | off } ] [ fdb ] [ type TYPE [ TYPE_ARGS ] ]
```

## sr

IPv6 Segment Routing.

### Commands

| Command | Description |
|---------|-------------|
| `hmac show` | Show HMAC configuration |
| `hmac set KEYID ALGO` | Set HMAC key (ALGO: sha1, sha256) |
| `tunsrc show` | Show tunnel source address |
| `tunsrc set ADDRESS` | Set tunnel source address |

## stats

Manage and show interface statistics.

### Commands

| Command | Description |
|---------|-------------|
| `show` | Show statistics (default) |
| `set` | Set statistics configuration |

### Syntax

```
ip stats show [ dev DEV ] [ group GROUP [ subgroup SUBGROUP [ suite SUITE ] ] ... ] ...
ip stats set dev DEV l3_stats { on | off }

GROUPS: link | offload | xstats | xstats_slave | afstats
```

## References

- [ip-addrlabel(8)](https://man7.org/linux/man-pages/man8/ip-addrlabel.8.html)
- [ip-fou(8)](https://man7.org/linux/man-pages/man8/ip-fou.8.html)
- [ip-ioam(8)](https://man7.org/linux/man-pages/man8/ip-ioam.8.html)
- [ip-l2tp(8)](https://man7.org/linux/man-pages/man8/ip-l2tp.8.html)
- [ip-macsec(8)](https://man7.org/linux/man-pages/man8/ip-macsec.8.html)
- [ip-maddress(8)](https://man7.org/linux/man-pages/man8/ip-maddress.8.html)
- [ip-monitor(8)](https://man7.org/linux/man-pages/man8/ip-monitor.8.html)
- [ip-mptcp(8)](https://man7.org/linux/man-pages/man8/ip-mptcp.8.html)
- [ip-mroute(8)](https://man7.org/linux/man-pages/man8/ip-mroute.8.html)
- [ip-netconf(8)](https://man7.org/linux/man-pages/man8/ip-netconf.8.html)
- [ip-nexthop(8)](https://man7.org/linux/man-pages/man8/ip-nexthop.8.html)
- [ip-ntable(8)](https://man7.org/linux/man-pages/man8/ip-ntable.8.html)
- [ip-sr(8)](https://man7.org/linux/man-pages/man8/ip-sr.8.html)
- [ip-stats(8)](https://man7.org/linux/man-pages/man8/ip-stats.8.html)
- [ip-tcp_metrics(8)](https://man7.org/linux/man-pages/man8/ip-tcp_metrics.8.html)
- [ip-token(8)](https://man7.org/linux/man-pages/man8/ip-token.8.html)
- [ip-vrf(8)](https://man7.org/linux/man-pages/man8/ip-vrf.8.html)
