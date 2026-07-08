# ip xfrm — IPsec Transform Management

Manage IPsec Security Associations (SAD) and Security Policies (SPD).

## Object Structure

```
ip xfrm { state | policy | monitor } { COMMAND | help }
```

## state (Security Association Database — SAD)

### Commands

| Command | Description |
|---------|-------------|
| `add` | Add a new SA |
| `update` | Update an existing SA |
| `allocspi` | Allocate an SPI |
| `delete` | Delete an SA |
| `get` | Get an SA |
| `deleteall` | Delete all SAs matching criteria |
| `list` | List SAs |
| `flush` | Flush all SAs |
| `count` | Count SAs |

### ID

```
ID := [ src ADDR ] [ dst ADDR ] [ proto XFRM-PROTO ] [ spi SPI ]
```

### XFRM-PROTO

| Value | Description |
|-------|-------------|
| `esp` | Encapsulating Security Payload |
| `ah` | Authentication Header |
| `comp` | IP Payload Compression |
| `route2` | Route header (IPv6) |
| `hao` | Home Address Option (IPv6) |

### MODE

| Value | Description |
|-------|-------------|
| `transport` | Transport mode |
| `tunnel` | Tunnel mode |
| `beet` | Bound End-to-End Tunnel |
| `ro` | Route Optimization |
| `in_trigger` | In trigger |

### ALGO

```
ALGO := { enc | auth } ALGO-NAME ALGO-KEYMAT |
        auth-trunc ALGO-NAME ALGO-KEYMAT ALGO-TRUNC-LEN |
        aead ALGO-NAME ALGO-KEYMAT ALGO-ICV-LEN |
        comp ALGO-NAME
```

### FLAG

`noecn` | `decap-dscp` | `nopmtudisc` | `wildrecv` | `icmp` | `af-unspec` | `align4` | `esn`

### ENCAP

```
ENCAP := { espinudp | espinudp-nonike | espintcp } SPORT DPORT OADDR
```

### dir (SA direction)

`in` | `out`

## policy (Security Policy Database — SPD)

### Commands

| Command | Description |
|---------|-------------|
| `add` | Add a new policy |
| `update` | Update a policy |
| `delete` | Delete a policy |
| `get` | Get a policy |
| `deleteall` | Delete all policies |
| `list` | List policies |
| `flush` | Flush all policies |
| `count` | Count policies |
| `set` | Set policy thresholds (`hthresh4 LBITS RBITS`, `hthresh6 LBITS RBITS`) |
| `setdefault` | Set default policy |
| `getdefault` | Get default policy |

### SELECTOR

```
SELECTOR := [ src ADDR[/PLEN] ] [ dst ADDR[/PLEN] ] [ dev DEV ] [ UPSPEC ]
```

### dir (policy direction)

`in` | `out` | `fwd`

### ptype

`main` | `sub`

### action

`allow` | `block`

### TMPL

```
TMPL := ID [ mode MODE ] [ reqid REQID ] [ level LEVEL ]
LEVEL := required | use
```

### LIMIT

```
LIMIT := { time-soft | time-hard | time-use-soft | time-use-hard } SECONDS |
         { byte-soft | byte-hard } SIZE |
         { packet-soft | packet-hard } COUNT
```

## monitor

```
ip xfrm monitor [ all-nsid ] [ nokeys ] [ all | LISTofXFRM-OBJECTS ]
LISTofXFRM-OBJECTS := acquire | expire | SA | policy | aevent | report
```

## References

- [ip-xfrm(8)](https://man7.org/linux/man-pages/man8/ip-xfrm.8.html)
