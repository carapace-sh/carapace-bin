package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzz_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new fuzz harness for a program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzz_initCmd).Standalone()

	fuzz_initCmd.Flags().BoolP("help", "h", false, "Print help")
	fuzzCmd.AddCommand(fuzz_initCmd)
}
