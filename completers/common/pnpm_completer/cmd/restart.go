package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a package. Runs \"stop\", \"restart\", and \"start\" scripts, and associated pre- and post- scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()

	restartCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	restartCmd.Flags().Bool("if-present", false, "Avoid exiting with a non-zero exit code when a script is undefined")
	rootCmd.AddCommand(restartCmd)
}
