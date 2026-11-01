package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var program_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close a program or buffer account and withdraw all lamports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_closeCmd).Standalone()

	program_closeCmd.Flags().String("authority", "", "Authority keypair (defaults to configured wallet)")
	program_closeCmd.Flags().Bool("bypass-warning", false, "Bypass warning prompts")
	program_closeCmd.Flags().BoolP("help", "h", false, "Print help")
	program_closeCmd.Flags().StringP("program-name", "p", "", "Program name to close (from workspace). Used when account is not provided")
	program_closeCmd.Flags().String("recipient", "", "Recipient address for reclaimed lamports (defaults to authority)")
	programCmd.AddCommand(program_closeCmd)

	carapace.Gen(program_closeCmd).FlagCompletion(carapace.ActionMap{
		"authority":    carapace.ActionFiles(),
		"program-name": anchor.ActionPrograms(),
	})
}
