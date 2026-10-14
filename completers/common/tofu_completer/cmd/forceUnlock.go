package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forceUnlockCmd = &cobra.Command{
	Use:   "force-unlock LOCK_ID",
	Short: "Release a stuck lock on the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forceUnlockCmd).Standalone()

	forceUnlockCmd.Flags().BoolS("force", "force", false, "Don't ask for input for unlock confirmation.")
	forceUnlockCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	forceUnlockCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	forceUnlockCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	forceUnlockCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(forceUnlockCmd)

	forceUnlockCmd.Flag("json-into").NoOptDefVal = " "
	forceUnlockCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(forceUnlockCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})
}
