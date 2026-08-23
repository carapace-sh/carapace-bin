package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportsauthrestartCmd = &cobra.Command{
	Use:   "supportsauthrestart",
	Short: "Check if the system supports authenticated restart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportsauthrestartCmd).Standalone()
	rootCmd.AddCommand(supportsauthrestartCmd)
}
