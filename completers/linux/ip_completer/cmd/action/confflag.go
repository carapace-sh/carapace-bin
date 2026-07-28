package action

import "github.com/carapace-sh/carapace"

func ActionConfFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"autojoin", "enable autojoin",
		"home", "(IPv6 only) designates this address the \"home address\"",
		"mngtmpaddr", "(IPv6 only) make the kernel manage temporary addresses created from this one as template",
		"nodad", "(IPv6 only) do not perform Duplicate Address Detection",
		"noprefixroute", "Do not automatically create a route for the network prefix of the added address",
		"optimistic", "(IPv6 only) When performing Duplicate Address Detection, use the RFC 4429 optimistic variant",
	)
}
