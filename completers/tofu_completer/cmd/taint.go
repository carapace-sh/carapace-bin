package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taintCmd = &cobra.Command{
	Use:   "taint [options] <address>",
	Short: "Mark a resource instance as not fully functional",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taintCmd).Standalone()

	taintCmd.Flags().BoolS("allow-missing", "allow-missing", false, "succeed even if the resource is missing.")
	taintCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	taintCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	taintCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	rootCmd.AddCommand(taintCmd)

	taintCmd.Flag("lock-timeout").NoOptDefVal = " "

	// TODO positional completion
}
