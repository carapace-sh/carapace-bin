---
name: iproute2
description: >
  Use when working with the Linux iproute2 networking commands — ip, ss, bridge, tc, and related tools.
  Covers the ip command argument structure, object types (address, link, route, rule, neighbour,
  tunnel, xfrm, netns, mptcp, ioam, l2tp, etc.), subcommands, selectors, and value types.
  Triggers on: "iproute2", "ip command", "ip address", "ip link", "ip route", "ip rule",
  "ip neighbour", "ip neigh", "ip tunnel", "ip xfrm", "ip netns", "ip mptcp", "ip ioam",
  "ip l2tp", "ip vrf", "ip token", "ip ntable", "ip tcp_metrics", "ip maddress",
  "ip mroute", "ip mrule", "ip addrlabel", "ip tuntap", "ip monitor", "ip stats",
  "ip nexthop", "ip netconf", "ip fou", "ip macsec", "ip sr", "ss command",
  "bridge command", "tc command", "IFADDR", "CONFFLAG", "SELECTOR", "ROUTE",
  "NUD state", "addrlabel", "ipproto".
user-invocable: true
---

# iproute2 In-Depth Reference

Reference for the [iproute2](https://wiki.linuxfoundation.org/networking/iproute2) networking utilities — the standard Linux network configuration toolkit. The primary focus is the `ip` command's argument structure, object types, subcommands, selectors, and value types.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| ip global options, -family, -json, -oneline, -netns, -color, -brief, object list, command overview, help | [references/overview.md](references/overview.md) |
| ip address, addr, IFADDR, PREFIX, CONFFLAG, FLAG-LIST, LIFETIME, SCOPE, scope, dev, label, broadcast, anycast, peer, mngtmpaddr, nodad, optimistic, noprefixroute, autojoin, home, permanent, dynamic, secondary, primary, tentative, deprecated, dadfailed, temporary, showdump, restore, save, flush, proto, ADDRPROTO | [references/address.md](references/address.md) |
| ip link, link add, link set, link show, link delete, TYPE, device type, bond, bridge, dummy, vlan, vxlan, gre, veth, macvlan, tun, tap, mtu, address, LLADDR, txqueuelen, master, nomaster, vrf, arp, multicast, promisc, xdp, vf, addrgenmode, altname, property, xstats, afstats, ETYPE, netns | [references/link.md](references/link.md) |
| ip route, route add, route del, route show, route flush, route get, ROUTE, NODE_SPEC, INFO_SPEC, NH, nexthop, nhid, via, dev, src, table, scope, proto, metric, encap, mpls, seg6, seg6local, ioam6, bpf, ip, xfrm, pref, onlink, ttl-propagate, mtu, rtt, window, cwnd, initcwnd, congctl, expires, SELECTOR, root, match, exact, fibmatch, save, restore, append, change, replace | [references/route.md](references/route.md) |
| ip rule, rule add, rule del, rule show, rule flush, rule save, rule restore, SELECTOR, not, from, to, tos, dsfield, fwmark, iif, oif, priority, preference, order, l3mdev, uidrange, ipproto, sport, dport, tun_id, ACTION, table, lookup, protocol, nat, map-to, realms, goto, type, unicast, blackhole, unreachable, prohibit, SUPPRESSOR, suppress_prefixlength, suppress_ifgroup | [references/rule.md](references/rule.md) |
| ip neighbour, ip neigh, ip n, neighbour add, neighbour del, neighbour show, neighbour flush, neighbour change, neighbour replace, neighbour get, NUD state, permanent, noarp, stale, reachable, none, incomplete, delay, probe, failed, lladdr, proxy, router, use, managed, extern_learn, extern_valid, vrf, nomaster | [references/neighbour.md](references/neighbour.md) |
| ip tunnel, tunnel add, tunnel change, tunnel del, tunnel show, tunnel prl, tunnel 6rd, MODE, ipip, gre, sit, isatap, vti, ip6ip6, ipip6, ip6gre, vti6, any, remote, local, ttl, hoplimit, tos, dsfield, tclass, flowlabel, key, ikey, okey, csum, icsum, ocsum, encaplimit, pmtudisc, ignore-df, allow-localremote, prl-default, prl-nodefault, prl-delete, 6rd-prefix, 6rd-relay_prefix, 6rd-reset | [references/tunnel.md](references/tunnel.md) |
| ip xfrm, xfrm state, xfrm policy, xfrm monitor, SAD, SPD, IPsec, ALGO, enc, auth, auth-trunc, aead, comp, MODE, transport, tunnel, beet, reqid, spi, dir, SA-DIR, SELECTOR, TMPL, LIMIT, FLAG, ENCAP, espinudp, espintcp, offload, tfcpad, allocspi, deleteall, count, set, setdefault, getdefault, hthresh4, hthresh6 | [references/xfrm.md](references/xfrm.md) |
| ip netns, netns list, netns add, netns attach, netns del, netns set, netns identify, netns pids, netns exec, netns monitor, netns list-id, NETNSID, network namespace | [references/netns.md](references/netns.md) |
| ip addrlabel, ip maddress, ip mroute, ip mrule, ip ntable, ip token, ip tcp_metrics, ip tcpmetrics, ip ioam, ip l2tp, ip mptcp, ip vrf, ip tuntap, ip monitor, ip fou, ip ila, ip macsec, ip netconf, ip nexthop, ip sr, ip stats, multicast, IOAM, namespace, schema, MPTCP, endpoint, limits, VRF, TUN, TAP, foo-over-udp, segment routing | [references/misc.md](references/misc.md) |

## Quick Guide

- **What are the global ip options?** → [references/overview.md](references/overview.md)
- **What object types does ip support?** → [references/overview.md](references/overview.md)
- **How does `ip address add` work?** → [references/address.md](references/address.md)
- **What link types can I create?** → [references/link.md](references/link.md)
- **How does `ip route` syntax work?** → [references/route.md](references/route.md)
- **What selectors does `ip rule` accept?** → [references/rule.md](references/rule.md)
- **What NUD states are valid?** → [references/neighbour.md](references/neighbour.md)
- **What tunnel modes exist?** → [references/tunnel.md](references/tunnel.md)
- **How does `ip xfrm` work?** → [references/xfrm.md](references/xfrm.md)
- **How do network namespaces work with ip?** → [references/netns.md](references/netns.md)
- **What about addrlabel, maddress, mptcp, l2tp, ioam, etc.?** → [references/misc.md](references/misc.md)
