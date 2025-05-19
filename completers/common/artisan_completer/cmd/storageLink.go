package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storageLinkCmd = &cobra.Command{
	Use:   "storage:link [--relative] [--force]",
	Short: "Create the symbolic links configured for the application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storageLinkCmd).Standalone()

	storageLinkCmd.Flags().Bool("force", false, "Recreate existing symbolic links")
	storageLinkCmd.Flags().Bool("relative", false, "Create the symbolic link using relative paths")
	rootCmd.AddCommand(storageLinkCmd)
}
