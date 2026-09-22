package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var program_extendCmd = &cobra.Command{
	Use:   "extend",
	Short: "Extend the length of an upgradeable program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_extendCmd).Standalone()

	program_extendCmd.Flags().BoolP("help", "h", false, "Print help")
	program_extendCmd.Flags().StringP("program-name", "p", "", "Program name to extend (from workspace). Used when program_id is not provided")
	programCmd.AddCommand(program_extendCmd)

	carapace.Gen(program_extendCmd).FlagCompletion(carapace.ActionMap{
		"program-name": anchor.ActionPrograms(),
	})
}
