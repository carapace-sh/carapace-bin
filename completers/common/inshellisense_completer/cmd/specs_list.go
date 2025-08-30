package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var specs_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the names of all available specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(specs_listCmd).Standalone()

	specs_listCmd.Flags().BoolP("help", "h", false, "display help for command")
	specs_listCmd.Flags().StringP("shell", "s", "", "shell to use alias specs, supported shells: bash, zsh")
	specsCmd.AddCommand(specs_listCmd)

	carapace.Gen(specs_listCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "zsh"),
	})
}
