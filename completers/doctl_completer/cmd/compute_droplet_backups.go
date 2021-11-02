package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_backupsCmd = &cobra.Command{
	Use:   "backups",
	Short: "List Droplet backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_backupsCmd).Standalone()
	compute_droplet_backupsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_droplet_backupsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletCmd.AddCommand(compute_droplet_backupsCmd)

	carapace.Gen(compute_droplet_backupsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
