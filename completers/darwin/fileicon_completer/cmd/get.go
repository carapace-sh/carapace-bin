package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a file or folder's custom icon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolP("force", "f", false, "force replacement of existing output file")
	getCmd.Flags().BoolP("quiet", "q", false, "silence status output")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).PositionalCompletion(carapace.ActionFiles())
}
