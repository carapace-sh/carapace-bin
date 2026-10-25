package qemu

import "github.com/carapace-sh/carapace"

// ActionNetdevTypes completes QEMU network backend types
//
//	user (user-mode networking (SLIRP))
//	tap (host TAP network backend)
func ActionNetdevTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"af-xdp", "AF_XDP socket",
		"bridge", "host TAP connected to bridge",
		"dgram", "datagram socket",
		"hubport", "hub port",
		"l2tpv3", "L2TPv3 pseudowire",
		"passt", "passt network backend",
		"socket", "socket network backend",
		"stream", "stream socket",
		"tap", "host TAP network backend",
		"user", "user-mode networking (SLIRP)",
		"vde", "VDE switch",
		"vhost-user", "vhost-user network",
		"vhost-vdpa", "vhost-vDPA network",
	).Tag("netdev types").Uid("qemu", "netdev-type")
}
