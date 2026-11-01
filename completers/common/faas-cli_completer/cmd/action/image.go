package action

import "github.com/carapace-sh/carapace"

// ActionArchitectures completes architecture tokens used in store manifests,
// referenced by `generate --arch`, `store --platform` and `template store list --platform`.
func ActionArchitectures() carapace.Action {
	return carapace.ActionValuesDescribed(
		"arm64", "ARMv8 64-bit",
		"armhf", "ARMv6 32-bit",
		"x86_64", "x86 64-bit",
	)
}

// ActionPlatforms completes comma-separated buildx platforms for publishing.
func ActionPlatforms() carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		return carapace.ActionValuesDescribed(
			"linux/amd64", "x86 64-bit",
			"linux/arm64", "ARMv8 64-bit",
			"linux/arm/v5", "ARMv5 32-bit",
			"linux/arm/v6", "ARMv6 32-bit",
			"linux/arm/v7", "ARMv7 32-bit",
			"linux/ppc64le", "IBM POWER8",
			"linux/s390x", "IBM z Systems",
		)
	})
}
