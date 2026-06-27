package mount

import (
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
)

// ActionSources completes sources
//
//	/dev/sda1
//	LABEL=ROOT
func ActionSources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var keyValues carapace.Action
		switch runtime.GOOS {
		case "darwin":
			keyValues = carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"LABEL", "specifies device by volume name",
						"UUID", "specifies device by volume UUID",
					).Suffix("=")
				case 1:
					switch c.Parts[0] {
					case "LABEL":
						return fs.ActionLabels()
					case "UUID":
						return fs.ActionUuids()
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		default:
			keyValues = carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
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
			})
		}
		return carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
			keyValues,
		).ToA()
	}).Tag("sources")
}
