package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showAccountCmd = &cobra.Command{
	Use:   "show-account",
	Short: "Show the contents of an account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showAccountCmd).Standalone()

	showAccountCmd.Flags().BoolP("help", "h", false, "Print help")
	showAccountCmd.Flags().Bool("lamports", false, "Display balance in lamports instead of SOL")
	showAccountCmd.Flags().String("output", "", "Return information in specified output format")
	showAccountCmd.Flags().StringP("output-file", "o", "", "Write the account data to this file")
	rootCmd.AddCommand(showAccountCmd)
}
