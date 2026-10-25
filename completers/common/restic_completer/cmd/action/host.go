package action

import (
	"github.com/carapace-sh/carapace"
)

func ActionRepo() carapace.Action {
	return carapace.Batch(
		carapace.ActionDirectories(),
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"azure", "Microsoft Azure Blob Storage",
					"b2", "Backblaze B2",
					"gs", "Google Cloud Storage",
					"rclone", "rclone",
					"rest", "HTTP",
					"s3", "Amazon S3",
					"sftp", "SFTP",
					"swift", "OpenStack Swift",
				).Suffix(":")
			default:
				// TODO completion
				return carapace.ActionValues()
			}
		}),
	).ToA()
}
