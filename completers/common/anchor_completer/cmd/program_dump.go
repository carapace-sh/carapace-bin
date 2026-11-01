package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var program_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Write the program data to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_dumpCmd).Standalone()

	program_dumpCmd.Flags().BoolP("help", "h", false, "Print help")
	programCmd.AddCommand(program_dumpCmd)
}
