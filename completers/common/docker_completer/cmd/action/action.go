package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
)

func ActionIpcModes() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValues("host", "private", "shareable"),
				carapace.ActionValues("container").Suffix(":"),
			).ToA()
		default:
			return docker.ActionContainers()
		}
	})
}

func ActionLinks() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return docker.ActionContainers().Suffix(":")
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionNetworkModes() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				docker.ActionNetworks(),
				carapace.ActionValues("bridge", "host", "none"),
				carapace.ActionValues("container").Suffix(":"),
			).ToA()
		default:
			return docker.ActionContainers()
		}
	})
}

func ActionPidModes() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValues("host"),
				carapace.ActionValues("container").Suffix(":"),
			).ToA()
		default:
			return docker.ActionContainers()
		}
	})
}

func ActionUtsModes() carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValues("host", "private"),
				carapace.ActionValues("container").Suffix(":"),
			).ToA()
		default:
			return docker.ActionContainers()
		}
	})
}
