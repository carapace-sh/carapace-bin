package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfupdateCmd = &cobra.Command{
	Use:   "selfupdate",
	Short: "Update MacPorts system, ports tree, and base tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfupdateCmd).Standalone()
	selfupdateCmd.Flags().Bool("no-sync", false, "Only update MacPorts itself, do not update ports tree")
	rootCmd.AddCommand(selfupdateCmd)
}
