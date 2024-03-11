package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()
	startCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	startCmd.Flags().String("script-shell", "", "shell to use")

	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"script-shell": os.ActionShells(),
	})
}
