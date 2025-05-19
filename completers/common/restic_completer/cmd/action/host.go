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
					"sftp", "SFTP",
					"rest", "HTTP",
					"s3", "Amazon S3",
					"swift", "OpenStack Swift",
					"b2", "Backblaze B2",
					"azure", "Microsoft Azure Blob Storage",
					"gs", "Google Cloud Storage",
					"rclone", "rclone",
				).Suffix(":")
			default:
				// TODO completion
				return carapace.ActionValues()
			}
		}),
	).ToA()
}
