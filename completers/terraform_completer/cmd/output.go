package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Show output values from your root module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outputCmd).Standalone()

	outputCmd.Flags().Bool("json", false, "Machine readable output will be printed in JSON format.")
	outputCmd.Flags().Bool("no-color", false, "Output won't contain any color")
	outputCmd.Flags().Bool("raw", false, "print the raw string directly")
	outputCmd.Flags().String("state", "", "Path to the state file to read")
	rootCmd.AddCommand(outputCmd)

	carapace.Gen(outputCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})

	// TODO name positional completion
}
