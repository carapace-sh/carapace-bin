package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Move files out of the Tailscale file inbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_getCmd).Standalone()

	file_getCmd.Flags().String("conflict", "", "behavior when a conflicting file already exists in the target directory")
	file_getCmd.Flags().Bool("loop", false, "run get in a loop, receiving files as they come in")
	file_getCmd.Flags().Bool("verbose", false, "verbose output")
	file_getCmd.Flags().Bool("wait", false, "wait for a file to arrive if inbox is empty")
	fileCmd.AddCommand(file_getCmd)

	carapace.Gen(file_getCmd).FlagCompletion(carapace.ActionMap{
		"conflict": carapace.ActionValues("overwrite", "rename", "skip"),
	})

	carapace.Gen(file_getCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
