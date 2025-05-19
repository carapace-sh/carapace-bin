package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untaintCmd = &cobra.Command{
	Use:   "untaint [options] name",
	Short: "Remove the 'tainted' state from a resource instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untaintCmd).Standalone()

	untaintCmd.Flags().BoolS("allow-missing", "allow-missing", false, "succeed even if the resource is missing.")
	untaintCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	untaintCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	untaintCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	rootCmd.AddCommand(untaintCmd)

	untaintCmd.Flag("lock-timeout").NoOptDefVal = " "

	// TODO positional completion
}
