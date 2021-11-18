package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var taintCmd = &cobra.Command{
	Use:   "taint",
	Short: "Mark a resource instance as not fully functional",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taintCmd).Standalone()

	taintCmd.Flags().Bool("allow-missing", false, "succeed even if the resource is missing.")
	taintCmd.Flags().Bool("ignore-remote-version", false, "A rare option used for the remote backend only.")
	taintCmd.Flags().String("lock", "", "Don't hold a state lock during the operation.")
	taintCmd.Flags().String("lock-timeout", "", "Duration to retry a state lock.")
	rootCmd.AddCommand(taintCmd)

	taintCmd.Flag("lock").NoOptDefVal = " "
	taintCmd.Flag("lock-timeout").NoOptDefVal = " "

	carapace.Gen(taintCmd).FlagCompletion(carapace.ActionMap{
		"lock": action.ActionBool(),
	})

	// TODO positional completion
}
