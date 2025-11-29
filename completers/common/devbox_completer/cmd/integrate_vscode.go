package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integrate_vscodeCmd = &cobra.Command{
	Use:    "vscode",
	Short:  "Integrate devbox environment with VSCode or other VSCode-based editors.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integrate_vscodeCmd).Standalone()

	integrate_vscodeCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	integrate_vscodeCmd.Flags().Bool("debugmode", false, "enable debug outputs to a file.")
	integrate_vscodeCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	integrate_vscodeCmd.Flags().String("ide", "", "name of the currently open editor to reopen after it's closed.")
	integrateCmd.AddCommand(integrate_vscodeCmd)

	// TODO environment, ide
	carapace.Gen(integrate_vscodeCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
