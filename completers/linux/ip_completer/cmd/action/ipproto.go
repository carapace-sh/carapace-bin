package action

import "github.com/carapace-sh/carapace"

// ActionIpProtocols completes IP protocols (from /etc/protocols)
//
//	tcp (Transmission Control Protocol)
//	udp (User Datagram Protocol)
//	icmp (Internet Control Message Protocol)
func ActionIpProtocols() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ah", "Authentication Header",
		"dccp", "Datagram Congestion Control Protocol",
		"egp", "Exterior Gateway Protocol",
		"eigrp", "Enhanced Interior Gateway Routing Protocol",
		"esp", "Encapsulating Security Payload",
		"gre", "Generic Routing Encapsulation",
		"hopopt", "IPv6 Hop-by-Hop Option",
		"icmp", "Internet Control Message Protocol",
		"icmpv6", "Internet Control Message Protocol v6",
		"igmp", "Internet Group Management Protocol",
		"igp", "Interior Gateway Protocol",
		"ipcomp", "IP Payload Compression Protocol",
		"ipip", "IP Encapsulation within IP",
		"ipv6", "Internet Protocol version 6",
		"ipv6-route", "Routing Header for IPv6",
		"ipv6-frag", "Fragment Header for IPv6",
		"ipv6-icmp", "ICMP for IPv6",
		"ipv6-nonxt", "No Next Header for IPv6",
		"ipv6-opts", "Destination Options for IPv6",
		"l2tp", "Layer Two Tunneling Protocol",
		"ospf", "Open Shortest Path First",
		"pim", "Protocol Independent Multicast",
		"rsvp", "Resource Reservation Protocol",
		"rspd", "Radio Standard Simulation Protocol",
		"sctp", "Stream Control Transmission Protocol",
		"st", "ST Datamode",
		"tcp", "Transmission Control Protocol",
		"udp", "User Datagram Protocol",
		"udplite", "UDP-Lite",
		"vrrp", "Virtual Router Redundancy Protocol",
	).Tag("ip protocols")
}
