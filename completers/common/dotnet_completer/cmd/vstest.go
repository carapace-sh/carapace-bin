package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vstestCmd = &cobra.Command{
	Use:   "vstest",
	Short: "run VSTest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vstestCmd).Standalone()
	rootCmd.AddCommand(vstestCmd)
}
