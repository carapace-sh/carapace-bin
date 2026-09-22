package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzz_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available fuzz tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzz_listCmd).Standalone()

	fuzz_listCmd.Flags().StringP("harness-dir", "C", "", "Use this directory instead of ./fuzz/<program_name>/")
	fuzz_listCmd.Flags().BoolP("help", "h", false, "Print help")
	fuzzCmd.AddCommand(fuzz_listCmd)

	carapace.Gen(fuzz_listCmd).FlagCompletion(carapace.ActionMap{
		"harness-dir": carapace.ActionFiles(),
	})
}
