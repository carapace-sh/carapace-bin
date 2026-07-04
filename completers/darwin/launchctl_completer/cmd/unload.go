package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unload a service (deprecated, use bootout)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unloadCmd).Standalone()
	unloadCmd.Flags().BoolP("all", "a", false, "Unload all services in the directory")
	unloadCmd.Flags().BoolP("write", "w", false, "Override the Disabled key and persist the change")
	rootCmd.AddCommand(unloadCmd)
	carapace.Gen(unloadCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
