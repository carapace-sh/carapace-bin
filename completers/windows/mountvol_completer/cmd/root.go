package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mountvol",
	Short: "create, delete, or list a volume mount point",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mountvol",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"list", "list mounted volumes",
			"delete", "delete a mount point",
			"noautodismount", "disable automount",
			"autodismount", "enable automount",
			"noserve", "disable direct mount",
			"serve", "enable direct mount",
		),
	)
}
