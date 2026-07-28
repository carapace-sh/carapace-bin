package tailscale

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type peerStatus struct {
	HostName       string
	DNSName        string
	TailscaleIPs   []string
	Online         bool
	ExitNode       bool
	ExitNodeOption bool
}

type status struct {
	Self        peerStatus
	Peer        map[string]peerStatus
	CertDomains []string
}

func actionStatus(f func(s status) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("tailscale", "status", "--json")(func(output []byte) carapace.Action {
		var s status
		if err := json.Unmarshal(output, &s); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(s)
	})
}

// ActionHosts completes peers and self by hostname, DNS name, and Tailscale IPs
//
//	machine-alpha (100.64.0.1)
//	machine-beta (100.64.0.2)
func ActionHosts() carapace.Action {
	return actionStatus(func(s status) carapace.Action {
		vals := make([]string, 0)
		for _, peer := range s.peers() {
			desc := peer.HostName
			if len(peer.TailscaleIPs) > 0 {
				if desc != "" {
					desc += " "
				}
				desc += "(" + peer.TailscaleIPs[0] + ")"
			}
			vals = append(vals, peer.HostName, desc)
			if peer.DNSName != "" && peer.DNSName != peer.HostName {
				vals = append(vals, strings.TrimSuffix(peer.DNSName, "."), desc)
			}
			for _, ip := range peer.TailscaleIPs {
				vals = append(vals, ip, desc)
			}
		}
		return carapace.ActionValuesDescribed(vals...).Tag("tailscale peers").MultiParts(".")
	})
}

// ActionExitNodes completes peers that offer exit node capability
//
//	machine-gateway (100.64.0.3)
func ActionExitNodes() carapace.Action {
	return actionStatus(func(s status) carapace.Action {
		vals := make([]string, 0)
		for _, peer := range s.peers() {
			if !peer.ExitNodeOption && !peer.ExitNode {
				continue
			}
			desc := peer.HostName
			if len(peer.TailscaleIPs) > 0 {
				if desc != "" {
					desc += " "
				}
				desc += "(" + peer.TailscaleIPs[0] + ")"
			}
			vals = append(vals, peer.HostName, desc)
			for _, ip := range peer.TailscaleIPs {
				vals = append(vals, ip, desc)
			}
		}
		vals = append(vals, "auto:any", "automatically select an exit node")
		return carapace.ActionValuesDescribed(vals...).Tag("exit nodes")
	})
}

func (s status) peers() []peerStatus {
	peers := make([]peerStatus, 0, len(s.Peer)+1)
	peers = append(peers, s.Self)
	for _, p := range s.Peer {
		peers = append(peers, p)
	}
	return peers
}

// ActionCertDomains completes cert domains from tailscale status
//
//	machine-alpha.tailnet.ts.net
func ActionCertDomains() carapace.Action {
	return actionStatus(func(s status) carapace.Action {
		return carapace.ActionValues(s.CertDomains...).Tag("cert domains")
	})
}
