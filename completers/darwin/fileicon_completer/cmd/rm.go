package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a custom icon from a file or folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().BoolP("quiet", "q", false, "silence status output")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(carapace.ActionFiles())
}
