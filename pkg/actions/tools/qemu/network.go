package qemu

import "github.com/carapace-sh/carapace"

// ActionNetdevTypes completes QEMU network backend types
//
//	user (user-mode networking (SLIRP))
//	tap (host TAP network backend)
func ActionNetdevTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"user", "user-mode networking (SLIRP)",
		"tap", "host TAP network backend",
		"bridge", "host TAP connected to bridge",
		"passt", "passt network backend",
		"l2tpv3", "L2TPv3 pseudowire",
		"vde", "VDE switch",
		"af-xdp", "AF_XDP socket",
		"socket", "socket network backend",
		"stream", "stream socket",
		"dgram", "datagram socket",
		"vhost-user", "vhost-user network",
		"vhost-vdpa", "vhost-vDPA network",
		"hubport", "hub port",
	).Tag("netdev types").Uid("qemu", "netdev-type")
}
