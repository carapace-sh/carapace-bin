package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot all machines in the deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebootCmd).Standalone()

	rebootCmd.Flags().StringSlice("exclude", nil, "Reboot all except the specified machines")
	rebootCmd.Flags().Bool("hard", false, "Send a hard reset to the machines")
	rebootCmd.Flags().StringSlice("include", nil, "Reboot only the specified machines")
	rebootCmd.Flags().Bool("no-wait", false, "Do not wait until the machines are up again")
	rebootCmd.Flags().Bool("rescue", false, "Reboot machines into the rescue system")

	rootCmd.AddCommand(rebootCmd)
}
