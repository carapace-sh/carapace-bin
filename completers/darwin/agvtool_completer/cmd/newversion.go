package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newVersionCmd = &cobra.Command{
	Use:   "new-version",
	Short: "Set a new version number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newVersionCmd).Standalone()
	rootCmd.AddCommand(newVersionCmd)
}
