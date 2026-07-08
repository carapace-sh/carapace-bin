# ip link — Network Device Configuration

Manage network devices (interfaces) — create virtual links, change device attributes, display device information.

## Commands

| Command | Description |
|---------|-------------|
| `add` | Add a virtual link (create device) |
| `delete` | Delete a virtual link |
| `set` | Change device attributes |
| `show` | Display device attributes (default) |
| `xstats` | Show extended statistics by device type |
| `afstats` | Show address-family specific stats |
| `help` | Show help for link types |
| `property add` | Add alternative names (altname) to a device |
| `property del` | Delete alternative names from a device |

## link add

```
ip link add [ link DEVICE ] [ name ] NAME
    [ txqueuelen PACKETS ] [ address LLADDR ] [ broadcast LLADDR ]
    [ mtu MTU ] [ index IDX ]
    [ numtxqueues QUEUE_COUNT ] [ numrxqueues QUEUE_COUNT ]
    [ gso_max_size BYTES ] [ gso_ipv4_max_size BYTES ]
    [ gso_max_segs SEGMENTS ]
    [ gro_max_size BYTES ] [ gro_ipv4_max_size BYTES ]
    [ netns { PID | NETNSNAME | NETNSFILE } ]
    type TYPE [ ARGS ]
```

## link set

```
ip link set { DEVICE | group GROUP }
    [ { up | down } ]
    [ type ETYPE TYPE_ARGS ]
    [ arp { on | off } ]
    [ dynamic { on | off } ]
    [ multicast { on | off } ]
    [ allmulticast { on | off } ]
    [ promisc { on | off } ]
    [ protodown { on | off } ]
    [ trailers { on | off } ]
    [ txqueuelen PACKETS ]
    [ gso_max_size BYTES ] [ gso_ipv4_max_size BYTES ]
    [ gso_max_segs SEGMENTS ]
    [ gro_max_size BYTES ] [ gro_ipv4_max_size BYTES ]
    [ name NEWNAME ]
    [ address LLADDR ]
    [ broadcast LLADDR ]
    [ mtu MTU ]
    [ netns { PID | NETNSNAME | NETNSFILE } ]
    [ link-netnsid ID ]
    [ alias NAME ]
    [ vf NUM [ mac LLADDR ] [ VFVLAN-LIST ]
             [ rate TXRATE ] [ max_tx_rate TXRATE ] [ min_tx_rate TXRATE ]
             [ spoofchk { on | off } ] [ query_rss { on | off } ]
             [ state { auto | enable | disable } ] [ trust { on | off } ]
             [ node_guid eui64 ] [ port_guid eui64 ] ]
    [ { xdp | xdpgeneric | xdpdrv | xdpoffload }
      { off | object FILE [ { section | program } NAME ] [ verbose ]
      | pinned FILE } ]
    [ master DEVICE ] [ nomaster ]
    [ vrf NAME ]
    [ addrgenmode { eui64 | none | stable_secret | random } ]
    [ macaddr [ MACADDR ] [ { flush | add | del } MACADDR ] [ set MACADDR ] ]
```

### set Attributes (value-taking keywords)

| Keyword | Value Type | Description |
|---------|------------|-------------|
| `arp` | `on`/`off` | ARP on device |
| `dynamic` | `on`/`off` | Dynamic addressing |
| `multicast` | `on`/`off` | Multicast |
| `allmulticast` | `on`/`off` | All multicast |
| `promisc` | `on`/`off` | Promiscuous mode |
| `protodown` | `on`/`off` | Protocol down |
| `trailers` | `on`/`off` | Trailer encapsulation |
| `mtu` | `NUMBER` | Maximum transmission unit |
| `txqueuelen` | `PACKETS` | Transmit queue length |
| `name` | `NEWNAME` | Rename device |
| `address` | `LLADDR` | MAC address |
| `broadcast` | `LLADDR` | Broadcast address |
| `netns` | `PID`/`NETNSNAME`/`NETNSFILE` | Move to network namespace |
| `master` | `DEVICE` | Set master device |
| `vrf` | `NAME` | Set VRF |
| `alias` | `NAME` | Device alias |
| `addrgenmode` | `eui64`/`none`/`stable_secret`/`random` | IPv6 address generation mode |
| `xdp` | `off`/`object`/`pinned` | XDP program |
| `vf` | `NUM` + sub-options | Virtual function config |
| `macaddr` | `MACADDR` | MAC address list management |
| `protodown` | `on`/`off` | Protocol down state |
| `link-netnsid` | `ID` | Link network namespace ID |
| `gso_max_size` | `BYTES` | GSO max size |
| `gso_ipv4_max_size` | `BYTES` | GSO IPv4 max size |
| `gso_max_segs` | `SEGMENTS` | GSO max segments |
| `gro_max_size` | `BYTES` | GRO max size |
| `gro_ipv4_max_size` | `BYTES` | GRO IPv4 max size |

### Boolean keywords (no value)

`up`, `down`, `nomaster`

## link show

```
ip link show [ DEVICE | group GROUP ] [ up ] [ master DEVICE ]
    [ type ETYPE ] [ vrf NAME ] [ nomaster ] [ novf ]
```

## TYPE — Device Types

```
amt | bareudp | bond | bridge | can | dsa | dummy | erspan |
geneve | gre | gretap | gtp | hsr | ifb | ip6erspan | ip6gre |
ip6gretap | ip6tnl | ipip | ipoib | ipvlan | ipvtap | lowpan |
macsec | macvlan | macvtap | netdevsim | nlmon | rmnet | sit |
vcan | veth | virt_wifi | vlan | vrf | vti | vxcan | vxlan | wwan |
xfrm
```

### ETYPE

`TYPE` plus slave types: `bridge_slave`, `bond_slave`

## property

```
ip link property add dev DEVICE [ altname NAME ... ]
ip link property del dev DEVICE [ altname NAME ... ]
```

## References

- [ip-link(8)](https://man7.org/linux/man-pages/man8/ip-link.8.html)
