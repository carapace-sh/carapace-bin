package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "runas",
	Short: "run programs and tools with permissions different from the user's current logon",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/runas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("noprofile", "noprofile", false, "do not load the user profile")
	rootCmd.Flags().BoolP("env", "env", false, "use current network environment")
	rootCmd.Flags().BoolP("netonly", "netonly", false, "use specified credentials for remote access only")
	rootCmd.Flags().BoolP("trustlevel", "trustlevel", false, "run at a specified trust level")
}
