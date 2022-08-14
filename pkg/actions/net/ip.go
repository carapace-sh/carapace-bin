package net

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/number"
)

// ActionIpv4Addresses completes ipv4 addresses
//
//	192.168.1.100
//	192.168.100.100/16
func ActionIpv4Addresses() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return number.ActionRange(number.RangeOpts{Format: "%v", Start: 0, End: 255}).Invoke(c).Suffix(".").ToA()
				case 1:
					return number.ActionRange(number.RangeOpts{Format: "%v", Start: 0, End: 255}).Invoke(c).Suffix(".").ToA()
				case 2:
					return number.ActionRange(number.RangeOpts{Format: "%v", Start: 0, End: 255}).Invoke(c).Suffix(".").ToA()
				case 3:
					return number.ActionRange(number.RangeOpts{Format: "%v", Start: 0, End: 255})
				default:
					return carapace.ActionValues()
				}
			})
		case 1:
			return ActionSubnets()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionSubnets completes subnets
//
//	24 (255.255.255.0 - 256)#
//	16 (255.255.0.0 - 65536)
func ActionSubnets() carapace.Action {
	return carapace.ActionValuesDescribed(
		"30", "255.255.255.252 - 4",
		"29", "255.255.255.248 - 8",
		"28", "255.255.255.240 - 16",
		"27", "255.255.255.224 - 32",
		"26", "255.255.255.192 - 64",
		"25", "255.255.255.128 - 128",
		"24", "255.255.255.0 - 256",
		"23", "255.255.254.0 - 512",
		"22", "255.255.252.0 - 1024",
		"21", "255.255.248.0 - 2048",
		"20", "255.255.240.0 - 4096",
		"19", "255.255.224.0 - 8192",
		"18", "255.255.192.0 - 16384",
		"17", "255.255.128.0 - 32768",
		"16", "255.255.0.0 - 65536",
	)
}
