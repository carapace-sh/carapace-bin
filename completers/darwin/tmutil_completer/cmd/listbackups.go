package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listbackupsCmd)
}

var listbackupsCmd = &cobra.Command{
	Use:   "listbackups",
	Short: "list all of this computer's completed backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listbackupsCmd).Standalone()

	listbackupsCmd.Flags().StringP("destination", "d", "", "Destination volume to list backups from")
	listbackupsCmd.Flags().BoolP("mount", "m", false, "Attempt to mount backups")
	listbackupsCmd.Flags().BoolP("timestamp-only", "t", false, "Show only the backup timestamp")
}