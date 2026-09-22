package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var program_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display information about a buffer or program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_showCmd).Standalone()

	program_showCmd.Flags().Bool("all", false, "Show all accounts")
	program_showCmd.Flags().Bool("get-buffers", false, "Get account information from the Solana config file")
	program_showCmd.Flags().Bool("get-programs", false, "Get account information from the Solana config file")
	program_showCmd.Flags().BoolP("help", "h", false, "Print help")
	programCmd.AddCommand(program_showCmd)
}
