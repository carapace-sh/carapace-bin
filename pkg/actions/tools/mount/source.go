package mount

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
)

// ActionSources completes sources
func ActionSources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
			carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"LABEL", "specifies device by filesystem label",
						"UUID", "specifies device by filesystem UUID",
						"PARTLABEL", "specifies device by partition label",
						"PARTUUID", "specifies device by partition UUID",
						"ID", "specifies device by udev hardware ID",
					).Suffix("=")
				case 1:
					switch c.Parts[0] {
					case "LABEL":
						return fs.ActionLabels()
					case "UUID":
						return fs.ActionUuids()
					case "PARTLABEL":
						return fs.ActionPartitionLabels()
					case "PARTUUID":
						return fs.ActionPartitionUuids()
					case "ID":
						// TODO
						return carapace.ActionValues()
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			}),
		).ToA()
	}).Tag("sources")
}
